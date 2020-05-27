package cmd

import (
	"SecondBrain/src/util"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var suppress bool

func init() {
	rootCmd.AddCommand(findCmd)
	findCmd.Flags().BoolVarP(&suppress, "suppress", "s", false, "Print less information")

	findCmd.AddCommand(findNoteCmd)
}

var findCmd = &cobra.Command{
	Use:   "find PATTERN [PATH]",
	Short: "Search through contents of Notes (default searches in the current Project)",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 && len(args) != 1 {
			return fmt.Errorf("this command takes 1 or 2 arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		searchPath := Config.Project
		if len(args) == 2 {
			searchPath = args[1]
		}

		pattern := args[0]
		grepCmd := fmt.Sprintf("grep -r -n -i '%s' %s", pattern, searchPath)
		grep, err := util.Exec(grepCmd)
		if err != nil {
			return fmt.Errorf("no results for '%+v'", pattern)
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
				if !suppress {
					fmt.Println("")
				}
				fmt.Printf("%+v\n", filename)
				foundMap[filename] = true
			}
			if !suppress {
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
