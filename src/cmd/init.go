package cmd

import (
	"SecondBrain/src/util"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gobuffalo/packr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your System of Notes",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create directores
		util.Exec("mkdir _data")
		util.Exec("mkdir _templates")
		util.Exec("mkdir Staging")

		// Create config file
		Config = ConfigFile{
			Project:      "Staging",
			ProjectDepth: 1,
			HotKeys:      map[string]string{},
			History: HistoryConfig{
				Log:            make([]string, 20),
				StartIndex:     1,
				EndIndex:       0,
				Length:         0,
				Capacity:       20,
				CommonCommands: map[string]int{},
			},
		}
		configBytes, err := json.MarshalIndent(Config, "", " ")
		if err != nil {
			return fmt.Errorf("failed to serialize Vault Config: %+v", err)
		}
		err = ioutil.WriteFile(ConfigName, configBytes, 0644)
		if err != nil {
			return fmt.Errorf("failed to write to %s: %+v", ConfigName, err)
		}

		// Create Templates
		box := packr.NewBox("../../_templates")
		for _, templateName := range []string{"BookSummary", "Contact", "Journal", "Note"} {
			templateBytes, err := box.Find(fmt.Sprintf("%s.md", templateName))
			err = ioutil.WriteFile(fmt.Sprintf("_templates/%s.md", templateName), templateBytes, 0644)
			if err != nil {
				return fmt.Errorf("failed to write to _templates/%s.md: %+v", templateName, err)
			}
		}

		// Bash Completion - https://github.com/spf13/cobra/blob/master/bash_completions.md
		fmt.Printf("For Command Completion copy out.sh\n")
		rootCmd.GenBashCompletionFile("out.sh")

		return nil
	},
}
