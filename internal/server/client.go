package server

import (
	"context"
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

func (c *Client) Read(ctx context.Context) {
	for {
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
	}
}
