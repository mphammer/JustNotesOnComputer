package cmd

import (
	"github.com/spf13/cobra"
)

// func init() {
// 	rootCmd.AddCommand(infoCmd)
// }

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print info about project, file, tag, or entire vault",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Print number of files in a project
		// Print number of files that reference a file
		// Print number files that have a tag
		// Print all of the tags and how many there are
		// Print the number of tags, files, and projects in the vault
		return nil
	},
}
