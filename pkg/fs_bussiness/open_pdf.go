package fsbussiness

import (
	"fmt"

	"github.com/ledongthuc/pdf"
)

func (m *model) OpenPDF(filePath string) (string, error) {
	file, res, err := pdf.Open(filePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	var content string

	for i := 0; i < res.NumPage(); i++ {
		page := res.Page(i + 1)

		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			fmt.Printf("Warning: failed to extract text from page %d: %v\n", i+1, filePath)
			continue
		}

		content += text
	}

	if content == "" {
		return "", fmt.Errorf("no text extracted from PDF: %s", filePath)
	}

	return content, nil
}
