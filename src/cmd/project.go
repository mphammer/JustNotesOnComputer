package cmd

import (
	"SecondBrain/src/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var newProject bool
var listProjects bool

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.Flags().BoolVarP(&newProject, "new", "n", false, "Create a new Project")
	projectCmd.Flags().BoolVarP(&listProjects, "list", "l", false, "List all Projects")
}

var projectCmd = &cobra.Command{
	Use:     "project NAME",
	Short:   "List, Create, and Set Projects",
	Aliases: []string{"p", "proj"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("this command takes up to 1 argument")
		}
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		LoadConfig()
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
		if listProjects {
			// TODO list all projects with numbers (ex: [1] Staging [2] SlipBox)
			// TODO add ability for user to set the project by number
			return nil
		}
		if newProject {
			// Create a new Project
			execCmd := fmt.Sprintf("mkdir -p %s", args[0])
			_, err := util.Exec(execCmd)
			if err != nil {
				return fmt.Errorf("failed to execute '%+v': %+v", cmd, err)
			}
		}
		// Set the project
		projectPath := args[0]
		if num, err := strconv.Atoi(args[0]); err == nil {
			projectPath, err = GetDirByIndex(num)
			if err != nil {
				return fmt.Errorf("failed to get project number '%d': %+v", num, err)
			}
		}

		Config.Project = projectPath
		newDepth := len(strings.Split(projectPath, "/"))
		Config.ProjectDepth = newDepth

		configBytes, err := json.MarshalIndent(Config, "", " ")
		if err != nil {
			return fmt.Errorf("failed to serialize Vault Config: %+v", err)
		}

		err = ioutil.WriteFile(".vaultConfig.json", configBytes, 0644)
		if err != nil {
			return fmt.Errorf("failed to write to .vaultConfig.json: %+v", err)
		}
		fmt.Printf("Set Project: %s\n", projectPath)
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
	if NotProject(path) {
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
	if NotProject(currPath) {
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

func NotProject(projectPath string) bool {
	if len(projectPath) > 1 && (projectPath[0:1] == "." || projectPath[0:1] == "_") {
		return true
	}
	return false
}
