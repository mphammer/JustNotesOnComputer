package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bangVal int

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().IntVarP(&bangVal, "bang", "b", 0, "Execute a command from history")
}

var historyCmd = &cobra.Command{
	Use:     "history",
	Short:   "Print history information",
	Aliases: []string{"h"},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		startIndex := Config.History.StartIndex
		currIndex := startIndex
		numToPrint := Config.History.Length
		for numToPrint > 0 {
			// Print the command
			fmt.Printf("[%d] %s\n", numToPrint, Config.History.Log[currIndex])

			// Increment the current index
			currIndex = getHistoryIndexForward(currIndex)

			// Decrement number left to print
			numToPrint = numToPrint - 1
		}
		return nil
	},
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
