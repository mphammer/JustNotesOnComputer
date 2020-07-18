package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"SecondBrain/src/util"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use: "remove NOTE",
	Example: `Remove a file:
$ ./jnoc rm <note_path>`,
	Aliases: []string{"rm"},
	Short:   "Remvoe a Note",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("this command takes 1 argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		targetPath := strings.Trim(args[0], "/")

		if !util.PathIsToFile(targetPath) {
			return fmt.Errorf("%+v is not a Note", targetPath)
		}

		MoveFile(targetPath, fmt.Sprintf(".trash/%s", targetPath))
		return nil
	},
}
