package util

import (
	"fmt"
	"os/exec"
	"strings"
)

func Exec(command string) (string, error) {
	cmdValues := strings.Fields(command)
	cmd := exec.Command(cmdValues[0], cmdValues[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%+v", out)
	}
	return string(out), nil
}
