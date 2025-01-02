package util

import (
	"fmt"
	"os"
	"time"

	"github.com/nxdir-s/telemetry"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
)

type ErrTracer struct {
	err error
}

func (e *ErrTracer) Error() string {
	return "error reading tracer from context: " + e.err.Error()
}

type ErrNilTracer struct{}

func (e *ErrNilTracer) Error() string {
	return "nil tracer in context"
}

// GetTracer checks the context for a tracer, returns an error if a nil pointer is found
func GetTracer(ctx context.Context) (trace.Tracer, error) {
	tracer, err := telemetry.TracerFromContext(ctx)
	if err != nil {
		return nil, &ErrTracer{err}
	}

	if tracer == nil {
		return nil, &ErrNilTracer{}
	}

	return tracer, nil
}

type ErrMeter struct {
	err error
}

func (e *ErrMeter) Error() string {
	return "error reading meter from context: " + e.err.Error()
}

type ErrNilMeter struct{}

func (e *ErrNilMeter) Error() string {
	return "nil meter in context"
}

// GetMeter checks the context for a meter, returns an error if a nil pointer is found
func GetMeter(ctx context.Context) (metric.Meter, error) {
	meter, err := telemetry.MeterFromContext(ctx)
	if err != nil {
		return nil, &ErrMeter{err}
	}

	if meter == nil {
		return nil, &ErrNilMeter{}
	}

	return meter, nil
}

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
