package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nxdir-s/IdleRpg/internal/server"
)

const (
	DefaultAddr string = "0.0.0.0:3000"
)

type ArgError struct{}

func (e ArgError) Error() string {
	return "please provide an address to listen on as the first argument"
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", DefaultAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "listening on ws://%v\n", listener.Addr())

	gs := server.NewGameServer(ctx)
	server := &http.Server{
		Handler:      gs,
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
		fmt.Fprintf(os.Stdout, "context canceled: %+v\n", ctx.Err())
	case err := <-errChan:
		fmt.Fprintf(os.Stdout, "failed to serve: %+v\n", err)
	}

	ctx, timeout := context.WithTimeout(ctx, time.Second*10)
	defer timeout()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("error shutting down server: %+v\n", err)
	}
}
