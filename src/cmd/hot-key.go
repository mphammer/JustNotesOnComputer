package cmd

import (
	"SecondBrain/src/util"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(hotKeyCmd)

	hotKeyCmd.AddCommand(setHotKeyCmd)
}

var hotKeyCmd = &cobra.Command{
	Use:     "hot-key [HOT_KEY]",
	Short:   "Execute a hot key",
	Aliases: []string{"hk"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			if len(Config.HotKeys) == 0 {
				fmt.Printf("No Hot Keys have been set\n")
				return nil
			}
			var keys []string
			for k := range Config.HotKeys {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, hk := range keys {
				fmt.Printf("%s - %s\n", hk, Config.HotKeys[hk])
			}
		} else {
			execCmd := Config.HotKeys[args[0]]
			out, err := util.Exec(execCmd)
			if err != nil {
				return fmt.Errorf("%s", err)
			}
			fmt.Printf("%s", out)
		}
		return nil
	},
}

var setHotKeyCmd = &cobra.Command{
	Use:     "set HOT_KEY COMMAND",
	Short:   "Set a hot key value",
	Example: "# set a hot-key\njnoc hot-key set 1 ls\n\n# delete a hot key\njnoc hot-key set 1 \"\"",
	RunE: func(cmd *cobra.Command, args []string) error {
		Config.HotKeys[args[0]] = args[1]
		if args[1] == "" {
			delete(Config.HotKeys, args[0])
		}

		// Save the config
		err := SaveConfig()
		if err != nil {
			return fmt.Errorf("failed to save config file: %+v", err)
		}
		return nil
	},
}
