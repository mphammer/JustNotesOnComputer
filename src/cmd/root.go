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

func LoadConfig() {
	file, err := ioutil.ReadFile(".vaultConfig.json")
	if err != nil {
		fmt.Println("failed to read .vaultConfig.json")
		os.Exit(1)
	}
	err = json.Unmarshal([]byte(file), &Config)
	if err != nil {
		fmt.Println("failed to unmarshal .vaultConfig.json")
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "vault",
	Short: "Your second brain",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
