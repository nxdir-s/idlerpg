package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/grafana/pyroscope-go"
	"github.com/nxdir-s/idlerpg/web"
	"github.com/nxdir-s/telemetry"
)

const (
	DefaultAddr string = "0.0.0.0:8080"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

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

	cfg := &telemetry.Config{
		ServiceName:        serviceName,
		OtelEndpoint:       otelEndpoint,
		Insecure:           true,
		EnableSpanProfiles: true,
	}

	ctx, cleanup, err := telemetry.InitProviders(ctx, cfg)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to initialize telemetry: %s\n", err.Error())
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
		fmt.Fprintf(os.Stdout, "failed to start profiler: %+v\n", err)
		os.Exit(1)
	}
	defer profiler.Stop()

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", DefaultAddr)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to create tcp listener: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "listening on %v\n", listener.Addr())

	ws, err := web.NewServer(ctx)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to create web server: %s\n", err.Error())
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
		fmt.Fprint(os.Stdout, "starting server...\n")
		errChan <- server.Serve(listener)
	}()

	select {
	case <-ctx.Done():
		fmt.Fprintf(os.Stdout, "%s\n", ctx.Err().Error())
	case err := <-errChan:
		if err != nil {
			fmt.Fprintf(os.Stdout, "failed to serve: %s\n", err.Error())
		}
	}

	ctx, timeout := context.WithTimeout(ctx, time.Second*10)
	defer timeout()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stdout, "failed to shutdown server: %s\n", err.Error())
		os.Exit(1)
	}
}
