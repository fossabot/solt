package cmd

import (
	"fmt"
	"github.com/aegoroff/godatastruct/collections"
	"github.com/aegoroff/godatastruct/rbtree"
	goahocorasick "github.com/anknown/ahocorasick"
	"github.com/spf13/afero"
	"os"
	"path/filepath"
	"solt/internal/msvc"
	"strings"

	"github.com/spf13/cobra"
)

// lostprojectsCmd represents the lostprojects command
var lostprojectsCmd = &cobra.Command{
	Use:     "lp",
	Aliases: []string{"lostprojects"},
	Short:   "Find projects that not included into any solution",
	RunE: func(cmd *cobra.Command, args []string) error {

		foldersTree := msvc.ReadSolutionDir(sourcesPath, appFileSystem, func(string) {})

		solutions := msvc.SelectSolutions(foldersTree)

		var projectsInSolutions []string
		projectsBySolution := make(map[string]collections.StringHashSet)
		// Each found solution
		for _, sln := range solutions {
			solutionProjectPaths := msvc.SelectAllSolutionProjectPaths(sln, false)
			projectsBySolution[sln.Path] = solutionProjectPaths
			for _, item := range solutionProjectPaths.Items() {
				projectsInSolutions = append(projectsInSolutions, strings.ToUpper(item))
			}
		}

		// Create projects matching machine
		pmm, err := createAhoCorasickMachine(projectsInSolutions)
		if err != nil {
			return err
		}

		projectsOutsideSolutions, filesInsideSolution := getOutsideProjectsAndFilesInsideSolution(foldersTree, pmm)

		lostProjects, lostProjectsThatIncludeSolutionProjectsFiles := separateProjects(projectsOutsideSolutions, filesInsideSolution)

		sortAndOutput(appWriter, lostProjects)

		if len(lostProjectsThatIncludeSolutionProjectsFiles) > 0 {
			_, _ = fmt.Fprintf(appWriter, "\nThese projects are not included into any solution but files from the projects' folders are used in another projects within a solution:\n\n")
		}

		sortAndOutput(appWriter, lostProjectsThatIncludeSolutionProjectsFiles)

		unexistProjects := getUnexistProjects(projectsBySolution, appFileSystem)

		if len(unexistProjects) > 0 {
			_, _ = fmt.Fprintf(appWriter, "\nThese projects are included into a solution but not found in the file system:\n")
		}

		outputSortedMap(appWriter, unexistProjects, "Solution")

		if showMemUsage {
			printMemUsage(appWriter)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(lostprojectsCmd)
}

func getUnexistProjects(projectsInSolutions map[string]collections.StringHashSet, fs afero.Fs) map[string][]string {
	var result = make(map[string][]string)
	for sol, projects := range projectsInSolutions {
		for _, prj := range projects.Items() {
			if _, err := fs.Stat(prj); !os.IsNotExist(err) {
				continue
			}

			if found, ok := result[sol]; ok {
				found = append(found, prj)
				result[sol] = found
			} else {
				result[sol] = []string{prj}
			}
		}
	}
	return result
}

func getOutsideProjectsAndFilesInsideSolution(foldersTree rbtree.RbTree, pmm *goahocorasick.Machine) ([]*msvc.MsbuildProject, collections.StringHashSet) {
	var projectsOutsideSolution []*msvc.MsbuildProject
	var filesInsideSolution = make(collections.StringHashSet)

	msvc.WalkProjects(foldersTree, func(prj *msvc.MsbuildProject, fold *msvc.Folder) {
		// Path in upper registry is the project's key
		projectKey := strings.ToUpper(prj.Path)

		ok := Match(pmm, projectKey)
		if !ok {
			projectsOutsideSolution = append(projectsOutsideSolution, prj)
		} else {
			filesIncluded := msvc.GetFilesIncludedIntoProject(prj)

			for _, f := range filesIncluded {
				filesInsideSolution.Add(strings.ToUpper(f))
			}
		}
	})

	return projectsOutsideSolution, filesInsideSolution
}

func separateProjects(projectsOutsideSolution []*msvc.MsbuildProject, filesInsideSolution collections.StringHashSet) ([]string, []string) {
	var projectsOutside []string
	var projectsOutsideSolutionWithFilesInside []string
	for _, prj := range projectsOutsideSolution {
		projectFiles := msvc.GetFilesIncludedIntoProject(prj)

		var includedIntoOther = false
		for _, f := range projectFiles {
			pf := strings.ToUpper(f)
			if !filesInsideSolution.Contains(pf) {
				continue
			}

			dir := filepath.Dir(prj.Path)

			if strings.Contains(pf, strings.ToUpper(dir)) {
				includedIntoOther = true
				break
			}
		}

		if !includedIntoOther {
			projectsOutside = append(projectsOutside, prj.Path)
		} else {
			projectsOutsideSolutionWithFilesInside = append(projectsOutsideSolutionWithFilesInside, prj.Path)
		}
	}
	return projectsOutside, projectsOutsideSolutionWithFilesInside
}
