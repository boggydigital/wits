package wits

import (
	"github.com/boggydigital/testo"
	"strconv"
	"testing"
)

func TestKvToKvs(t *testing.T) {
	tests := []struct {
		input  KeyValue
		output KeyValues
	}{
		{
			input:  nil,
			output: nil,
		},
		{
			input:  KeyValue{"k": ""},
			output: KeyValues{"k": {""}},
		},
		{
			input:  KeyValue{"k": "v"},
			output: KeyValues{"k": {"v"}},
		},
		{
			input:  KeyValue{"k1": "v1", "k2": "v2"},
			output: KeyValues{"k1": {"v1"}, "k2": {"v2"}},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			kvs := kvToKvs(tt.input)
			testo.DeepEqual(t, kvs, tt.output)
		})
	}
}

func TestKvsToKv(t *testing.T) {
	tests := []struct {
		input  KeyValues
		output KeyValue
	}{
		{
			input:  nil,
			output: nil,
		},
		{
			input:  KeyValues{"k": nil},
			output: KeyValue{"k": ""},
		},
		{
			input:  KeyValues{"k": {}},
			output: KeyValue{"k": ""},
		},
		{
			input:  KeyValues{"k": {""}},
			output: KeyValue{"k": ""},
		},
		{
			input:  KeyValues{"k": {"v"}},
			output: KeyValue{"k": "v"},
		},
		{
			// k1:v2 should be dropped
			input:  KeyValues{"k1": {"v1", "v2"}, "k2": {"v3"}},
			output: KeyValue{"k1": "v1", "k2": "v3"},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			kv := kvsToKv(tt.input)
			testo.DeepEqual(t, kv, tt.output)
		})
	}
}

func TestKvsToSkv(t *testing.T) {
	tests := []struct {
		input  KeyValues
		output SectionKeyValue
	}{
		{
			input:  nil,
			output: nil,
		},
		{
			input:  KeyValues{"s": nil},
			output: SectionKeyValue{"s": {}},
		},
		{
			input:  KeyValues{"s": {}},
			output: SectionKeyValue{"s": {}},
		},
		{
			input:  KeyValues{"s": {""}},
			output: SectionKeyValue{"s": {"": ""}},
		},
		{
			input:  KeyValues{"s": {"="}},
			output: SectionKeyValue{"s": {"": ""}},
		},
		{
			input:  KeyValues{"s": {"", ""}},
			output: SectionKeyValue{"s": {"": ""}},
		},
		{
			input:  KeyValues{"s": {"k"}},
			output: SectionKeyValue{"s": {"k": ""}},
		},
		{
			input:  KeyValues{"s": {"k="}},
			output: SectionKeyValue{"s": {"k": ""}},
		},
		{
			input:  KeyValues{"s": {"k=v"}},
			output: SectionKeyValue{"s": {"k": "v"}},
		},
		{
			//last value with the same key overwrites the previous ones
			input:  KeyValues{"s": {"=1", "=2"}},
			output: SectionKeyValue{"s": {"": "2"}},
		},
		{
			input:  KeyValues{"s": {"k1=v1", "k2=v2;v3"}},
			output: SectionKeyValue{"s": {"k1": "v1", "k2": "v2;v3"}},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			skv := kvsToSkv(tt.input)
			testo.DeepEqual(t, skv, tt.output)
		})
	}
}

func TestKvsToSkvs(t *testing.T) {
	tests := []struct {
		input  KeyValues
		output SectionKeyValues
	}{
		{
			input:  nil,
			output: nil,
		},
		{
			input:  KeyValues{"s": nil},
			output: SectionKeyValues{"s": {}},
		},
		{
			input:  KeyValues{"s": {}},
			output: SectionKeyValues{"s": {}},
		},
		{
			input:  KeyValues{"s": {""}},
			output: SectionKeyValues{"s": {"": {""}}},
		},
		{
			input:  KeyValues{"s": {"="}},
			output: SectionKeyValues{"s": {"": {""}}},
		},
		{
			input:  KeyValues{"s": {"=;"}},
			output: SectionKeyValues{"s": {"": {"", ""}}},
		},
		{
			input:  KeyValues{"s": {";="}},
			output: SectionKeyValues{"s": {";": {""}}},
		},
		{
			input:  KeyValues{"s": {"=;;"}},
			output: SectionKeyValues{"s": {"": {"", "", ""}}},
		},
		{
			//last value with the same key overwrites the previous ones
			input:  KeyValues{"s": {"=1", "=2;3"}},
			output: SectionKeyValues{"s": {"": {"2", "3"}}},
		},
		{
			input:  KeyValues{"s": {"k1=v1", "k2=v2;v3"}},
			output: SectionKeyValues{"s": {"k1": {"v1"}, "k2": {"v2", "v3"}}},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			skvs := kvsToSkvs(tt.input)
			testo.DeepEqual(t, skvs, tt.output)
		})
	}
}

func TestSkvToKvs(t *testing.T) {
	tests := []struct {
		input  SectionKeyValue
		output KeyValues
	}{
		{
			input:  nil,
			output: nil,
		},
		{
			input:  SectionKeyValue{"s": {}},
			output: KeyValues{"s": {}},
		},
		{
			input:  SectionKeyValue{"s": {"": ""}},
			output: KeyValues{"s": {"="}},
		},
		{
			input:  SectionKeyValue{"s": {"k": ""}},
			output: KeyValues{"s": {"k="}},
		},
		{
			input:  SectionKeyValue{"s": {"k": "v"}},
			output: KeyValues{"s": {"k=v"}},
		},
		{
			input:  SectionKeyValue{"s": {"k1": "v2;v3"}},
			output: KeyValues{"s": {"k1=v2;v3"}},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			skvs := skvToKvs(tt.input)
			testo.DeepEqual(t, skvs, tt.output)
		})
	}
}

func TestSkvsToKvs(t *testing.T) {
	tests := []struct {
		input  SectionKeyValues
		output KeyValues
	}{
		{
			input:  nil,
			output: nil,
		},
		{
			input:  SectionKeyValues{"s": {}},
			output: KeyValues{"s": {}},
		},
		{
			input:  SectionKeyValues{"s": {"": {""}}},
			output: KeyValues{"s": {"="}},
		},
		{
			input:  SectionKeyValues{"s": {"": {"", ""}}},
			output: KeyValues{"s": {"=;"}},
		},
		{
			input:  SectionKeyValues{"s": {";": {""}}},
			output: KeyValues{"s": {";="}},
		},
		{
			input:  SectionKeyValues{"s": {"": {"", "", ""}}},
			output: KeyValues{"s": {"=;;"}},
		},
		{
			input:  SectionKeyValues{"s": {"k1": {"v1", "v2"}}},
			output: KeyValues{"s": {"k1=v1;v2"}},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			kvs := skvsToKvs(tt.input)
			testo.DeepEqual(t, kvs, tt.output)
		})
	}
}
