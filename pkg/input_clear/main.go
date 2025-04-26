package inputclear

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func GetInputText(label string) (string, error) {
	re := regexp.MustCompile(`[^\w\s]`)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label)
	search, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	cleaned := re.ReplaceAllString(search, "")
	cleaned = strings.TrimSpace(cleaned)

	if cleaned == "" {
		return "", fmt.Errorf("input cannot be empty")
	}
	return cleaned, nil
}
