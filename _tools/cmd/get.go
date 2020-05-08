package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get files from Desktop or Downloads",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		return nil
	},
}
