package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"SecondBrain/src/util"
)

func init() {
	rootCmd.AddCommand(moveCmd)
}

var moveCmd = &cobra.Command{
	Use:     "move SOURCE TARGET",
	Example: "Rename a file: move path/name.md newpath/newname.md\nMove a directory: move path/directory path/newdirectory",
	Aliases: []string{"mv"},
	Short:   "Move or rename files",
	Long: `If target is a filename: Rename source file to new filename and path.
	If target is a directory: Move files from source directory to target directory.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("this command takes 2 arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		source := args[0]
		target := args[1]
		err := move(source, target)
		if err != nil {
			return fmt.Errorf("failed to move files: %+v", err)
		}
		return nil
	},
}

// TODO check if the current project was renamed

func move(startPath, destPath string) error {
	startPath = strings.Trim(startPath, "/")
	destPath = strings.Trim(destPath, "/")
	destFilename := ""

	dirpath, dirnames, filenames, err := util.GetFilesAndDirectories(startPath)
	if err != nil {
		return fmt.Errorf("failed to get files and directories: %+v", err)
	}
	if util.PathIsToFile(startPath) {
		dirpath = filepath.Dir(startPath)
		filenames = []string{filepath.Base(startPath)}
		dirnames = []string{}
		// handle new destination filename
		if util.PathIsToFile(destPath) {
			destPath = filepath.Dir(destPath)
			destFilename = filepath.Base(destPath)
			if !util.FileHasID(destFilename) {
				destFilename = util.AddFileID(destFilename)
			}
		}
	} else if util.PathIsToFile(destPath) {
		return fmt.Errorf("source is a directory - target cannot be a file")
	}

	for _, filename := range filenames {
		err = MoveFile(filename, dirpath, destPath, destFilename)
		if err != nil {
			return fmt.Errorf("failed to move file '%+v': %+v", filename, err)
		}
	}

	for _, dirname := range dirnames {
		err = move(fmt.Sprintf("%s/%s", startPath, dirname), fmt.Sprintf("%s/%s", destPath, dirname))
		if err != nil {
			return fmt.Errorf("failed to move subdirectory: %+v", err)
		}
	}

	return nil

}

func MoveFile(filename, dirpath, destPath, destFilename string) error {
	oldPath := fmt.Sprintf("%s/%s", dirpath, filename)
	newPath := fmt.Sprintf("%s/%s", destPath, filename)
	if destFilename != "" {
		newPath = fmt.Sprintf("%s/%s", destPath, destFilename)
	}

	// Change all references to this file
	cmd := fmt.Sprintf("find . -type f -name \"*\\.md\" -print0 | xargs -0 sed -i '' -e 's~%s~%s~g'", oldPath, newPath)
	// _, err := util.Exec(cmd)
	// if err != nil {
	// 	return fmt.Errorf("failed to execute '%+v': %+v", cmd, err)
	// }
	fmt.Printf("%+v\n", cmd)

	// Change references within this file
	oldDepth := len(strings.Split(oldPath, "/")) - 1
	oldRootPath := strings.Repeat("../", oldDepth)
	newDepth := len(strings.Split(newPath, "/")) - 1
	newRootPath := strings.Repeat("../", newDepth)

	// lines, err := util.ReadFileLines(oldPath)
	// if err != nil {
	// 	return fmt.Errorf("failed read old file: %s", err)
	// }
	// currFile, err := os.Open(oldPath)
	// if err != nil {
	// 	return fmt.Errorf("failed create file: %s", err)
	// }
	// regOldRootPath := strings.Replace(oldRootPath, ".", "\\.", -1)
	// re := regexp.MustCompile(regOldRootPath)
	// for _, txt := range lines {
	// 	txt = string(re.ReplaceAll([]byte(txt), []byte(newRootPath)))
	// 	fmt.Fprintln(currFile, txt)
	// }
	// currFile.Close()

	fmt.Printf("%s -> %s\n", oldRootPath, newRootPath)

	// Create new directories
	cmd = fmt.Sprintf("mkdir -p %s", destPath)
	// _, err = util.Exec(cmd)
	// if err != nil {
	// 	return fmt.Errorf("failed to execute '%+v': %+v", cmd, err)
	// }
	fmt.Printf("%+v\n", cmd)

	// Move the file
	cmd = fmt.Sprintf("mv %s %s", oldPath, newPath)
	// _, err = util.Exec(cmd)
	// if err != nil {
	// 	return fmt.Errorf("failed to execute '%+v': %+v", cmd, err)
	// }
	fmt.Printf("%+v\n", cmd)
	fmt.Println("")
	return nil
}
