package test

import (
	"io"
	"log/slog"
	"testing"
)

// $ make bench
// goos: darwin
// goarch: amd64
// pkg: github.com/kei2100/protoc-gen-go-log-valuer/test
// cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
// BenchmarkTypes_LogValue-16                182263              6357 ns/op            1850 B/op         11 allocs/op
// BenchmarkTypes_NotUse_LogValue-16          61605             19753 ns/op            2291 B/op         76 allocs/op
// BenchmarkTypes_Manually-16                378253              3140 ns/op             600 B/op          2 allocs/op
// PASS
// ok      github.com/kei2100/protoc-gen-go-log-valuer/test        4.042s

func BenchmarkTypes_LogValue(b *testing.B) {
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
	}
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	for i := 0; i < b.N; i++ {
		logger.Info("bench", "message", m)
	}
}

func BenchmarkTypes_NotUse_LogValue(b *testing.B) {
	type types2 Types
	m := &types2{
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
	}
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	for i := 0; i < b.N; i++ {
		logger.Info("bench", "message", m)
	}
}

func BenchmarkTypes_Manually(b *testing.B) {
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
	}
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	for i := 0; i < b.N; i++ {
		logger.Info("bench", "message", slog.GroupValue(
			slog.String("secret_val", "[REDACTED]"),
			slog.Float64("double_val", m.DoubleVal),
			slog.Float64("float_val", float64(m.FloatVal)),
			slog.Int64("int32_val", int64(m.Int32Val)),
			slog.Int64("int64_val", m.Int64Val),
			slog.Uint64("uint32_val", uint64(m.Uint32Val)),
			slog.Uint64("uint64_val", m.Uint64Val),
			slog.Int64("sint32_val", int64(m.Sint32Val)),
			slog.Int64("sint64_val", m.Sint64Val),
			slog.Uint64("fixed32_val", uint64(m.Fixed32Val)),
			slog.Uint64("fixed64_val", m.Fixed64Val),
			slog.Int64("sfixed32_val", int64(m.Sfixed32Val)),
			slog.Int64("sfixed64_val", m.Sfixed64Val),
		))
	}
}
