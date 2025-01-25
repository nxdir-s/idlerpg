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
		fmt.Fprintf(os.Stdout, "failed to get RLIMIT: %s\n", err.Error())
		os.Exit(1)
	}

	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		fmt.Fprintf(os.Stdout, "failed to set RLIMIT: %s\n", err.Error())
		os.Exit(1)
	}

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
		fmt.Fprintf(os.Stdout, "failed to start profiler: %+v\n", err)
		os.Exit(1)
	}
	defer profiler.Stop()

	// cfg := &telemetry.Config{
	// 	ServiceName:  serviceName,
	// 	OtelEndpoint: otelEndpoint,
	// }
	//
	// ctx, cleanup, err := telemetry.InitProviders(ctx, cfg)
	// if err != nil {
	// 	fmt.Fprintf(os.Stdout, "failed to initialize telemetry: %s\n", err.Error())
	// 	os.Exit(1)
	// }

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", DefaultAddr)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to create tcp listener: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "listening on ws://%v\n", listener.Addr())

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

	fmt.Fprintf(os.Stdout, "BROKERS: %s\n", brokerStr)

	var kafka ports.KafkaPort
	kafka, err = secondary.NewFranzAdapter(ctx, "user-events", logger, secondary.WithFranzProducer(strings.Split(brokerStr, ","), rpUser, rpPass))
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to create kafka adapter: %s\n", err.Error())
		os.Exit(1)
	}
	defer kafka.CloseProducer()

	pool := server.NewPool(ctx)
	ngin := engine.NewGameEngine(ctx, pool, kafka)

	epoll, err := server.NewEpoll(ctx, pool)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to create epoll: %s\n", err.Error())
		os.Exit(1)
	}

	server := server.NewGameServer(ctx, listener, epoll, ngin, pool)

	fmt.Fprint(os.Stdout, "starting server...\n")

	go server.Start(ctx)

	select {
	case <-ctx.Done():
		fmt.Fprintf(os.Stdout, "%s\n", ctx.Err().Error())
	}

	// ctx, timeout := context.WithTimeout(ctx, time.Second*10)
	// defer timeout()
	//
	// go cleanup(ctx)
}
