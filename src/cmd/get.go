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

var getProjectName string

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&getProjectName, "project", "p", "", "Project to get data for")
}

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get data from Desktop or Downloads",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 3 {
			return fmt.Errorf("this command takes up to 3 arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			err := listFiles()
			if err != nil {
				return fmt.Errorf("failed to list files: %+v", err)
			}
			return nil
		}

		filePath := args[0]
		if num, err := strconv.Atoi(args[0]); err == nil {
			filePath, err = GetFileByIndex(num)
			if err != nil {
				return fmt.Errorf("failed to get file number %d: %+v", num, err)
			}
		}

		var err error
		if len(args) == 1 {
			err = get(filePath, "")
		} else {
			err = get(filePath, args[1])
		}
		if err != nil {
			return fmt.Errorf("failed to get file: %+v", err)
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
	cmd := fmt.Sprintf("mv \"%s\" \"%s\"", path, newPath)
	_, err := util.Exec(cmd)
	if err != nil {
		return fmt.Errorf("failed to execute '%+v': %+v", cmd, err)
	}

	projectPath := Config.Project
	if getProjectName != "" {
		projectPath = getProjectName
	}
	projectDepth := len(strings.Split(projectPath, "/"))
	pathDepth := strings.Repeat("../", projectDepth)
	markdownLink := fmt.Sprintf(fmt.Sprintf("![%s](%s%s)", filename, pathDepth, newPath))
	err = clipboard.WriteAll(markdownLink)
	if err == nil {
		fmt.Printf("Copied to Clipboard: ")
	} else {
		fmt.Printf("Failed to Copy to Clipboard\n")
	}
	fmt.Printf("%s\n", markdownLink)
	return nil
}

func listFiles() error {
	home := homeDir()
	desktopPath := fmt.Sprintf("%s/Desktop", home)
	downloadsPath := fmt.Sprintf("%s/Downloads", home)

	paths := []string{
		desktopPath,
		downloadsPath,
	}

	fileCount := 1

	for _, path := range paths {
		fmt.Printf("%+v\n", path)
		dirpath, _, filenames, err := util.GetFilesAndDirectories(path)
		if err != nil {
			return fmt.Errorf("failed to list files at '%+v': %+v", path, err)
		}
		sort.Strings(filenames)
		for _, filename := range filenames {
			if filename[0:1] == "." {
				continue
			}
			fmt.Printf("[%d] \"%s/%s\"\n", fileCount, dirpath, filename)
			fileCount = fileCount + 1
		}
		fmt.Println("")
	}
	return nil
}

func GetFileByIndex(index int) (string, error) {
	home := homeDir()
	desktopPath := fmt.Sprintf("%s/Desktop", home)
	downloadsPath := fmt.Sprintf("%s/Downloads", home)

	paths := []string{
		desktopPath,
		downloadsPath,
	}

	fileCount := 1

	for _, path := range paths {
		fmt.Printf("%+v\n", path)
		dirpath, _, filenames, err := util.GetFilesAndDirectories(path)
		if err != nil {
			return "", fmt.Errorf("failed to list files at '%+v': %+v", path, err)
		}
		sort.Strings(filenames)
		for _, filename := range filenames {
			if filename[0:1] == "." {
				continue
			}
			if fileCount == index {
				return fmt.Sprintf("%s/%s", dirpath, filename), nil
			}
			fileCount = fileCount + 1
		}
	}

	return "", fmt.Errorf("failed to locate file number %d", index)
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
