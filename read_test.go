package wits

import (
	"github.com/boggydigital/testo"
	"strconv"
	"strings"
	"testing"
)

func TestReadKeyValue(t *testing.T) {
	tests := []struct {
		content string
		output  KeyValue
		expErr  bool
	}{
		{
			"",
			KeyValue{},
			false,
		},
		{
			// document cannot start with a whitespace
			" ",
			nil,
			true,
		},
		{
			// document cannot start with a whitespace (including a comment before it)
			"#\n ",
			nil,
			true,
		},
		{
			"1",
			KeyValue{"1": ""},
			false,
		},
		{
			"1\n" +
				" 2",
			KeyValue{"1": "2"},
			false,
		},
		{
			// ignore second value under the same key
			"1\n" +
				" 2\n" +
				" 3",
			KeyValue{"1": "2"},
			false,
		},
		{
			"1\n" +
				"  2",
			KeyValue{"1": "2"},
			false,
		},
		{
			"1\n" +
				"\t2",
			KeyValue{"1": "2"},
			false,
		},
		{
			"1\n" +
				"\t\t2",
			KeyValue{"1": "2"},
			false,
		},
		{
			"1\n" +
				" \t2",
			KeyValue{"1": "2"},
			false,
		},
		{
			"1\n" +
				" \t 2",
			KeyValue{"1": "2"},
			false,
		},
		{
			"1\n" +
				"\t 2",
			KeyValue{"1": "2"},
			false,
		},
		{
			"1\n" +
				"\t \t2",
			KeyValue{"1": "2"},
			false,
		},
		{
			"1\n" +
				"# comment\n" +
				" 2",
			KeyValue{"1": "2"},
			false,
		},
		{
			"1\n" +
				" #comment\n" +
				" 2",
			KeyValue{"1": "2"},
			false,
		},
		{
			"1\n" +
				"\t#comment\n" +
				"\t2",
			KeyValue{"1": "2"},
			false,
		},
		{
			"1\n" +
				" 2\n" +
				"3\n" +
				" 4",
			KeyValue{"1": "2", "3": "4"},
			false,
		},
		{
			// duplicate keys/sections produce an error
			"1\n" +
				"2\n" +
				"1",
			nil,
			true,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			kv, err := ReadKeyValue(strings.NewReader(tt.content))

			testo.EqualInterfaces(t, kv, tt.output)
			testo.Error(t, err, tt.expErr)
		})
	}
}

func TestReadKeyValues(t *testing.T) {
	tests := []struct {
		content string
		output  KeyValues
		expErr  bool
	}{
		// given that ReadKeyValues calls ReadKeyValue,
		// try to only add KeyValues specific test cases
		{
			"1\n" +
				" 2\n" +
				" 3",
			KeyValues{"1": {"2", "3"}},
			false,
		},
		{
			// duplicate keys/sections produce an error
			"1\n2\n1",
			KeyValues{"1": {}, "2": {}},
			true,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			kvs, err := ReadKeyValues(strings.NewReader(tt.content))

			testo.EqualInterfaces(t, kvs, tt.output)
			testo.Error(t, err, tt.expErr)
		})
	}
}

func TestReadSectionKeyValue(t *testing.T) {
	tests := []struct {
		content string
		output  SectionKeyValue
		expErr  bool
	}{
		// given that ReadSectionKeyValue calls ReadKeyValue,
		// try to only add SectionKeyValue specific test cases
		{
			"=\n",
			SectionKeyValue{"=": {}},
			false,
		},
		{
			"\n" +
				"=",
			SectionKeyValue{"=": {}},
			false,
		},
		{
			"1\n" +
				" 2=",
			SectionKeyValue{"1": {"2": ""}},
			false,
		},
		{
			"1\n" +
				" =2",
			SectionKeyValue{"1": {"": "2"}},
			false,
		},
		{
			"1\n" +
				" 2=3",
			SectionKeyValue{"1": {"2": "3"}},
			false,
		},
		{
			"1\n" +
				" 2=3;4",
			SectionKeyValue{"1": {"2": "3;4"}},
			false,
		},
		{
			// duplicate keys/sections produce an error
			"1\n2\n1",
			nil,
			true,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			skv, err := ReadSectionKeyValue(strings.NewReader(tt.content))

			testo.EqualInterfaces(t, skv, tt.output)
			testo.Error(t, err, tt.expErr)
		})
	}
}

func TestReadSectionKeyValues(t *testing.T) {
	tests := []struct {
		content string
		output  SectionKeyValues
		expErr  bool
	}{
		// given that ReadSectionKeyValues calls ReadKeyValue,
		// try to only add SectionKeyValues specific test cases
		{
			"1\n" +
				" 2=3;4",
			SectionKeyValues{"1": {"2": {"3", "4"}}},
			false,
		},
		{
			"1\n" +
				"2\n" +
				"1",
			nil,
			true,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			skvs, err := ReadSectionKeyValues(strings.NewReader(tt.content))

			testo.EqualInterfaces(t, skvs, tt.output)
			testo.Error(t, err, tt.expErr)
		})
	}
}
