package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

type ConfigFile struct {
	Project      string        `json:"project"`
	ProjectDepth int           `json:"projectDepth"`
	History      HistoryConfig `json:"history"`
}

type HistoryConfig struct {
	Log            []string       `json:"log"`
	StartIndex     int            `json:"startIndex"`
	EndIndex       int            `json:"endIndex"`
	Length         int            `json:"length"`
	Capacity       int            `json:"capacity"`
	CommonCommands map[string]int `json:"commonCommands"`
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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		LoadConfig()
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		// Get Executed Command
		command := cmd.CommandPath()

		// Increment the Indexes
		Config.History.EndIndex = getHistoryIndexForward(Config.History.EndIndex)
		if Config.History.Length == Config.History.Capacity {
			Config.History.StartIndex = getHistoryIndexForward(Config.History.StartIndex)
		}

		// Update the Size
		if Config.History.Length != Config.History.Capacity {
			Config.History.Length = Config.History.Length + 1
		}

		// Write the command to history
		Config.History.Log[Config.History.EndIndex] = command

		// Save the config
		err := SaveConfig()
		if err != nil {
			return fmt.Errorf("failed to save config file: %+v", err)
		}

		return nil
	},
}

func getHistoryIndexForward(currIndex int) int {
	nextIndex := currIndex + 1
	if nextIndex >= Config.History.Capacity {
		nextIndex = 0
	}
	return nextIndex
}

func getHistoryIndexBackward(currIndex int) int {
	nextIndex := currIndex - 1
	if nextIndex < 0 {
		nextIndex = Config.History.Capacity - 1
	}
	return nextIndex
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
