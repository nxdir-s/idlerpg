package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nxdir-s/idlerpg/internal/config"
	"github.com/nxdir-s/idlerpg/internal/logs"
	"github.com/nxdir-s/idlerpg/internal/observability"
	"github.com/nxdir-s/idlerpg/web"
	"github.com/nxdir-s/telemetry"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := slog.New(logs.NewHandler(slog.NewTextHandler(os.Stdout, nil)))
	slog.SetDefault(logger)

	cfg, err := config.New(
		config.WithListenerAddr(),
		config.WithOtelServiceName(),
		config.WithOtelEndpoint(),
		config.WithProfileURL(),
		config.WithGrafanaUsr(),
		config.WithGrafanaPass(),
	)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	otelCfg := &telemetry.Config{
		ServiceName:        cfg.OtelService,
		OtelEndpoint:       cfg.OtelEndpoint,
		Insecure:           true,
		EnableSpanProfiles: true,
	}

	ctx, cleanup, err := telemetry.InitProviders(ctx, otelCfg)
	if err != nil {
		logger.Error("failed to initialize telemetry", slog.Any("err", err))
		os.Exit(1)
	}
	defer cleanup(ctx)

	profileCfg := &observability.ProfileConfig{
		ApplicationName: cfg.OtelService,
		ServerAddress:   cfg.ProfileURL,
		AuthUser:        cfg.GrafanaUsr,
		AuthPassword:    cfg.GrafanaPass,
	}

	profiler, err := observability.NewProfiler(profileCfg)
	if err != nil {
		logger.Error("failed to start profiler", slog.Any("err", err))
		os.Exit(1)
	}
	defer profiler.Stop()

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", cfg.ListenerAddr)
	if err != nil {
		logger.Error("failed to create tcp listener", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info(fmt.Sprintf("listening on %v", listener.Addr()))

	ws, err := web.NewServer(ctx, logger)
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
