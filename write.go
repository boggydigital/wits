package wits

import (
	"fmt"
	"io"
)

func (kvs KeyValues) Write(w io.Writer) error {
	for key, values := range kvs {
		if _, err := fmt.Fprintf(w, "%s\n", key); err != nil {
			return err
		}
		for _, value := range values {
			if _, err := fmt.Fprintf(w, "%s%s\n", spacePfx, value); err != nil {
				return err
			}
		}
	}

	return nil
}

func (kv KeyValue) Write(w io.Writer) error {
	return kvToKvs(kv).Write(w)
}

func (skv SectionKeyValue) Write(w io.Writer) error {
	return skvToKvs(skv).Write(w)
}

func (skvs SectionKeyValues) Write(w io.Writer) error {
	return skvsToKvs(skvs).Write(w)
}
