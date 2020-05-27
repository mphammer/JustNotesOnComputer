package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ConfigFile struct {
	Project      string            `json:"project"`
	ProjectDepth int               `json:"projectDepth"`
	HotKeys      map[string]string `json:"hotKeys"`
	History      HistoryConfig     `json:"history"`
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

var ConfigName = ".jnocConfig.json"

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

var version bool

var rootCmd = &cobra.Command{
	Use: "jnoc",
	Example: `Create a Project:
$ ./jnoc project My/First/Project --new

Checkout the Project to easily create Notes in it:
$ ./jnoc checkout My/First/Project

Create a Note:
$ ./jnoc note # You will be prompted for a Note Name
`,
	SilenceErrors: true,
	SilenceUsage:  true,
	Short:         "Just a tool for managing your Notes On Computer",
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			v := "1.0.0"
			fmt.Printf("Just Notes On Computer %s\n", v)
		}
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

		// Create args string
		argsString := strings.Join(args, " ")

		// Create flags string
		flagsString := ""
		flags := cmd.Flags()
		flags.Visit(func(f *pflag.Flag) {
			name := f.Name
			value := f.Value.String()
			flagsString = fmt.Sprintf("%s --%s %s", flagsString, name, value)
		})

		// Write the command to history
		commandString := fmt.Sprintf("./%s", command)
		if argsString != "" {
			commandString = fmt.Sprintf("%s %s", commandString, argsString)
		}
		if flagsString != "" {
			commandString = fmt.Sprintf("%s %s", commandString, flagsString)
		}
		Config.History.Log[Config.History.EndIndex] = commandString

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
	rootCmd.Flags().BoolVar(&version, "version", false, "Version of jnoc")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
