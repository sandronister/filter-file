package fsbussiness

import (
	"fmt"
	"os"
	"strings"
)

func (m *model) ListPDFInDirectory(directory string) ([]string, error) {
	if directory == "" {
		return nil, fmt.Errorf("directory path is empty")
	}

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory does not exist: %s", directory)
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".pdf") {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}
