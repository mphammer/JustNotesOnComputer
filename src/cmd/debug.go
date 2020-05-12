package cmd

import (
	"github.com/spf13/cobra"
)

// func init() {
// 	rootCmd.AddCommand(debugCmd)
// }

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug issues with your System of Notes",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check all links point to files
		// Check all ../'s go to the correct depth
		// Check no duplicate project names
		// Check the _data and _templates directories exist
		// Check the .vaultconfig.json exists
		// Check the .vaultconfig.json points to existing Project
		return nil
	},
}
