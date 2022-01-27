# wits

wits is a trivial whitespace indented text structure of sections and lines:

```text
# comment1
section1
    # comment2
    line1-1
    line1-2
section2
    line2-1
    line2-2 
```

## Why wits?

Why you'd use wits vs other options? In most cases there are existing and better general purpose
text formats: JSON, YAML, TOML, etc. Yet, sometimes you don't need most of the functionality they
offer and might appreciate lighter solution for simple and specific cases.

wits work great when you need simple structural data (that's fixed and not generic, by design) and
you want to allow effortless human data edit or entry into those files, that's harder to mess up.

## Most common type: KeyValues

In essence, wits files contain lines and line prefix defines line role:

- `#` character prefix makes line a comment
- any amount of `(space)` or `(tab)` character prefixes (or combinations of those characters)  make
  line... a line (under a section)
- no prefix makes line a section

Structure described above is mapped to `map[string][]string` in Go - wits provide`KeyValues` type.

## Specialized types: from simpler to more complex than KeyValues

wits provide additional types for various needs:

### KeyValue

If your data needs are limited to a single value for a key, `KeyValue` might be better option
than `KeyValues`:

```text
key1
    value1
key2
    value2
    # the next value would be ignored when read as KeyValue
    value3
```

### SectionKeyValue

For a more complex data that looks like this:

```text
section1
    key1=value1
    key2=value2
section2
    key3=value3
    key4=value4
```

This structure would be represented as `map[string]map[string]string` in Go
or `wits.SectionKeyValue`.

### SectionKeyValues

The most complex data case would be similar to:

```text
section1
    key1=value1;value2;value3
    key2=value4
section2
    key3=value5;value6
    key4=value7;value8
```

This structure would be represented as `map[string]map[string][]string` in Go
or `wits.SectionKeyValues`.

## Using wits

Adding wits module to your Go app: `go get github.com/boggydigital/wits`

wits provide the following types:

- `KeyValue`: `map[string]string`
- `KeyValues`: `map[string][]string`
- `SectionKeyValue`: `map[string]KeyValue`: `map[string]map[string]string`
- `SectionKeyValues`: `map[string]KeyValues`: `map[string]map[string][]string`

All types can be read and written using the family of `Read*` functions e.g.:

- `ReadKeyValue(r io.Reader) (KeyValue, error)`
- `ReadKeyValues(r io.Reader) (KeyValues, error)`
- `ReadSectionKeyValue(r io.Reader) (SectionKeyValue, error)`
- `ReadSectionKeyValues(r io.Reader) (SectionKeyValues, error)`

Instances of every type can write data:

- `(kv KeyValue) Write(w io.Writer) error`
- `(kvs KeyValues) Write(w io.Writer) error`
- `(skv SectionKeyValue) Write(w io.Writer) error`
- `(skvs SectionKeyValues) Write(w io.Writer) error`
