package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/nxdir-s/idlerpg/internal/core/entity"
	"github.com/nxdir-s/idlerpg/internal/core/valobj"
)

type Client struct {
	Conn net.Conn
	Fd   int32
	User *entity.User
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

func (c *Client) ReadMessage(ctx context.Context, epoller *Epoll, logger *slog.Logger) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			header, err := ws.ReadHeader(c.Conn)
			if err != nil {
				if err == io.EOF {
					logger.Info(fmt.Sprintf("%d disconnected", c.User.Id))
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

			// TODO: update user state

			if header.OpCode == ws.OpClose {
				return nil
			}
		}
	}
}
