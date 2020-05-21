package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

var printTree bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&printTree, "tree", "t", false, "Print a tree of files")
}

var listCmd = &cobra.Command{
	Use:     "list [PROJECT]",
	Short:   "List Notes, Projects, and Tags",
	Aliases: []string{"ls"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("this command takes up to 1 argument")
		}
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		LoadConfig()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		if printTree {
			if len(args) == 0 {
				path = "."
			}
			err := PrintTree(path, 0)
			if err != nil {
				return fmt.Errorf("%+v", err)
			}
		} else {
			err := PrintDirectoryContents(path)
			if err != nil {
				return fmt.Errorf("%+v", err)
			}
		}
		return nil
	},
}

func PrintDirectoryContents(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("%+v", err)
	}

	for _, f := range files {
		if f.Name()[0:1] == "." || f.Name()[0:1] == "_" {
			continue
		}
		if f.IsDir() {
			fmt.Printf("%+v\n", f.Name())
		} else {
			fmt.Printf("%+v\n", f.Name())
		}
	}

	return nil
}

func PrintTree(path string, level int) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("%+v", err)
	}

	indent := ""
	if level > 0 {
		indent = strings.Repeat(" ", 4*level)
	}

	for _, f := range files {
		if f.Name()[0:1] == "." || f.Name()[0:1] == "_" {
			continue
		}
		if f.IsDir() {
			fmt.Printf("%s%s/\n", indent, f.Name())
			err = PrintTree(fmt.Sprintf("%s/%+v", path, f.Name()), level+1)
			if err != nil {
				return fmt.Errorf("%+v", err)
			}
		} else {
			fmt.Printf("%s%s\n", indent, f.Name())
		}
	}

	return nil
}
