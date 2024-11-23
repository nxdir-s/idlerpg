package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/nxdir-s/IdleRpg/internal/adapters/secondary"
	"github.com/nxdir-s/IdleRpg/internal/engine"
	"github.com/nxdir-s/IdleRpg/internal/ports"
	"github.com/nxdir-s/IdleRpg/internal/server"
)

const (
	DefaultAddr string = "0.0.0.0:3000"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		fmt.Fprintf(os.Stdout, "%+v\n", err)
		os.Exit(1)
	}

	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		fmt.Fprintf(os.Stdout, "%+v\n", err)
		os.Exit(1)
	}

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

	fmt.Fprintf(os.Stdout, "BROKERS: %s\n", brokerStr)

	var kafka ports.KafkaPort
	kafka, err = secondary.NewSaramaAdapter(strings.Split(brokerStr, ","))
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to create kafka adapter: %s\n", err.Error())
		os.Exit(1)
	}
	defer kafka.CloseProducer()

	pool := server.NewPool()
	ngin := engine.NewGameEngine(pool, kafka)

	epoll, err := server.NewEpoll(pool)
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
		os.Exit(0)
	}
}
