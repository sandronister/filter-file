package inputclear

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInputText(label string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label)
	search, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	cleaned := strings.TrimSpace(search)

	if cleaned == "" {
		return "", fmt.Errorf("input cannot be empty")
	}
	return cleaned, nil
}
