package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"SecondBrain/src/util"
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
	PreRun: func(cmd *cobra.Command, args []string) {
		LoadConfig()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			listFiles()
			return nil
		}

		filePath := args[0]
		if num, err := strconv.Atoi(args[0]); err == nil {
			filePath = GetFileByIndex(num)
		}

		var err error
		if len(args) == 1 {
			err = get(filePath, "")
		} else {
			err = get(filePath, args[1])
		}
		if err != nil {
			return fmt.Errorf("%+v", err)
		}
		return nil
	},
}

func get(path, newName string) error {
	id := util.GetID()

	basename := filepath.Base(path)
	extension := filepath.Ext(basename)
	name := strings.Replace(basename, extension, "", -1)
	if newName != "" {
		name = newName
	}

	filename := fmt.Sprintf("%s-%s%s", name, id, extension)
	newPath := fmt.Sprintf("_data/%s", filename)
	_, err := util.Exec(fmt.Sprintf("mv \"%s\" \"%s\"", path, newPath))
	if err != nil {
		return fmt.Errorf("%+v", err)
	}

	pathDepth := strings.Repeat("../", Config.ProjectDepth)
	markdownLink := fmt.Sprintf(fmt.Sprintf("![%s](%s%s)", filename, pathDepth, newPath))
	err = clipboard.WriteAll(markdownLink)
	if err == nil {
		fmt.Printf("Copied to Clipboard: ")
	}
	fmt.Printf("%s\n", markdownLink)
	return nil
}

func listFiles() {
	home := homeDir()
	desktopPath := fmt.Sprintf("%s/Desktop", home)
	downloadsPath := fmt.Sprintf("%s/Downloads", home)

	fileCount := 1

	fmt.Printf("Desktop:\n")
	dirpath, _, filenames, _ := util.GetFilesAndDirectories(desktopPath)
	sort.Strings(filenames)
	for _, filename := range filenames {
		if filename[0:1] == "." {
			continue
		}
		fmt.Printf("[%d] \"%s/%s\"\n", fileCount, dirpath, filename)
		fileCount = fileCount + 1
	}

	fmt.Println("")

	fmt.Printf("Downloads:\n")
	dirpath, _, filenames, _ = util.GetFilesAndDirectories(downloadsPath)
	sort.Strings(filenames)
	for _, filename := range filenames {
		if filename[0:1] == "." {
			continue
		}
		fmt.Printf("[%d] \"%s/%s\"\n", fileCount, dirpath, filename)
		fileCount = fileCount + 1
	}
}

func GetFileByIndex(index int) string {
	home := homeDir()
	desktopPath := fmt.Sprintf("%s/Desktop", home)
	downloadsPath := fmt.Sprintf("%s/Downloads", home)

	fileCount := 1

	dirpath, _, filenames, _ := util.GetFilesAndDirectories(desktopPath)
	sort.Strings(filenames)
	for _, filename := range filenames {
		if filename[0:1] == "." {
			continue
		}
		if fileCount == index {
			return fmt.Sprintf("%s/%s", dirpath, filename)
		}
		fileCount = fileCount + 1
	}

	dirpath, _, filenames, _ = util.GetFilesAndDirectories(downloadsPath)
	sort.Strings(filenames)
	for _, filename := range filenames {
		if filename[0:1] == "." {
			continue
		}
		if fileCount == index {
			return fmt.Sprintf("%s/%s", dirpath, filename)
		}
		fileCount = fileCount + 1
	}
	return ""
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
