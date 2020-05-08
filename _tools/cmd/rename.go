package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(renameCmd)
}

var renameCmd = &cobra.Command{
	Use:     "rename",
	Aliases: []string{"rn"},
	Short:   "Print the version number of Hugo",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		return nil
	},
}
