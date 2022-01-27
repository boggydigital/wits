package wits

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	spacePfx   = " "
	tabPfx     = "\t"
	commentPfx = "#"
)

func ReadKeyValue(r io.Reader) (KeyValue, error) {
	if kvs, err := ReadKeyValues(r); err != nil {
		return nil, err
	} else {
		return kvsToKv(kvs), nil
	}
}

func ReadKeyValues(r io.Reader) (KeyValues, error) {
	data := make(KeyValues)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	sect := ""

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" ||
			strings.HasPrefix(line, commentPfx) {
			continue
		}

		if strings.HasPrefix(line, spacePfx) || strings.HasPrefix(line, tabPfx) {
			if sect == "" {
				return data, fmt.Errorf("cannot start with a whitespace prefixed line")
			}
			tline := strings.Trim(line, tabPfx+spacePfx)
			if strings.HasPrefix(tline, commentPfx) {
				continue
			}
			data[sect] = append(data[sect], tline)
		} else {
			sect = line
			if _, ok := data[sect]; ok {
				return data, fmt.Errorf("key or section %s is present more than once", sect)
			}
			data[sect] = make([]string, 0)
		}
	}

	return data, nil
}

func ReadSectionKeyValue(r io.Reader) (SectionKeyValue, error) {
	if lines, err := ReadKeyValues(r); err != nil {
		return nil, err
	} else {
		return kvsToSkv(lines), nil
	}
}

func ReadSectionKeyValues(r io.Reader) (SectionKeyValues, error) {
	if kvs, err := ReadKeyValues(r); err != nil {
		return nil, err
	} else {
		return kvsToSkvs(kvs), nil
	}
}
