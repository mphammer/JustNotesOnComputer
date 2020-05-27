package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(hotKeyCmd)
}

var hotKeyCmd = &cobra.Command{
	Use:   "hot-key NUM",
	Short: "Execute a saved command",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
