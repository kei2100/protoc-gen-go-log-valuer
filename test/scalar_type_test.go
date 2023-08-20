package test

import (
	"testing"
)

func TestScalarType_LogValue(t *testing.T) {
	m := &ScalarType{
		SecretVal:   "secret",
		DoubleVal:   0.1,
		FloatVal:    0.1,
		Int32Val:    1,
		Int64Val:    1,
		Uint32Val:   1,
		Uint64Val:   1,
		Sint32Val:   2,
		Sint64Val:   2,
		Fixed32Val:  2,
		Fixed64Val:  2,
		Sfixed32Val: 3,
		Sfixed64Val: 3,
		BoolVal:     true,
		StringVal:   "string",
		BytesVal:    []byte{1, 2, 3}, // expect base64 encode
		EnumVal:     ScalarType_ENUM_1,
		OneofVal:    &ScalarType_OneofStringVal{OneofStringVal: "oneof_string"},
		XString:     "_string",
	}
	want := `
{
	"secret_val": "[REDACTED]",
	"double_val": 0.1,
	"float_val": 0.1,
	"int32_val": 1,
	"int64_val": 1,
	"uint32_val": 1,
	"uint64_val": 1,
	"sint32_val": 2,
	"sint64_val": 2,
	"fixed32_val": 2,
	"fixed64_val": 2,
	"sfixed32_val": 3,
	"sfixed64_val": 3,
	"bool_val": true,
	"string_val": "string",
	"bytes_val": "AQID",
	"enum_val": "ENUM_1",
	"oneof_string_val": "oneof_string",
	"_String": "_string"
}
`
	testJSON(t, m, want)
}
