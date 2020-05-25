package cmd

import (
	"SecondBrain/src/util"
	"fmt"

	"github.com/spf13/cobra"
)

var editWithTextEditor bool

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().BoolVarP(&editWithTextEditor, "open-with-text-editor", "o", false, "Open with default text editor")
}

var editCmd = &cobra.Command{
	Use:     "edit NOTE_PATH",
	Short:   "Edit a note",
	Aliases: []string{"e"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("this command takes 1 argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		notePath := args[0]
		execCmd := ""
		if editWithTextEditor {
			execCmd = fmt.Sprintf("open %s", notePath)
			_, err := util.Exec(execCmd)
			if err != nil {
				return fmt.Errorf("failed to exec: %+v", err)
			}
		} else {
			execCmd = fmt.Sprintf("vim %s", notePath)
			err := util.ExecShell(execCmd)
			if err != nil {
				return fmt.Errorf("failed to exec: %+v", err)
			}
		}
		return nil
	},
}
