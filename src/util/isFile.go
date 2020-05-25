package util

import (
	"fmt"
	"os"
	"path/filepath"
)

// IsFile ...
func IsFile(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, fmt.Errorf("%+v", err)
	}
	mode := fi.Mode()
	return !mode.IsDir(), nil
}

// IsDir ...
func IsDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, fmt.Errorf("%+v", err)
	}
	mode := fi.Mode()
	return mode.IsDir(), nil
}

// PathIsToFile ...
func PathIsToFile(path string) bool {
	ext := filepath.Ext(path)
	return ext != ""
}

// FileExists ...
func FileExists(path string) bool {
	// TODO
	return false
}

// DirectoryExists ...
func DirectoryExists(path string) bool {
	// TODO
	return false
}
