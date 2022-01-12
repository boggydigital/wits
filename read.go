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

type SectLines map[string][]string
type SectMap map[string]map[string]string

func ReadSectLines(path string) (SectLines, error) {
	data := make(SectLines)

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
			tline := strings.TrimSpace(line)
			if strings.HasPrefix(tline, commentPrefix) {
				continue
			}
			data[sect] = append(data[sect], tline)
		} else {
			sect = line
			data[sect] = make([]string, 0)
		}
	}

	return data, nil
}

func ReadSectMap(path string) (SectMap, error) {
	if lines, err := ReadSectLines(path); err != nil {
		return nil, err
	} else {
		return SectLinesToMap(lines), nil
	}
}
