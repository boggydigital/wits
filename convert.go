package wits

import (
	"fmt"
	"strings"
)

const kvSep = "="

// FIXME: replace calls to this with strings.Cut when 1.18 releases
func cut(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

func SectLinesToMap(data SectLines) SectMap {
	res := make(SectMap, len(data))

	for sect, lines := range data {
		res[sect] = make(map[string]string, len(lines))
		for _, line := range lines {
			key, val, _ := cut(line, kvSep)
			res[sect][key] = val
		}
	}

	return res
}

func SectMapToLines(data SectMap) SectLines {
	res := make(SectLines, len(data))

	for sect, mp := range data {
		res[sect] = make([]string, 0, len(mp))
		for key, val := range mp {
			res[sect] = append(res[sect],
				fmt.Sprintf("%s%s%s", key, kvSep, val))
		}
	}

	return res
}
