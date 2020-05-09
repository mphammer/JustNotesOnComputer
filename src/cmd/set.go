package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Get files from Desktop or Downloads",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		return nil
	},
}
