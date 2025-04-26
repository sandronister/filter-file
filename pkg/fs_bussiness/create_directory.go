package fsbussiness

import "os"

func (m *model) CreateDirectory(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.MkdirAll(directory, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
