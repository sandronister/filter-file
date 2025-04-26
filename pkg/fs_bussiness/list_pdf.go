package fsbussiness

import (
	"os"
	"strings"
)

func (m *model) ListPDFInDirectory(directory string) ([]string, error) {
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
