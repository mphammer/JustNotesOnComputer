package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"SecondBrain/_tools/util"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get files from Desktop or Downloads",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 3 {
			return fmt.Errorf("this command takes up to 3 arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			listFiles()
		} else if len(args) == 1 {
			get(args[0], "")
		} else {
			get(args[0], args[1])
		}
		return nil
	},
}

func get(path, newName string) {
	id := util.GetID()

	basename := filepath.Base(path)
	extension := filepath.Ext(basename)
	name := strings.Replace(basename, extension, "", -1)
	if newName != "" {
		name = newName
	}

	filename := fmt.Sprintf("%s-%s%s", name, id, extension)
	newPath := fmt.Sprintf("_data/%s", filename)
	util.Exec(fmt.Sprintf("mv \"%s\" \"%s\"", path, newPath))
	markdownLink := fmt.Sprintf(fmt.Sprintf("![%s](../%s)", filename, newPath))
	err := clipboard.WriteAll(markdownLink)
	if err == nil {
		fmt.Printf("Copied to Clipboard: ")
	}
	fmt.Printf("%s\n", markdownLink)
}

func listFiles() {
	home := homeDir()
	desktopPath := fmt.Sprintf("%s/Desktop", home)
	downloadsPath := fmt.Sprintf("%s/Downloads", home)

	fmt.Printf("Desktop:\n")
	dirpath, _, filenames, _ := util.GetFilesAndDirectories(desktopPath)
	sort.Strings(filenames)
	for _, filename := range filenames {
		if filename[0:1] == "." {
			continue
		}
		fmt.Printf("\"%s/%s\"\n", dirpath, filename)
	}

	fmt.Println("")

	fmt.Printf("Downloads:\n")
	dirpath, _, filenames, _ = util.GetFilesAndDirectories(downloadsPath)
	sort.Strings(filenames)
	for _, filename := range filenames {
		if filename[0:1] == "." {
			continue
		}
		fmt.Printf("\"%s/%s\"\n", dirpath, filename)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
