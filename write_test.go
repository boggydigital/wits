package wits

import (
	"github.com/boggydigital/testo"
	"strconv"
	"testing"
)

type stringWriteCloser struct {
	content string
}

func (swc *stringWriteCloser) Write(b []byte) (int, error) {
	swc.content += string(b)
	return len(b), nil
}

func (swc *stringWriteCloser) Close() error {
	return nil
}

func TestKeyValuesWrite(t *testing.T) {
	tests := []struct {
		input  KeyValues
		output string
	}{
		{
			nil,
			"",
		},
		{
			KeyValues{},
			"",
		},
		{
			KeyValues{"": {""}},
			"\n \n",
		},
		{
			KeyValues{"": {"1"}},
			"\n 1\n",
		},
		{
			KeyValues{"1": {"2"}},
			"1\n 2\n",
		},
		{
			KeyValues{"1": {"2", "3"}},
			"1\n 2\n 3\n",
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			swc := &stringWriteCloser{}
			err := tt.input.Write(swc)

			testo.EqualValues(t, swc.content, tt.output)
			testo.Error(t, err, false)
		})
	}
}

func TestKeyValueWrite(t *testing.T) {
	tests := []struct {
		input  KeyValue
		output string
	}{
		// given that KeyValue.Write calls KeyValues.Write,
		// try to only add KeyValue specific test cases
		{
			KeyValue{"1": "2"},
			"1\n 2\n",
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			swc := &stringWriteCloser{}
			err := tt.input.Write(swc)

			testo.EqualValues(t, swc.content, tt.output)
			testo.Error(t, err, false)
		})
	}
}

func TestSectionKeyValueWrite(t *testing.T) {
	tests := []struct {
		input  SectionKeyValue
		output string
	}{
		// given that SectionKeyValue.Write calls KeyValues.Write,
		// try to only add SectionKeyValue specific test cases
		{
			SectionKeyValue{"1": nil},
			"1\n",
		},
		{
			SectionKeyValue{"1": {}},
			"1\n",
		},
		{
			SectionKeyValue{"1": {"": ""}},
			"1\n =\n",
		},
		{
			SectionKeyValue{"1": {"": "2"}},
			"1\n =2\n",
		},
		{
			SectionKeyValue{"1": {"2": "3"}},
			"1\n 2=3\n",
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			swc := &stringWriteCloser{}
			err := tt.input.Write(swc)

			testo.EqualValues(t, swc.content, tt.output)
			testo.Error(t, err, false)
		})
	}
}

func TestSectionKeyValuesWrite(t *testing.T) {
	tests := []struct {
		input  SectionKeyValues
		output string
	}{
		// given that SectionKeyValues.Write calls KeyValues.Write,
		// try to only add SectionKeyValues specific test cases
		{
			SectionKeyValues{"1": nil},
			"1\n",
		},
		{
			SectionKeyValues{"1": {}},
			"1\n",
		},
		{
			SectionKeyValues{"1": {"": nil}},
			"1\n =\n",
		},
		{
			SectionKeyValues{"1": {"": {}}},
			"1\n =\n",
		},
		{
			SectionKeyValues{"1": {"": {""}}},
			"1\n =\n",
		},
		{
			SectionKeyValues{"1": {"": {"2"}}},
			"1\n =2\n",
		},
		{
			SectionKeyValues{"1": {"2": {"3"}}},
			"1\n 2=3\n",
		},
		{
			SectionKeyValues{"1": {"2": {"3", "4"}}},
			"1\n 2=3;4\n",
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			swc := &stringWriteCloser{}
			err := tt.input.Write(swc)

			testo.EqualValues(t, swc.content, tt.output)
			testo.Error(t, err, false)
		})
	}
}
