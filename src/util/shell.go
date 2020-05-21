package util

import (
	"fmt"
	"os"
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

func ExecShell(command string) error {
	// cmdValues := strings.Fields(command)
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%+v", err)
	}
	return nil
}
