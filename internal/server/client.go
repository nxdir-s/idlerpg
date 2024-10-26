package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/gobwas/ws"
	"github.com/nxdir-s/IdleRpg/internal/core/entity"
	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
)

type Client struct {
	Conn   net.Conn
	Fd     int
	Msgs   chan *valobj.Message
	Player *entity.Player
}

func (c *Client) ReadMessage(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			header, err := ws.ReadHeader(c.Conn)
			if err != nil {
				fmt.Fprintf(os.Stdout, "failed to read header: %+v\n", err)
				return
			}

			payload := make([]byte, header.Length)
			_, err = io.ReadFull(c.Conn, payload)
			if err != nil {
				fmt.Fprintf(os.Stdout, "failed to read client message: %+v\n", err)
				return
			}

			// TODO: update player state

			if header.OpCode == ws.OpClose {
				return
			}
		}
	}
}
