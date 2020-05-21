package cmd

import (
	"SecondBrain/src/util"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var verbose bool
var all bool
var findCmdProjectName string

func init() {
	rootCmd.AddCommand(findCmd)
	findCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print more information")
	findCmd.Flags().BoolVarP(&all, "all", "a", false, "Search in all projects")
	findCmd.Flags().StringVarP(&findCmdProjectName, "project", "p", "", "Project to search in")

	findCmd.AddCommand(findNoteCmd)
}

var findCmd = &cobra.Command{
	Use:   "find PATTERN",
	Short: "Search through contents of Notes (default searches in the current Project)",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("this command takes 1 argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		searchPath := Config.Project
		if findCmdProjectName != "" {
			searchPath = findCmdProjectName
		}
		if all {
			searchPath = "."
		}

		pattern := args[0]
		grep, err := util.Exec(fmt.Sprintf("grep -r -n '%s' %s", pattern, searchPath))
		if err != nil {
			return fmt.Errorf("failed to grep: %+v", err)
		}

		grepLines := strings.Split(grep, "\n")

		foundMap := map[string]bool{}
		for _, line := range grepLines {
			splitLine := strings.SplitN(line, ":", 3)
			if len(splitLine) != 3 {
				continue
			}
			filename := splitLine[0]
			lineNum := splitLine[1]
			result := splitLine[2]
			if _, ok := foundMap[filename]; !ok {
				if verbose {
					fmt.Println("")
				}
				fmt.Printf("%+v\n", filename)
				foundMap[filename] = true
			}
			if verbose {
				fmt.Printf("[%s] %s\n", lineNum, result)
			}
		}
		return nil
	},
}

var findNoteCmd = &cobra.Command{
	Use:   "note PATTERN",
	Short: "Search through note names (default searches in the current project)",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("this command takes 1 argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		searchPath := Config.Project
		if findCmdProjectName != "" {
			searchPath = findCmdProjectName
		}
		if all {
			searchPath = "."
		}

		pattern := args[0]
		execCmd := fmt.Sprintf("find %s | grep '%s'", searchPath, pattern)
		grep, err := util.Exec(execCmd)
		if err != nil {
			return fmt.Errorf("failed to grep: %+v", err)
		}

		grepLines := strings.Split(grep, "\n")

		for _, line := range grepLines {
			if line == "" {
				continue
			}
			fmt.Printf("%+v\n", line)
		}
		return nil
	},
}
