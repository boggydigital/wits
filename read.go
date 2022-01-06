package wits

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	spacePrefix   = " "
	commentPrefix = "#"
)

func Read(path string) (map[string][]string, error) {
	data := make(map[string][]string)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return data, nil
	}

	listFile, err := os.Open(path)
	if err != nil {
		return data, err
	}

	scanner := bufio.NewScanner(listFile)
	scanner.Split(bufio.ScanLines)

	sect := ""

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" ||
			strings.HasPrefix(line, commentPrefix) {
			continue
		}

		if strings.HasPrefix(line, spacePrefix) {
			if sect == "" {
				return data, fmt.Errorf("cannot start with a space prefixed line")
			}
			data[sect] = append(data[sect], strings.TrimSpace(line))
		} else {
			sect = line
			data[sect] = make([]string, 0)
		}
	}

	return data, nil
}
