package test

import (
	"log/slog"
	"os"
)

var exampleLoggerOptions = &slog.HandlerOptions{
	ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
		// Remove time.
		if attr.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}
		return attr
	},
}

func ExampleExampleMessage_LogValue() {
	m := &ExampleMessage{
		Message: "foo",
		Count:   1,
		Secret:  "secret",
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, exampleLoggerOptions))
	logger.Info("out example message", "message", m)

	// Output:
	// level=INFO msg="out example message" message.message=foo message.count=1 message.secret=[REDACTED]
}
