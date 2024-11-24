package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nxdir-s/idlerpg/web/server"
)

const (
	DefaultAddr string = "0.0.0.0:3001"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", DefaultAddr)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to create tcp listener: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "listening on %v\n", listener.Addr())

	ws, err := server.NewWebServer()
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
