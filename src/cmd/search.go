package cmd

import (
	"SecondBrain/src/util"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var verbose bool
var all bool
var searchProjectName string

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print more information")
	searchCmd.Flags().BoolVarP(&all, "all", "a", false, "Search in all projects")
	searchCmd.Flags().StringVarP(&searchProjectName, "project", "p", "", "Project to search in")
}

var searchCmd = &cobra.Command{
	Use:   "search PATTERN",
	Short: "Search through notes (default searches in the current project)",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("this command takes 1 argument")
		}
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		LoadConfig()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		searchPath := Config.Project
		if searchProjectName != "" {
			searchPath = searchProjectName
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
			}
			if verbose {
				fmt.Printf("[%s] %s\n", lineNum, result)
			}
		}
		return nil
	},
}
