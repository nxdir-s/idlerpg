package server

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/gobwas/ws"
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

			for _, client := range clients {
				go client.ReadMessage(ctx, gs.epoller)
			}
		}
	}
}
