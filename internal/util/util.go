package util

import (
	"fmt"
	"os"
	"time"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// RecordError sets the span status and records the error
func RecordError(span trace.Span, msg string, err error) {
	span.SetStatus(codes.Error, msg)
	span.RecordError(err)
}

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Fprintf(os.Stdout, "%s finished, duration: %s\n", name, time.Since(start).String())
	}
}
