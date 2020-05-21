package cmd

import (
	"SecondBrain/src/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var setNewProject bool

func init() {
	rootCmd.AddCommand(checkoutCmd)
	checkoutCmd.Flags().BoolVarP(&setNewProject, "new", "n", false, "Create and set to a new Project")
}

var checkoutCmd = &cobra.Command{
	Use:     "checkout NAME",
	Short:   "Switch Projects to work in",
	Aliases: []string{"c"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("this command takes up to 1 argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Printf("Current Project: %+v\n", Config.Project)
			err := ListProjects()
			if err != nil {
				return fmt.Errorf("%+v", err)
			}
			return nil
		}
		if setNewProject {
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

		if !util.NoteOrProjectExists(projectPath) {
			return fmt.Errorf("Project '%+v' does not exist", projectPath)
		}

		Config.Project = projectPath
		newDepth := len(strings.Split(projectPath, "/"))
		Config.ProjectDepth = newDepth

		err := SaveConfig()
		if err != nil {
			return fmt.Errorf("failed to save config file: %+v", err)
		}
		fmt.Printf("Set Project: %s\n", projectPath)
		return nil
	},
}
