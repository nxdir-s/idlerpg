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

	"github.com/nxdir-s/idlerpg/internal/adapters/primary"
	"github.com/nxdir-s/idlerpg/internal/adapters/secondary"
	"github.com/nxdir-s/idlerpg/internal/auth"
	"github.com/nxdir-s/idlerpg/internal/config"
	"github.com/nxdir-s/idlerpg/internal/core/domain"
	"github.com/nxdir-s/idlerpg/internal/core/service"
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

	// secondary adapters
	var database ports.DatabasePort

	// services
	var userService ports.UserServicePort

	// domains
	var users ports.UsersPort
	var jwt ports.JWTPort

	// primary adapters
	var authAdapter ports.AuthPort

	var pgxPool secondary.PgxPool
	pgxPool, err = secondary.NewPgxPool(ctx, "dbUrl")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	database, err = secondary.NewPostgresAdapter(pgxPool, logger, otel.Tracer("postgres"))
	if err != nil {
		logger.Error("failed to create postgres adapter", slog.Any("err", err))
		os.Exit(1)
	}

	userService, err = service.NewUserService(ctx, database)
	if err != nil {
		logger.Error("failed to create user service", slog.Any("err", err))
		os.Exit(1)
	}

	jwt = domain.NewJWT("accessKey", "refreshKey")

	users, err = domain.NewUsers(ctx, userService, jwt)
	if err != nil {
		logger.Error("failed to create users domain orchestrator", slog.Any("err", err))
		os.Exit(1)
	}

	authAdapter, err = primary.NewAuthAdapter(ctx, users, jwt)
	if err != nil {
		logger.Error("failed to create auth adapter", slog.Any("err", err))
		os.Exit(1)
	}

	s, err := auth.NewServer(ctx, authAdapter)
	if err != nil {
		logger.Error("failed to create auth server", slog.Any("err", err))
		os.Exit(1)
	}

	server := &http.Server{
		Handler:      s,
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
