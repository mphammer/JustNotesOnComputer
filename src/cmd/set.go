package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set PROJECT",
	Short: "Set a Project to be working in",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("this command takes 1 argument")
		}
		// TODO verify project exists
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		Config.Project = args[0]
		newDepth := len(strings.Split(args[0], "/"))
		Config.ProjectDepth = newDepth

		configBytes, err := json.MarshalIndent(Config, "", " ")
		if err != nil {
			return fmt.Errorf("failed to serialize Vault Config: %+v", err)
		}

		err = ioutil.WriteFile(".vaultConfig.json", configBytes, 0644)
		if err != nil {
			return fmt.Errorf("failed to write to .vaultConfig.json: %+v", err)
		}
		return nil
	},
}
