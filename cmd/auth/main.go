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
	"github.com/nxdir-s/idlerpg/internal/core/domain"
	"github.com/nxdir-s/idlerpg/internal/core/service"
	"github.com/nxdir-s/idlerpg/internal/logs"
	"github.com/nxdir-s/idlerpg/internal/ports"
)

const (
	DefaultAddr string = "0.0.0.0:9000"
)

type ErrService struct {
	service string
	err     error
}

func (e *ErrService) Error() string {
	return "error creating " + e.service + " service: " + e.err.Error()
}

type ErrOrchestrator struct {
	domain string
	err    error
}

func (e *ErrOrchestrator) Error() string {
	return "error creating " + e.domain + " orchestrator: " + e.err.Error()
}

type ErrAdapter struct {
	adapter string
	err     error
}

func (e *ErrAdapter) Error() string {
	return "error creating " + e.adapter + " adapter: " + e.err.Error()
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := slog.New(logs.NewHandler(slog.NewTextHandler(os.Stdout, nil)))
	slog.SetDefault(logger)

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", DefaultAddr)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to create tcp listener: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "listening on %v\n", listener.Addr())

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
		fmt.Fprintf(os.Stdout, "%+v\n", err)
		os.Exit(1)
	}

	database, err = secondary.NewPostgresAdapter(ctx, pgxPool, logger)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%+v\n", &ErrAdapter{"postgres", err})
		os.Exit(1)
	}

	userService, err = service.NewUserService(ctx, database)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%+v\n", &ErrService{"user", err})
		os.Exit(1)
	}

	jwt = domain.NewJWT("accessKey", "refreshKey")

	users, err = domain.NewUsers(ctx, userService, jwt)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%+v\n", &ErrOrchestrator{"users", err})
		os.Exit(1)
	}

	authAdapter, err = primary.NewAuthAdapter(ctx, users, jwt)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%+v\n", &ErrAdapter{"auth", err})
		os.Exit(1)
	}

	s, err := auth.NewServer(ctx, authAdapter)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to create auth server: %s\n", err.Error())
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
