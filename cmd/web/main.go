package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/grafana/pyroscope-go"
	"github.com/nxdir-s/idlerpg/internal/logs"
	"github.com/nxdir-s/idlerpg/web"
	"github.com/nxdir-s/telemetry"
)

const (
	DefaultAddr string = "0.0.0.0:8080"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := slog.New(logs.NewHandler(slog.NewTextHandler(os.Stdout, nil)))
	slog.SetDefault(logger)

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

	logger.Info(fmt.Sprintf("listening on %v", listener.Addr()))

	ws, err := web.NewServer(ctx)
	if err != nil {
		logger.Error("failed to create web server", slog.Any("err", err))
		os.Exit(1)
	}

	server := &http.Server{
		Handler:      ws,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	errChan := make(chan error, 1)
	go func() {
		logger.Info("starting server")
		errChan <- server.Serve(listener)
	}()

	select {
	case <-ctx.Done():
		logger.Info(ctx.Err().Error())
	case err := <-errChan:
		if err != nil {
			logger.Error("failed to serve", slog.Any("err", err))
		}
	}

	ctx, timeout := context.WithTimeout(ctx, time.Second*10)
	defer timeout()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("failed to shutdown server", slog.Any("err", err))
		os.Exit(1)
	}
}
