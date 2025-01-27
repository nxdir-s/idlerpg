package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/nxdir-s/idlerpg/internal/adapters/secondary"
	"github.com/nxdir-s/idlerpg/internal/config"
	"github.com/nxdir-s/idlerpg/internal/engine"
	"github.com/nxdir-s/idlerpg/internal/observability"
	"github.com/nxdir-s/idlerpg/internal/ports"
	"github.com/nxdir-s/idlerpg/internal/server"
	"github.com/nxdir-s/telemetry"
	"go.opentelemetry.io/otel"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
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

	cfg, err := config.New(
		config.WithListenerAddr(),
		config.WithBrokers(),
		config.WithRedPandaUsr(),
		config.WithRedPandaPass(),
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

	cleanup, err := telemetry.InitProviders(ctx, otelCfg)
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

	logger.Info(fmt.Sprintf("listening on ws://%v", listener.Addr()))

	var kafka ports.KafkaPort
	kafka, err = secondary.NewFranzAdapter(
		logger,
		otel.Tracer("kafka.franz"),
		secondary.WithFranzProducer(
			strings.Split(cfg.Brokers, ","),
			cfg.RedPandaUsr,
			cfg.RedPandaPass,
		),
	)
	if err != nil {
		logger.Error("failed to create kafka adapter", slog.Any("err", err))
		os.Exit(1)
	}
	defer kafka.CloseProducer()

	pool := server.NewPool(ctx, logger, otel.Tracer("server.pool"))
	ngin := engine.NewGameEngine(pool, kafka, logger, otel.Tracer("engine"))

	epoll, err := server.NewEpoll(pool, logger, otel.Tracer("server.epoll"))
	if err != nil {
		logger.Error("failed to create epoll", slog.Any("err", err))
		os.Exit(1)
	}

	server := server.NewGameServer(ctx, listener, epoll, ngin, pool, logger)

	logger.Info("starting server")

	go server.Start(ctx)

	select {
	case <-ctx.Done():
		logger.Info(ctx.Err().Error())
		os.Exit(0)
	}
}
