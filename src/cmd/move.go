package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"SecondBrain/src/util"
)

func init() {
	rootCmd.AddCommand(moveCmd)
}

var moveCmd = &cobra.Command{
	Use: "move SOURCE TARGET",
	Example: `Move a file:
$ ./jason move <path/name.md> <newpath>

Rename a file:
$ ./jason move <path/name-ID.md> <newpath/newname>       // a new ID will be added
$ ./jason move <path/name-ID.md> <newpath/newname-ID.md> // provided ID is used

Move and rename a file:
$ ./jason move <path/name.md> <newpath/newname-ID.md>

Move a directory:
$ ./jason move <directorypath/directoryname> <newdirectorypath/directoryname>

Rename a directory:
$ ./jason move <directorypath/directoryname> <directorypath/newdirectoryname>

Move and rename a directory:
$ ./jason move <directorypath/directoryname> <newdirectorypath/newdirectoryname>`,
	Aliases: []string{"mv"},
	Short:   "Move and/or rename Notes and Projects",
	Long: `Move and/or rename Notes and Projects:

SOURCE: Path to a file or project.
TARGET: Path to a file or project.

- If SOURCE is a file: 
  - If TARGET is a file: Move SOURCE to new path and rename it to the TARGET filename.
  - If TARGET is a project: Move SOURCE to new path.
- If SOURCE is a project: 
  - If TARGET name is the same: Move SOURCE to new path.
  - If TARGET name is different: Move SOURCE to new path and rename it to the TARGET name.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("this command takes 2 arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		sourcePath := strings.Trim(args[0], "/")
		targetPath := strings.Trim(args[1], "/")

		var err error
		if util.PathIsToFile(sourcePath) {
			err = MoveFile(sourcePath, targetPath)
			if err != nil {
				return fmt.Errorf("failed to move files: %+v", err)
			}
		} else {
			err = MoveDirectory(sourcePath, targetPath)
			if err != nil {
				return fmt.Errorf("failed to move files: %+v", err)
			}
			err = UpdateCurrProjectPath(sourcePath, targetPath)
			if err != nil {
				return fmt.Errorf("failed to update current project path: %+v", err)
			}
		}
		return nil
	},
}

func MoveFile(sourceFilePath, targetPath string) error {
	targetDirPath := targetPath
	targetFilename := ""

	sourceDirPath := filepath.Dir(sourceFilePath)
	sourceFilename := filepath.Base(sourceFilePath)
	if util.PathIsToFile(targetPath) {
		// User provided new name for file
		targetDirPath = filepath.Dir(targetPath)
		targetFilename = filepath.Base(targetPath)
		if !util.FileHasID(targetFilename) {
			targetFilename = util.AddFileID(targetFilename)
		}
	}
	err := MoveFileAndUpdateReferences(sourceDirPath, sourceFilename, targetDirPath, targetFilename)
	if err != nil {
		return fmt.Errorf("failed to move file '%+v': %+v", sourceFilename, err)
	}

	return nil
}

func MoveDirectory(sourceDirPath, targetDirPath string) error {
	if util.PathIsToFile(targetDirPath) {
		return fmt.Errorf("source is a directory - target cannot be a file")
	}

	currDirPath, dirnames, filenames, err := util.GetFilesAndDirectories(sourceDirPath)
	if err != nil {
		return fmt.Errorf("failed to get files and directories: %+v", err)
	}

	for _, filename := range filenames {
		err = MoveFileAndUpdateReferences(currDirPath, filename, targetDirPath, filename)
		if err != nil {
			return fmt.Errorf("failed to move file '%+v': %+v", filename, err)
		}
	}

	for _, dirname := range dirnames {
		nextSourceDirPath := fmt.Sprintf("%s/%s", sourceDirPath, dirname)
		nextTargetDirPath := fmt.Sprintf("%s/%s", targetDirPath, dirname)
		err = MoveDirectory(nextSourceDirPath, nextTargetDirPath)
		if err != nil {
			return fmt.Errorf("failed to move subdirectory: %+v", err)
		}
	}
	// Remove source directory if empty
	_, dirnames, filenames, err = util.GetFilesAndDirectories(sourceDirPath)
	if err != nil {
		return fmt.Errorf("failed to get files and directories: %+v", err)
	}
	if len(dirnames) == 0 && len(filenames) == 0 {
		cmd := fmt.Sprintf("rmdir %s", sourceDirPath)
		_, err = util.Exec(cmd)
		if err != nil {
			return fmt.Errorf("failed to execute '%+v': %+v", cmd, err)
		}
		// fmt.Printf("%+v\n", cmd)
	}
	return nil
}

func MoveFileAndUpdateReferences(oldDirPath, oldFilename, newDirPath, newFilename string) error {
	oldFilePath := fmt.Sprintf("%s/%s", oldDirPath, oldFilename)
	var newFilePath string
	if newFilename != "" {
		newFilePath = fmt.Sprintf("%s/%s", newDirPath, newFilename)
	} else {
		newFilePath = fmt.Sprintf("%s/%s", newDirPath, oldFilename)
	}

	// Change all references to this file
	cmd := fmt.Sprintf("find . -type f -name \"*\\.md\" -print0 | xargs -0 sed -i '' -e 's~%s~%s~g'", oldFilePath, newFilePath)
	_, err := util.Exec(cmd)
	if err != nil {
		return fmt.Errorf("failed to execute '%+v': %+v", cmd, err)
	}
	// fmt.Printf("%+v\n", cmd)

	// Change references within this file
	oldDepth := len(strings.Split(oldDirPath, "/"))
	oldPathToRoot := strings.Repeat("../", oldDepth)
	newDepth := len(strings.Split(newDirPath, "/"))
	newPathToRoot := strings.Repeat("../", newDepth)

	lines, err := util.ReadFileLines(oldFilePath)
	if err != nil {
		return fmt.Errorf("failed read old file: %s", err)
	}
	currFile, err := os.Open(oldFilePath)
	if err != nil {
		return fmt.Errorf("failed create file: %s", err)
	}
	regOldPathToRoot := strings.Replace(oldPathToRoot, ".", "\\.", -1)
	re := regexp.MustCompile(regOldPathToRoot)
	for _, txt := range lines {
		txt = string(re.ReplaceAll([]byte(txt), []byte(newPathToRoot)))
		fmt.Fprintln(currFile, txt)
	}
	currFile.Close()

	// fmt.Printf("%s -> %s\n", oldPathToRoot, newPathToRoot)

	// Create new directories
	cmd = fmt.Sprintf("mkdir -p %s", newDirPath)
	_, err = util.Exec(cmd)
	if err != nil {
		return fmt.Errorf("failed to execute '%+v': %+v", cmd, err)
	}
	// fmt.Printf("%+v\n", cmd)

	// Move the file
	cmd = fmt.Sprintf("mv %s %s", oldFilePath, newFilePath)
	_, err = util.Exec(cmd)
	if err != nil {
		return fmt.Errorf("failed to execute '%+v': %+v", cmd, err)
	}
	// fmt.Printf("%+v\n", cmd)
	// fmt.Println("")
	return nil
}

func UpdateCurrProjectPath(oldUpdatedPath, newUpdatedPath string) error {
	oldItemsFromUpdatedPath := strings.Split(oldUpdatedPath, "/")
	newItemsFromUpdatedPath := strings.Split(newUpdatedPath, "/")

	currProjectPath := Config.Project
	itemsFromCurrProjectPath := strings.Split(currProjectPath, "/")

	if PathIsParentOrSame(oldUpdatedPath, currProjectPath) {
		i := 0
		for i < len(oldItemsFromUpdatedPath) {
			if itemsFromCurrProjectPath[i] != oldItemsFromUpdatedPath[i] {
				break
			}
			i = i + 1
		}
		remainingCurrProjectPathItems := itemsFromCurrProjectPath[i:]
		newProjectPathItems := append(newItemsFromUpdatedPath, remainingCurrProjectPathItems...)
		newProjectPath := ""
		for _, elem := range newProjectPathItems {
			newProjectPath = fmt.Sprintf("%s/%s", newProjectPath, elem)
		}
		Config.Project = strings.Trim(newProjectPath, "/")
		Config.ProjectDepth = len(newItemsFromUpdatedPath) + len(remainingCurrProjectPathItems)
		err := SaveConfig()
		if err != nil {
			return fmt.Errorf("failed to save config file: %+v", err)
		}
		fmt.Printf("Updated Current Project: %s\n", Config.Project)
	}
	return nil
}

func PathIsParentOrSame(parentPath, childPath string) bool {
	parentPathItems := strings.Split(parentPath, "/")
	childPathItems := strings.Split(childPath, "/")

	if len(parentPathItems) > len(childPathItems) {
		return false
	}
	if parentPathItems[0] == childPathItems[0] {
		return true
	}
	return false
}
