package main

import (
	"fmt"
	"path/filepath"
	"solt/solution"
	"strings"
)

func lostprojectscmd(opt options) error {

	var solutions []string
	folders := readProjectDir(opt.Path, func(we *walkEntry) {
		ext := strings.ToLower(filepath.Ext(we.Name))
		if ext == solutionFileExt {
			sp := filepath.Join(we.Parent, we.Name)
			solutions = append(solutions, sp)
		}
	})

	allProjectsWithinSolutions := getAllSolutionsProjects(solutions)

	projectsOutsideSolution, filesInsideSolution := getOutsideProjectsAndFilesInsideSolution(folders, allProjectsWithinSolutions)

	projectsOutside, projectsOutsideSolutionWithFilesInside := separateOutsideProjects(projectsOutsideSolution, filesInsideSolution)

	sortAndOutput(projectsOutside)

	if len(projectsOutsideSolutionWithFilesInside) > 0 {
		fmt.Printf("\nThese projects not included into any solution but their files used in projects that included into another projects within a solution.\n")
	}

	sortAndOutput(projectsOutsideSolutionWithFilesInside)

	return nil
}

func getOutsideProjectsAndFilesInsideSolution(folders []*folderInfo, allProjectsWithinSolutions map[string]interface{}) ([]*folderInfo, map[string]interface{}) {
	var projectsOutsideSolution []*folderInfo
	var filesInsideSolution = make(map[string]interface{})
	for _, info := range folders {
		if info.project == nil {
			continue
		}
		project := *info.projectPath
		_, ok := allProjectsWithinSolutions[project]
		if !ok {
			projectsOutsideSolution = append(projectsOutsideSolution, info)
		} else {
			filesIncluded := getFilesIncludedIntoProject(info)

			for _, f := range filesIncluded {
				filesInsideSolution[f] = nil
			}
		}
	}
	return projectsOutsideSolution, filesInsideSolution
}

func separateOutsideProjects(projectsOutsideSolution []*folderInfo, filesInsideSolution map[string]interface{}) ([]string, []string) {
	var projectsOutside []string
	var projectsOutsideSolutionWithFilesInside []string
	for _, info := range projectsOutsideSolution {
		filesIncluded := getFilesIncludedIntoProject(info)

		var includedIntoOther = false
		for _, f := range filesIncluded {
			if _, ok := filesInsideSolution[f]; ok {

				dir := filepath.Dir(*info.projectPath)

				if strings.Contains(f, dir) {
					includedIntoOther = true
					break
				}
			}
		}

		if !includedIntoOther {
			projectsOutside = append(projectsOutside, *info.projectPath)
		} else {
			projectsOutsideSolutionWithFilesInside = append(projectsOutsideSolutionWithFilesInside, *info.projectPath)
		}
	}
	return projectsOutside, projectsOutsideSolutionWithFilesInside
}

func getAllSolutionsProjects(solutions []string) map[string]interface{} {
	var projectsInSolution = make(map[string]interface{})
	for _, solpath := range solutions {
		sln, _ := solution.Parse(solpath)

		parent := filepath.Dir(solpath)

		for _, p := range sln.Projects {
			// Skip solution folders
			if p.TypeId == "{2150E333-8FDC-42A3-9474-1A3956D46DE8}" {
				continue
			}

			pp := filepath.Join(parent, p.Path)

			if _, ok := projectsInSolution[pp]; !ok {
				projectsInSolution[pp] = nil
			}
		}
	}
	return projectsInSolution
}
