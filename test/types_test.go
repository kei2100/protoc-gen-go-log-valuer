package test

import (
	"google.golang.org/protobuf/types/known/structpb"
	"log/slog"
	"testing"
)

func TestTypes_LogValue(t *testing.T) {
	m := &Types{
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
		BytesVal:    []byte{1, 2, 3},
		EnumVal:     Types_ENUM_1,
		OtherTypeVal: &OtherType1{
			OtherStringVal: "other_string",
			OtherSecretVal: "other_secret",
		},
		NestedTypeVal: &Types_NestedType{
			NestedStringVal: "nested_string",
			NestedSecretVal: "nested_secret",
		},
		OtherTypeNestedTypeVal: &OtherType2_NestedType{
			NestedStringVal: "other_nested_string",
			NestedSecretVal: "other_nested_secret",
		},
		OneofVal: &Types_OneofStringVal{
			OneofStringVal: "", // set zero value explicitly
		},
		MapVal1: map[string]string{
			"foo": "bar",
		},
		MapVal2: map[string]*OtherType1{
			"foo": {
				OtherStringVal: "other_string",
				OtherSecretVal: "other_secret",
			},
		},
		RepeatedVal1: []string{
			"foo", "bar",
		},
		RepeatedVal2: []Types_Enum{
			Types_ENUM_1, Types_ENUM_2,
		},
		RepeatedVal3: []*OtherType1{
			{
				OtherStringVal: "other_string",
				OtherSecretVal: "other_secret",
			},
		},
		RepeatedEmptyVal: []string{},
		StructVal: &structpb.Struct{
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
		XString:                   "foo",
		OptionalVal:               refs("foo"),
		OptionalNotPresentVal:     nil,
		OptionalEnum:              refs(Types_ENUM_1),
		OptionalNotPresentEnum:    nil,
		OptionalMessage:           &OtherType3{Val: "foo"},
		OptionalNotPresentMessage: nil,
		PresentMessage:            &OtherType3{Val: "foo"},
		NotPresentMessage:         nil,
	}
	got := m.LogValue().Group()
	want := []slog.Attr{
		{Key: "secret_val", Value: slog.StringValue("[REDACTED]")},
		{Key: "double_val", Value: slog.Float64Value(0.1)},
		{Key: "float_val", Value: slog.Float64Value(float64(float32(0.1)))},
		{Key: "int32_val", Value: slog.Int64Value(1)},
		{Key: "int64_val", Value: slog.Int64Value(1)},
		{Key: "uint32_val", Value: slog.Uint64Value(1)},
		{Key: "uint64_val", Value: slog.Uint64Value(1)},
		{Key: "sint32_val", Value: slog.Int64Value(2)},
		{Key: "sint64_val", Value: slog.Int64Value(2)},
		{Key: "fixed32_val", Value: slog.Uint64Value(2)},
		{Key: "fixed64_val", Value: slog.Uint64Value(2)},
		{Key: "sfixed32_val", Value: slog.Int64Value(3)},
		{Key: "sfixed64_val", Value: slog.Int64Value(3)},
		{Key: "bool_val", Value: slog.BoolValue(true)},
		{Key: "string_val", Value: slog.StringValue("string")},
		{Key: "bytes_val", Value: slog.AnyValue([]byte{1, 2, 3})},
		{Key: "enum_val", Value: slog.StringValue(Types_ENUM_1.String())},
		{Key: "other_type_val", Value: slog.GroupValue(
			slog.String("other_string_val", "other_string"),
			slog.String("other_secret_val", "[REDACTED]"),
		)},
		{Key: "nested_type_val", Value: slog.GroupValue(
			slog.String("nested_string_val", "nested_string"),
			slog.String("nested_secret_val", "[REDACTED]"),
		)},
		{Key: "other_type_nested_type_val", Value: slog.GroupValue(
			slog.String("nested_string_val", "other_nested_string"),
			slog.String("nested_secret_val", "[REDACTED]"),
		)},
		{Key: "oneof_string_val", Value: slog.StringValue("")}, // 20
		{Key: "map_val1", Value: slog.GroupValue(
			slog.String("foo", "bar"),
		)},
		{Key: "map_val2", Value: slog.GroupValue(
			slog.Group("foo",
				slog.String("other_string_val", "other_string"),
				slog.String("other_secret_val", "[REDACTED]"),
			),
		)},
		{Key: "map_empty_val", Value: slog.GroupValue()},
		{Key: "repeated_val1", Value: slog.GroupValue(
			slog.String("0", "foo"),
			slog.String("1", "bar"),
		)},
		{Key: "repeated_val2", Value: slog.GroupValue(
			slog.String("0", Types_ENUM_1.String()),
			slog.String("1", Types_ENUM_2.String()),
		)},
		{Key: "repeated_val3", Value: slog.GroupValue(
			slog.Group("0",
				slog.String("other_string_val", "other_string"),
				slog.String("other_secret_val", "[REDACTED]"),
			),
		)},
		{Key: "repeated_empty_val", Value: slog.GroupValue()},
		{Key: "struct_val", Value: slog.AnyValue(&structpb.Struct{
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
		})},
		{Key: "_String", Value: slog.StringValue("foo")},
		{Key: "optional_val", Value: slog.StringValue("foo")},
		//{Key: "optional_not_present_val", Value: slog.StringValue("")}, // must not contains key
		{Key: "optional_enum", Value: slog.StringValue(Types_ENUM_1.String())},
		//{Key: "optional_not_present_enum", Value: slog.StringValue("")}, // must not contains key
		{Key: "optional_message", Value: slog.GroupValue(
			slog.String("val", "foo"),
		)},
		//{Key: "optional_not_present_message", Value: slog.GroupValue()}, // must not contains key
		{Key: "present_message", Value: slog.GroupValue(
			slog.String("val", "foo"),
		)},
		//{Key: "not_present_message", Value: slog.GroupValue()}, // must not contains key
	}
	// assert len
	if g, w := len(got), len(want); g != w {
		t.Errorf("\n got len:%v\nwant len:%v", len(got), len(want))
	}
	// assert attrs
	for i := range want {
		if len(got) <= i {
			break
		}
		g := got[i]
		w := want[i]
		if w.Value.Kind() == slog.KindAny {
			// []byte の compare で panic するのを回避
			g.Value = slog.StringValue(g.Value.String())
			w.Value = slog.StringValue(w.Value.String())
		}
		if !g.Equal(w) {
			t.Errorf("\nindex %d\n  got key:%s value string:%s\n want key:%s value string:%s", i, g.Key, g.String(), w.Key, w.String())
		}
	}
}

func refs[T any](v T) *T {
	return &v
}
