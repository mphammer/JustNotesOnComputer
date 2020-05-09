package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed opening file: %s", err)
	}
	dataScanner := bufio.NewScanner(file)
	dataScanner.Split(bufio.ScanLines)

	lines := []string{}
	for dataScanner.Scan() {
		txt := dataScanner.Text()
		lines = append(lines, txt+"\n")
	}

	file.Close()

	return lines, nil
}
