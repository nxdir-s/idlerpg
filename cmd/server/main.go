package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/grafana/pyroscope-go"
	"github.com/nxdir-s/idlerpg/internal/adapters/secondary"
	"github.com/nxdir-s/idlerpg/internal/engine"
	"github.com/nxdir-s/idlerpg/internal/logs"
	"github.com/nxdir-s/idlerpg/internal/ports"
	"github.com/nxdir-s/idlerpg/internal/server"
	"github.com/nxdir-s/idlerpg/internal/util"
	"github.com/nxdir-s/telemetry"
)

const (
	DefaultAddr string = "0.0.0.0:3000"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := slog.New(logs.NewHandler(slog.NewTextHandler(os.Stdout, nil)))
	slog.SetDefault(logger)

	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		logger.Error("failed to get RLIMIT", slog.Any("err", err))
		os.Exit(1)
	}

	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		logger.Error("failed to set RLIMIT", slog.Any("err", err))
		os.Exit(1)
	}

	serviceName := os.Getenv("OTEL_SERVICE_NAME")
	if serviceName == "" {
		logger.Error("missing env var: OTEL_SERVICE_NAME")
		os.Exit(1)
	}

	otelEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if otelEndpoint == "" {
		logger.Error("missing env var: OTEL_EXPORTER_OTLP_ENDPOINT")
		os.Exit(1)
	}

	profileUrl := os.Getenv("PROFILE_URL")
	if profileUrl == "" {
		logger.Error("missing env var: PROFILE_URL")
		os.Exit(1)
	}

	gcUser := os.Getenv("GCLOUD_USER")
	if gcUser == "" {
		logger.Error("missing env var: GCLOUD_USER")
		os.Exit(1)
	}

	gcPass := os.Getenv("GCLOUD_PASSWORD")
	if gcUser == "" {
		logger.Error("missing env var: GCLOUD_PASSWORD")
		os.Exit(1)
	}

	cfg := &telemetry.Config{
		ServiceName:        serviceName,
		OtelEndpoint:       otelEndpoint,
		Insecure:           true,
		EnableSpanProfiles: true,
	}

	ctx, cleanup, err := telemetry.InitProviders(ctx, cfg)
	if err != nil {
		logger.Error("failed to initialize telemetry", slog.Any("err", err))
		os.Exit(1)
	}
	defer cleanup(ctx)

	tracer, err := util.GetTracer(ctx)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	profileCfg := pyroscope.Config{
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
	}

	profiler, err := pyroscope.Start(profileCfg)
	if err != nil {
		logger.Error("failed to start profiler", slog.Any("err", err))
		os.Exit(1)
	}
	defer profiler.Stop()

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", DefaultAddr)
	if err != nil {
		logger.Error("failed to create tcp listener", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info(fmt.Sprintf("listening on ws://%v", listener.Addr()))

	brokerStr := os.Getenv("BROKERS")
	if brokerStr == "" {
		logger.Error("found empty string for BROKERS")
		os.Exit(1)
	}

	logger.Info("BROKERS: " + brokerStr)

	rpUser := os.Getenv("REDPANDA_SASL_USERNAME")
	if rpUser == "" {
		logger.Error("found empty string for REDPANDA_SASL_USERNAME")
		os.Exit(1)
	}

	rpPass := os.Getenv("REDPANDA_SASL_PASSWORD")
	if rpPass == "" {
		logger.Error("found empty string for REDPANDA_SASL_PASSWORD")
		os.Exit(1)
	}

	var kafka ports.KafkaPort
	kafka, err = secondary.NewFranzAdapter(ctx, "user-events", logger, secondary.WithFranzProducer(strings.Split(brokerStr, ","), rpUser, rpPass))
	if err != nil {
		logger.Error("failed to create kafka adapter", slog.Any("err", err))
		os.Exit(1)
	}
	defer kafka.CloseProducer()

	pool := server.NewPool(ctx, tracer)
	ngin := engine.NewGameEngine(pool, kafka, logger, tracer)

	epoll, err := server.NewEpoll(ctx, pool, tracer)
	if err != nil {
		logger.Error("failed to create epoll", slog.Any("err", err))
		os.Exit(1)
	}

	server := server.NewGameServer(ctx, listener, epoll, ngin, pool)

	logger.Info("starting server")

	go server.Start(ctx)

	select {
	case <-ctx.Done():
		logger.Info(ctx.Err().Error())
	}
}
