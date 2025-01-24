package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"strings"

	"github.com/grafana/pyroscope-go"
	"github.com/nxdir-s/idlerpg/internal/adapters/secondary"
	"github.com/nxdir-s/idlerpg/internal/logs"
	"github.com/nxdir-s/idlerpg/internal/ports"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := slog.New(logs.NewHandler(slog.NewTextHandler(os.Stdout, nil)))
	slog.SetDefault(logger)

	serviceName := os.Getenv("OTEL_SERVICE_NAME")
	if serviceName == "" {
		fmt.Fprint(os.Stdout, "missing env var: OTEL_SERVICE_NAME\n")
		os.Exit(1)
	}

	otelEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if otelEndpoint == "" {
		fmt.Fprint(os.Stdout, "missing env var: OTEL_EXPORTER_OTLP_ENDPOINT\n")
		os.Exit(1)
	}

	profileUrl := os.Getenv("PROFILE_URL")
	if profileUrl == "" {
		fmt.Fprint(os.Stdout, "missing env var: PROFILE_URL\n")
		os.Exit(1)
	}

	gcUser := os.Getenv("GCLOUD_USER")
	if gcUser == "" {
		fmt.Fprint(os.Stdout, "missing env var: GCLOUD_USER\n")
		os.Exit(1)
	}

	gcPass := os.Getenv("GCLOUD_PASSWORD")
	if gcUser == "" {
		fmt.Fprint(os.Stdout, "missing env var: GCLOUD_PASSWORD\n")
		os.Exit(1)
	}

	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	pyroscope.Start(pyroscope.Config{
		ApplicationName:   serviceName,
		ServerAddress:     profileUrl,
		BasicAuthUser:     gcUser,
		BasicAuthPassword: gcPass,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

	brokerStr := os.Getenv("BROKERS")
	if brokerStr == "" {
		fmt.Fprint(os.Stdout, "found empty string for BROKERS\n")
		os.Exit(1)
	}

	rpUser := os.Getenv("REDPANDA_SASL_USERNAME")
	if rpUser == "" {
		fmt.Fprint(os.Stdout, "found empty string for REDPANDA_SASL_USERNAME\n")
		os.Exit(1)
	}

	rpPass := os.Getenv("REDPANDA_SASL_PASSWORD")
	if rpPass == "" {
		fmt.Fprint(os.Stdout, "found empty string for REDPANDA_SASL_PASSWORD\n")
		os.Exit(1)
	}

	var kafka ports.KafkaPort
	kafka, err := secondary.NewFranzAdapter(ctx, logger, secondary.WithFranzConsumer(strings.Split(brokerStr, ","), rpUser, rpPass))
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to create kafka adapter: %s\n", err.Error())
		os.Exit(1)
	}
	defer kafka.CloseConsumer()

	go kafka.ConsumeUserEvent(ctx)

	select {
	case <-ctx.Done():
		fmt.Fprintf(os.Stdout, "%s\n", ctx.Err().Error())
	}
}
