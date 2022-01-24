package wits

type (
	KeyValue         map[string]string
	KeyValues        map[string][]string
	SectionKeyValue  map[string]KeyValue
	SectionKeyValues map[string]KeyValues
)
