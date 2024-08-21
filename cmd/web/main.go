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

	"github.com/nxdir-s/IdleRpg/web/server"
)

const (
	DefaultAddr string = "0.0.0.0:8080"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", DefaultAddr)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error creating listener: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "listening on %v\n", listener.Addr())

	ws, err := server.NewWebServer()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error creating web server: %v\n", err)
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
