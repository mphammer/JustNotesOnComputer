package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(inputText string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(inputText)
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}
