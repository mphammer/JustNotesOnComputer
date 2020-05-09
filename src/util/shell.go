package util

import (
	"fmt"
	"os/exec"
)

func Exec(command string) (string, error) {
	// cmdValues := strings.Fields(command)
	cmd := exec.Command("sh", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf(string(out))
	}
	return string(out), nil
}
