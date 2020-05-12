package util

import "os"

func NoteOrProjectExists(projectPath string) bool {
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		return false
	}
	return true
}
