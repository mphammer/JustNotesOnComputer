package cmd

import (
	"SecondBrain/src/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gobuffalo/packr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your Vault",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create directores
		util.Exec("mkdir _data")
		util.Exec("mkdir _templates")
		util.Exec("mkdir Staging")

		// Create config file
		c := ConfigFile{
			Project:      "Staging",
			ProjectDepth: 1,
		}
		configBytes, err := json.MarshalIndent(c, "", " ")
		if err != nil {
			return fmt.Errorf("failed to serialize Vault Config: %+v", err)
		}
		err = ioutil.WriteFile(".vaultConfig.json", configBytes, 0644)
		if err != nil {
			return fmt.Errorf("failed to write to .vaultConfig.json: %+v", err)
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
		fmt.Printf("For Command Completion copy this to .bashrc:\n")
		rootCmd.GenBashCompletion(os.Stdout)

		return nil
	},
}
