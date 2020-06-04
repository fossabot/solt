package cmd

import (
	"fmt"
	"github.com/aegoroff/godatastruct/collections"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"log"
	"solt/internal/msvc"
	"strings"
)

var subfolderToExclude = []string{
	"obj",
}

const filterParamName = "file"
const removeParamName = "remove"
const onlyLostParamName = "onlylost"

// lostfilesCmd represents the lostfiles command
var lostfilesCmd = &cobra.Command{
	Use:     "lf",
	Aliases: []string{"lostfiles"},
	Short:   "Find lost files in the folder specified",
	RunE: func(cmd *cobra.Command, args []string) error {

		lostFilesFilter, err := cmd.Flags().GetString(filterParamName)

		if err != nil {
			return err
		}

		removeLostFiles, err := cmd.Flags().GetBool(removeParamName)

		if err != nil {
			return err
		}

		onlyLost, err := cmd.Flags().GetBool(onlyLostParamName)

		if err != nil {
			return err
		}

		return executeLostFilesCommand(lostFilesFilter, removeLostFiles, onlyLost, appFileSystem)
	},
}

func init() {
	rootCmd.AddCommand(lostfilesCmd)
	lostfilesCmd.Flags().StringP(filterParamName, "f", ".cs", "Lost files filter extension. If not set .cs extension used")
	lostfilesCmd.Flags().BoolP(removeParamName, "r", false, "Remove lost files")
	lostfilesCmd.Flags().BoolP(onlyLostParamName, "l", false, "Show only lost files. Don't show unexist files. If not set all shown")
}

func executeLostFilesCommand(lostFilesFilter string, removeLostFiles bool, onlyLost bool, fs afero.Fs) error {
	lh := newLostFilesHandler(lostFilesFilter, fs)

	foldersTree := msvc.ReadSolutionDir(sourcesPath, fs, lh)

	msvc.WalkProjects(foldersTree, lh.projectHandler)

	lostFiles, err := findLostFiles(lh.excludeFolders, lh.foundFiles, lh.includedFiles)

	if err != nil {
		return err
	}

	sortAndOutput(appWriter, lostFiles)

	if !onlyLost {
		if len(lh.unexistFiles) > 0 {
			_, _ = fmt.Fprintf(appWriter, "\nThese files included into projects but not exist in the file system.\n")
		}

		outputSortedMap(appWriter, lh.unexistFiles, "Project")
	}

	if removeLostFiles {
		removeLostfiles(lostFiles, fs)
	}

	if showMemUsage {
		printMemUsage(appWriter)
	}
	return nil
}

func findLostFiles(excludeFolders collections.StringHashSet, foundFiles []string, includedFiles collections.StringHashSet) ([]string, error) {
	exm, err := createAhoCorasickMachine(excludeFolders.ItemsDecorated(normalize))
	if err != nil {
		return nil, err
	}

	var result []string
	for _, file := range foundFiles {
		normalized := normalize(file)
		if !includedFiles.Contains(normalized) && !Match(exm, normalized) {
			result = append(result, file)
		}
	}

	return result, err
}

func normalize(s string) string {
	return strings.ToUpper(s)
}

func removeLostfiles(lostFiles []string, fs afero.Fs) {
	for _, f := range lostFiles {
		err := fs.Remove(f)
		if err != nil {
			log.Printf("%v\n", err)
		} else {
			_, _ = fmt.Fprintf(appWriter, "File: %s removed successfully.\n", f)
		}
	}
}
