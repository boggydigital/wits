# wits

wits is a trivial whitespace indented text structure of sections and lines:

```text
#comment1
section1
 line1-1
 line1-2
#comment2
section2
 line2-1
 line2-2 
```

In essence, wits files contain lines and line prefix defines line role:

- `#` character prefix makes line a comment
- `(space)` character prefix makes line... a line (under a section)
- no prefix makes line a section

Structure described above is mapped to `map[string][]string` in Go - we'll call it `SectLines` going
forward.

This project is a Go language module that provides helpers to work with wits files.

Lines can be used to represent key value pairs and wits provides helpers to work with key value
pairs:

```text
section
    key1=value1
    key2=value2
```

That structure is mapped to `map[string]map[string]string` in Go - we'll call it `SectMap` going
forward.

## Using wits

Adding wits module to your Go app: `go get github.com/boggydigital/wits`

wits provides the following methods to read data:

- `ReadSectLines(path string) (SectLines, error)`
- `ReadSectMap(path string) (SectMap, error)`

wits provides the following methods to write data:

- `(sl SectLines) Write(path string) error`
- `(sm SectMap) Write(path string) error`
