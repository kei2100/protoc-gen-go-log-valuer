package test

import (
	"testing"
)

func TestListType_LogValue(t *testing.T) {
	m := &ListType{
		StringList: []string{"a", "b"},
	}
	want := `
{
	"string_list": {
		"0": "a",
		"1": "b"
	}
}
`
	testJSON(t, m, want)
}
