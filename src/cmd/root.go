package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

type ConfigFile struct {
	Project      string `json:"project"`
	ProjectDepth int    `json:"projectDepth"`
}

var Config ConfigFile

var ConfigName = ".jasonConfig.json"

func LoadConfig() {
	file, err := ioutil.ReadFile(ConfigName)
	if err != nil {
		fmt.Printf("failed to read %s\n", ConfigName)
		os.Exit(1)
	}
	err = json.Unmarshal([]byte(file), &Config)
	if err != nil {
		fmt.Printf("failed to unmarshal %s\n", ConfigName)
		os.Exit(1)
	}
}

func SaveConfig() error {
	configBytes, err := json.MarshalIndent(Config, "", " ")
	if err != nil {
		return fmt.Errorf("failed to serialize Vault Config: %+v", err)
	}

	err = ioutil.WriteFile(ConfigName, configBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to %s: %+v", ConfigName, err)
	}
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "jason",
	Short: "Tool for managing Just A System Of Notes",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
