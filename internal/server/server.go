package server

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/gobwas/ws"
	"github.com/nxdir-s/pipelines"
)

const (
	MaxReadFan int = 3
)

type Startable interface {
	Start(ctx context.Context)
}

type GameServer struct {
	listener    net.Listener
	engine      Startable
	connections *Pool
	epoller     *Epoll
}

func NewGameServer(ctx context.Context, ln net.Listener, epoll *Epoll, ngin Startable, pool *Pool) *GameServer {
	go pool.Start(ctx)
	go epoll.Start(ctx)
	go ngin.Start(ctx)

	server := &GameServer{
		listener:    ln,
		engine:      ngin,
		connections: pool,
		epoller:     epoll,
	}

	return server
}

// Start listens for incoming tcp connections
func (gs *GameServer) Start(ctx context.Context) {
	go gs.listen(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			conn, err := gs.listener.Accept()
			if err != nil {
				fmt.Fprintf(os.Stdout, "failed to accept tcp connection: %+v\n", err)
				continue
			}

			_, err = ws.Upgrade(conn)
			if err != nil {
				fmt.Fprintf(os.Stdout, "failed to upgrade tcp connection: %+v\n", err)
				continue
			}

			gs.epoller.Add <- conn
		}
	}
}

// listen waits for incoming client messages
func (gs *GameServer) listen(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			clients, err := gs.epoller.Wait()
			if err != nil {
				fmt.Fprintf(os.Stdout, "failed to recieve epoll event: %+v\n", err)
				continue
			}

			readMsg := func(ctx context.Context, client *Client) error {
				return client.ReadMessage(ctx, gs.epoller)
			}

			stream := pipelines.StreamSlice(ctx, clients)
			fanOut := pipelines.FanOut(ctx, stream, readMsg, MaxReadFan)
			errChan := pipelines.FanIn(ctx, fanOut...)

			for err := range errChan {
				select {
				case <-ctx.Done():
					fmt.Fprintf(os.Stdout, "%s\n", ctx.Err().Error())
					return
				default:
					if err != nil {
						fmt.Fprintf(os.Stdout, "failed to read client message: %+v\n", err)
					}
				}
			}
		}
	}
}
