package cmd

import (
	"SecondBrain/src/util"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(debugCmd)
}

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug issues with your System of Notes",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO
		// Check all ../'s go to the correct depth
		// Check no duplicate project names

		// Verify _data directory exists
		_, err := util.IsDir("_data")
		if err != nil {
			return fmt.Errorf("Error checking _data: %+v", err)
		}

		// Verify _templates directory exists
		_, err = util.IsDir("_templates")
		if err != nil {
			return fmt.Errorf("Error checking _templates: %+v", err)
		}

		// Verify .jnocConfig file exists
		_, err = util.IsFile(".jnocConfig.json")
		if err != nil {
			return fmt.Errorf("Error checking .jnocConfig.json: %+v", err)
		}

		// Check the .vaultconfig.json points to existing Project
		_, err = util.IsDir(Config.Project)
		if err != nil {
			return fmt.Errorf("Error locating Project in file .jnocConfig: %+v", err)
		}

		err = CheckProjects(".")
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		return nil
	},
}

func CheckProjects(projectsPath string) error {
	_, dirnames, _, err := util.GetFilesAndDirectories(projectsPath)
	if err != nil {
		return fmt.Errorf("failed to get files and directories: %+v", err)
	}
	for _, dirname := range dirnames {
		if dirname[0:1] == "." || dirname[0:1] == "_" {
			continue
		}
		err = CheckDirectory(fmt.Sprintf("%s/%s", projectsPath, dirname))
		if err != nil {
			return fmt.Errorf("%+v", err)
		}
	}
	return nil
}

func CheckDirectory(dirpath string) error {
	_, dirnames, filenames, err := util.GetFilesAndDirectories(dirpath)
	if err != nil {
		return fmt.Errorf("failed to get files and directories: %+v", err)
	}

	for _, filename := range filenames {
		err = CheckFile(fmt.Sprintf("%s/%s", dirpath, filename))
		if err != nil {
			return fmt.Errorf("%+v", err)
		}
	}

	for _, dirname := range dirnames {
		if dirname[0:1] == "." || dirname[0:1] == "_" {
			continue
		}
		err = CheckDirectory(fmt.Sprintf("%s/%s", dirpath, dirname))
		if err != nil {
			return fmt.Errorf("%+v", err)
		}
	}
	return nil
}

func CheckFile(filePath string) error {
	lines, err := util.ReadFileLines(filePath)
	if err != nil {
		return fmt.Errorf("failed read old file: %s", err)
	}
	currFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed create file: %s", err)
	}
	regex := "\\[.+\\]\\((.+)\\)"
	re := regexp.MustCompile(regex)
	for _, txt := range lines {
		matches := re.FindStringSubmatch(txt)
		if len(matches) == 2 {
			link := matches[1]
			linkFilePath, err := filepath.Abs(fmt.Sprintf("%s/%s", filepath.Dir(filePath), link))
			if err != nil {
				return fmt.Errorf("%s", err)
			}
			isFile, err := util.IsFile(linkFilePath)
			if !isFile || err != nil {
				fmt.Printf("[ERROR] bad link in file '%s': '%s'\n", filePath, linkFilePath)
			}
		}
	}
	currFile.Close()

	return nil
}
