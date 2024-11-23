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

func (c *Client) SendMessage(ctx context.Context, msg *valobj.Message) error {
	wr := wsutil.NewWriter(c.Conn, ws.StateServerSide, ws.OpText)

	if err := json.NewEncoder(wr).Encode(msg); err != nil {
		return err
	}

	if err := wr.Flush(); err != nil {
		return err
	}

	return nil
}

func (c *Client) ReadMessage(ctx context.Context, epoller *Epoll) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "%s\n", ctx.Err().Error())
			return nil
		default:
			header, err := ws.ReadHeader(c.Conn)
			if err != nil {
				if err == io.EOF {
					fmt.Fprint(os.Stdout, "client disconnected\n")
					epoller.Remove <- c

					return nil
				}

				return err
			}

			payload := make([]byte, header.Length)
			_, err = io.ReadFull(c.Conn, payload)
			if err != nil {
				return err
			}

			// TODO: update player state

			if header.OpCode == ws.OpClose {
				return nil
			}
		}
	}
}
