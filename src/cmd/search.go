package cmd

import (
	"SecondBrain/src/util"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var verbose bool

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print more information")
}

var searchCmd = &cobra.Command{
	Use:   "search PATTERN",
	Short: "Search through notes",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("this command takes 1 argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		pattern := args[0]
		grep, err := util.Exec(fmt.Sprintf("grep -r -n '%s' .", pattern))
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

		// fmt.Printf("%+v\n", grepLines)
		return nil
	},
}
