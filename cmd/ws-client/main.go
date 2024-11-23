package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"

	"github.com/coder/websocket"
)

const (
	MaxClients int = 10
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if len(os.Args) < 2 {
		fmt.Fprint(os.Stdout, "please provide an address\n")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	for range MaxClients {
		wg.Add(1)
		go startWS(ctx, &wg, os.Args[1])
	}

	wg.Wait()
}

func startWS(ctx context.Context, wg *sync.WaitGroup, addr string) {
	defer wg.Done()

	conn, _, err := websocket.Dial(ctx, addr, nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error dialing websocket: %s\n", err.Error())
		return
	}
	defer conn.CloseNow()

	for {
		select {
		case <-ctx.Done():
			conn.Close(websocket.StatusNormalClosure, "")
			return
		default:
			msgType, r, err := conn.Reader(ctx)
			if err != nil {
				fmt.Fprintf(os.Stdout, "error getting reader from connection: %s\n", err.Error())
				return
			}

			var msg []byte
			msg, err = io.ReadAll(r)
			if err != nil {
				fmt.Fprintf(os.Stdout, "error reading message: %s\n", err.Error())
				return
			}

			switch msgType {
			case websocket.MessageText:
				fmt.Fprintf(os.Stdout, "recieved message: %s\n", string(msg))
			case websocket.MessageBinary:
				fmt.Fprint(os.Stdout, "recieved binary message\n")
			default:
				fmt.Fprintf(os.Stdout, "unknown websocket frame type: %d", msgType)
				return
			}
		}
	}
}
