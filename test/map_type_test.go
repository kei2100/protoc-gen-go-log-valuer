package test

import (
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func TestMapType_LogValue(t *testing.T) {
	m := &MapType{
		StringMap: map[string]string{
			"a": "aa",
		},
		FloatMap: map[string]float32{
			"a": 0.1,
		},
		IntKeyMap: map[int32]string{
			0: "zero",
		},
		MessageMap: map[string]*MapType_MapMessage{
			"a": {
				SecretVal: "secret",
				DoubleVal: 0.2,
				FloatVal:  0.1,
			},
		},
		StructMap: map[string]*structpb.Struct{
			"a": {
				Fields: map[string]*structpb.Value{
					"struct_string": {
						Kind: &structpb.Value_StringValue{StringValue: "foo"},
					},
					"struct_number": {
						Kind: &structpb.Value_NumberValue{NumberValue: 0.1},
					},
					"struct_bool": {
						Kind: &structpb.Value_BoolValue{BoolValue: true},
					},
				},
			},
		},
		EmptyStringMap:  map[string]string{},
		EmptyMessageMap: map[string]*MapType_MapMessage{},
		NullStringMap:   nil,
		NullMessageMap:  nil,
	}

	want := `
{
	"string_map": {
		"a": "aa"
	},
	"float_map": {
		"a": 0.1
	},
	"int_key_map": {
		"0": "zero"
	},
	"message_map": {
		"a": {
			"secret_val": "[REDACTED]",
			"double_val": 0.2,
			"float_val": 0.1
		}
	},
	"struct_map": {
		"a": {
			"struct_string": "foo",
			"struct_number": 0.1,
			"struct_bool": true
		}
	},
	"empty_string_map": {},
	"empty_message_map": {},
	"null_string_map": {},
	"null_message_map": {}
}`
	testJSON(t, m, want)
}
