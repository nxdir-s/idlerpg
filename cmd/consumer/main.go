package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strings"

	"github.com/nxdir-s/idlerpg/internal/adapters/secondary"
	"github.com/nxdir-s/idlerpg/internal/config"
	"github.com/nxdir-s/idlerpg/internal/consumers"
	"github.com/nxdir-s/idlerpg/internal/logs"
	"github.com/nxdir-s/idlerpg/internal/observability"
	"github.com/nxdir-s/idlerpg/internal/ports"
	"github.com/nxdir-s/telemetry"
	"go.opentelemetry.io/otel"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := slog.New(logs.NewHandler(slog.NewTextHandler(os.Stdout, nil)))
	slog.SetDefault(logger)

	cfg, err := config.New(
		config.WithBrokers(),
		config.WithRedPandaUsr(),
		config.WithRedPandaPass(),
		config.WithConsumerName(),
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

	var kafka ports.KafkaPort
	kafka, err = secondary.NewFranzAdapter(
		logger,
		otel.Tracer("kafka.franz"),
		secondary.WithFranzConsumer(
			"user-events",
			cfg.ConsumerName,
			strings.Split(cfg.Brokers, ","),
			cfg.RedPandaUsr,
			cfg.RedPandaPass,
		),
	)
	if err != nil {
		logger.Error("failed to create kafka adapter", slog.Any("err", err))
		os.Exit(1)
	}

	consumer := consumers.NewEventsConsumer(kafka, logger)
	defer consumer.Close()

	go consumer.Start(ctx)

	select {
	case <-ctx.Done():
		logger.Info(ctx.Err().Error())
	}
}
