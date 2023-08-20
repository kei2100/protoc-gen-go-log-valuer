package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"log/slog"
	"slices"
	"testing"
	"time"
)

// TODO map, list, message, explicit presence
// TODO bench

func testJSON(t *testing.T, v slog.LogValuer, wantJSON string) {
	t.Helper()
	// marshal v
	rec := slog.NewRecord(time.Now(), slog.LevelInfo, "", 0)
	rec.AddAttrs(v.LogValue().Group()...)
	var out bytes.Buffer
	opt := &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if slices.Contains([]string{"time", "level", "msg"}, a.Key) {
				return slog.Attr{}
			}
			return a
		},
	}
	if err := slog.NewJSONHandler(&out, opt).Handle(context.Background(), rec); err != nil {
		t.Error(err)
	}
	// TODO remove debug print
	fmt.Println(out.String())
	// compare JSON
	var got, want interface{}
	if err := json.Unmarshal(out.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(wantJSON), &want); err != nil {
		t.Fatal(err)
	}
	diff := cmp.Diff(got, want)
	if diff != "" {
		t.Error(diff)
	}
}
