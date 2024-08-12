package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/coder/websocket"
	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
)

type Client struct {
	conn *websocket.Conn
	msgs chan *valobj.Message
}

// Read waits for client events and notifies the game server
func (c *Client) Read(ctx context.Context, pool *Pool) {
	defer func() {
		pool.Unregister <- c
		c.conn.CloseNow()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "recieved done signal: %v\n", ctx.Err())
			return
		default:
			_, r, err := c.conn.Reader(ctx)
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

			fmt.Fprintf(os.Stdout, "Recieved Message: %s\n", string(msg))

			// TODO: implement notifying game server
		}
	}
}

// Listen waits for server events and writes back to the client
func (c Client) Listen(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-c.msgs:
			fmt.Fprintf(os.Stdout, "recieved message, getting connection writer...\n")

			wr, err := c.conn.Writer(ctx, websocket.MessageText)
			if err != nil {
				fmt.Fprintf(os.Stdout, "error getting connection writer: %s\n", err.Error())
				return err
			}

			fmt.Fprintf(os.Stdout, "preparing to send message...\n")

			err = json.NewEncoder(wr).Encode(msg)
			if err != nil {
				fmt.Fprintf(os.Stdout, "error encoding message: %s\n", err.Error())
				return err
			}

			err = wr.Close()
			if err != nil {
				fmt.Fprintf(os.Stdout, "error closing connection writer: %s\n", err.Error())
				return err
			}
		}
	}
}
