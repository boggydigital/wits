package wits

import (
	"fmt"
	"strings"
)

const (
	keyValueSep = "="
	valuesSep   = ";"
)

// FIXME: replace calls to this with strings.Cut when 1.18 releases
func cut(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

// The following conversions are available
// (left column = from, top row = to):
// |      | kv | kvs | skv | skvs |
// |   kv | == | YES | NAH | NOPE |
// |  kvs | OK | ==  | YES | YEAH |
// |  skv | NO | YES | === | NOPE |
// | skvs | NO | YES | NAH | ==== |
//
// Simple way to read this table:
// - KeyValues can be converted into everything else, it's the only type that's read natively
// - everything can be converted into KeyValues,  it's the only type that's writable natively

func kvToKvs(kv KeyValue) KeyValues {
	kvs := make(KeyValues, len(kv))

	for key, value := range kv {
		kvs[key] = []string{value}
	}

	return kvs
}

func kvsToKv(kvs KeyValues) KeyValue {
	kv := make(KeyValue, len(kvs))

	for key, values := range kvs {
		switch len(values) {
		case 0:
			kv[key] = ""
		case 1:
			kv[key] = values[0]
		}
	}

	return kv
}

func kvsToSkv(kvs KeyValues) SectionKeyValue {
	skv := make(SectionKeyValue, len(kvs))

	for section, keyValues := range kvs {
		skv[section] = make(KeyValue, len(keyValues))
		for _, keyValue := range keyValues {
			key, val, _ := cut(keyValue, keyValueSep)
			skv[section][key] = val
		}
	}

	return skv
}

func kvsToSkvs(kvs KeyValues) SectionKeyValues {
	skvs := make(SectionKeyValues, len(kvs))

	for section, keyValuePairs := range kvs {
		skvs[section] = make(KeyValues, len(keyValuePairs))
		for _, keyValues := range keyValuePairs {
			key, values, _ := cut(keyValues, keyValueSep)
			skvs[section][key] = strings.Split(values, valuesSep)
		}
	}

	return skvs
}

func skvToKvs(skv SectionKeyValue) KeyValues {
	kvs := make(KeyValues, len(skv))

	for section, kv := range skv {
		kvs[section] = make([]string, 0, len(kv))
		for key, val := range kv {
			kvs[section] = append(kvs[section],
				fmt.Sprintf("%s%s%s", key, keyValueSep, val))
		}
	}

	return kvs
}

func skvsToKvs(skvs SectionKeyValues) KeyValues {
	kvs := make(KeyValues, len(skvs))

	for section, keyValues := range skvs {
		for key, values := range keyValues {
			kvs[section] = append(kvs[section],
				fmt.Sprintf("%s%s%s",
					key,
					keyValueSep,
					strings.Join(values, valuesSep)))
		}
	}

	return kvs
}
