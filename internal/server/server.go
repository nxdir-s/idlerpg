package server

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/gobwas/ws"
	"github.com/nxdir-s/IdleRpg/internal/core/entity"
	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
)

type GameEngine interface {
	Start(ctx context.Context)
}

type ClientPool interface {
	Start(ctx context.Context)
}

type Client struct {
	Conn   net.Conn
	Fd     int
	Msgs   chan *valobj.Message
	Player *entity.Player
}

type GameServer struct {
	engine      GameEngine
	connections ClientPool
	epoller     *Epoll
	listener    net.Listener
}

func NewGameServer(ctx context.Context, pool ClientPool, ngin GameEngine, ln net.Listener) *GameServer {
	defer func() {
		go ngin.Start(ctx)
	}()
	go pool.Start(ctx)

	gs := &GameServer{
		engine:      ngin,
		connections: pool,
		listener:    ln,
	}

	return gs
}

func (gs *GameServer) Start(ctx context.Context) {
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
