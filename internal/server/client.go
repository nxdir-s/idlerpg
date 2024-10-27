package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/nxdir-s/IdleRpg/internal/core/entity"
	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
)

type Client struct {
	Conn   net.Conn
	Fd     int
	Msgs   chan *valobj.Message
	Player *entity.Player
}

func (c *Client) SendMessage(ctx context.Context, msg *valobj.Message) {
	wr := wsutil.NewWriter(c.Conn, ws.StateServerSide, ws.OpText)

	if err := json.NewEncoder(wr).Encode(msg); err != nil {
		fmt.Fprintf(os.Stdout, "error encoding message: %+v\n", err)
		return
	}

	if err := wr.Flush(); err != nil {
		fmt.Fprintf(os.Stdout, "error flushing writer: %+v\n", err)
	}
}

func (c *Client) ReadMessage(ctx context.Context, epoller *Epoll) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			header, err := ws.ReadHeader(c.Conn)
			if err != nil {
				if err == io.EOF {
					epoller.Remove <- c
				}

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
