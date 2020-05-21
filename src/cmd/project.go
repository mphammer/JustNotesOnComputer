package cmd

import (
	"SecondBrain/src/util"
	"fmt"

	"github.com/spf13/cobra"
)

var newProject bool
var newProjectName string

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.Flags().BoolVarP(&newProject, "new", "n", false, "Create a new Project")
	projectCmd.Flags().StringVarP(&newProjectName, "rename", "r", "", "Rename the Project")
}

var projectCmd = &cobra.Command{
	Use:     "project [NAME]",
	Short:   "List, Create, and Rename Projects",
	Aliases: []string{"p", "proj"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("this command takes up to 1 argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			// List Projects
			fmt.Printf("Current Project: %+v\n", Config.Project)
			err := ListProjects()
			if err != nil {
				return fmt.Errorf("%+v", err)
			}
			return nil
		}
		if newProject {
			// Create a new Project
			execCmd := fmt.Sprintf("mkdir -p %s", args[0])
			_, err := util.Exec(execCmd)
			if err != nil {
				return fmt.Errorf("failed to execute '%+v': %+v", cmd, err)
			}
			fmt.Printf("Created Project: %s\n", args[0])
		}
		if newProjectName != "" {
			// TODO MoveDirectory()
			fmt.Printf("functionality not implemented yet... use move command\n")
		}
		return nil
	},
}

func ListProjects() error {
	_, dirnames, _, err := util.GetFilesAndDirectories(".")
	if err != nil {
		return fmt.Errorf("failed to get files and directories: %+v", err)
	}
	count := 1
	for _, dir := range dirnames {
		err = ListProjectsHelper(&count, dir)
		if err != nil {
			return fmt.Errorf("failed to print projects in '%+v': %+v", dir, err)
		}
	}
	return nil
}

func ListProjectsHelper(count *int, path string) error {
	if NotProjectPath(path) {
		return nil
	}
	currProjStr := " "
	if path == Config.Project {
		currProjStr = "*"
	}
	fmt.Printf("[%d] %s %s\n", *count, currProjStr, path)
	*count = (*count + 1)

	_, dirnames, _, err := util.GetFilesAndDirectories(path)
	if err != nil {
		return fmt.Errorf("failed to get files and directories: %+v", err)
	}
	for _, dir := range dirnames {
		dirPath := fmt.Sprintf("%s/%s", path, dir)
		err = ListProjectsHelper(count, dirPath)
		if err != nil {
			return fmt.Errorf("failed to print projects in '%+v': %+v", dir, err)
		}
	}
	return nil
}

func GetDirByIndex(targetIndex int) (string, error) {
	_, dirnames, _, err := util.GetFilesAndDirectories(".")
	if err != nil {
		return "", fmt.Errorf("failed to get files and directories: %+v", err)
	}
	currIndex := 1
	for _, dir := range dirnames {
		foundPath, err := GetDirByIndexHelper(dir, &currIndex, targetIndex)
		if err != nil {
			return "", fmt.Errorf("%+v", err)
		}
		if foundPath != "" {
			return foundPath, nil
		}
	}
	return "", fmt.Errorf("failed to find project at index %d", targetIndex)
}

func GetDirByIndexHelper(currPath string, currIndex *int, targetIndex int) (string, error) {
	if NotProjectPath(currPath) {
		return "", nil
	}
	if *currIndex == targetIndex {
		return currPath, nil
	}
	*currIndex = (*currIndex + 1)
	_, dirnames, _, err := util.GetFilesAndDirectories(currPath)
	if err != nil {
		return "", fmt.Errorf("failed to get files and directories: %+v", err)
	}
	for _, dir := range dirnames {
		dirPath := fmt.Sprintf("%s/%s", currPath, dir)
		foundPath, err := GetDirByIndexHelper(dirPath, currIndex, targetIndex)
		if err != nil {
			return "", fmt.Errorf("%+v", err)
		}
		if foundPath != "" {
			return foundPath, nil
		}
	}
	return "", err
}

func NotProjectPath(projectPath string) bool {
	if len(projectPath) > 1 && (projectPath[0:1] == "." || projectPath[0:1] == "_") {
		return true
	}
	return false
}
