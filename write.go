package wits

import (
	"fmt"
	"os"
	"path/filepath"
)

func Write(data map[string][]string, path string) error {
	dir, _ := filepath.Split(path)
	if dir != "" {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
	}

	file, err := os.Create(path)
	defer file.Close()

	if err != nil {
		return err
	}

	for section, lines := range data {
		if _, err := fmt.Fprintf(file, "%s\n", section); err != nil {
			return err
		}
		for _, line := range lines {
			if _, err := fmt.Fprintf(file, "%s%s\n", spacePrefix, line); err != nil {
				return err
			}
		}
	}

	return nil
}