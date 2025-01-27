package server

import (
	"context"
	"log/slog"
	"net"
	"time"

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
	logger      *slog.Logger
}

func NewGameServer(ctx context.Context, ln net.Listener, epoll *Epoll, ngin Startable, pool *Pool, logger *slog.Logger) *GameServer {
	go pool.Start(ctx)
	go epoll.Start(ctx)
	go ngin.Start(ctx)

	server := &GameServer{
		listener:    ln,
		engine:      ngin,
		connections: pool,
		epoller:     epoll,
		logger:      logger,
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
				gs.logger.Error("failed to accept tcp connection", slog.Any("err", err))
				continue
			}

			_, err = ws.Upgrade(conn)
			if err != nil {
				gs.logger.Error("failed to upgrade tcp connection", slog.Any("err", err))
				continue
			}

			select {
			case gs.epoller.Add <- conn:
			case <-time.After(80 * time.Millisecond):
				gs.logger.Warn("add to epoller timed out")
			}
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
				gs.logger.Error("failed to recieve epoll event", slog.Any("err", err))
				continue
			}

			readMsg := func(ctx context.Context, client *Client) error {
				return client.ReadMessage(ctx, gs.epoller, gs.logger)
			}

			stream := pipelines.StreamSlice(ctx, clients)
			fanOut := pipelines.FanOut(ctx, stream, readMsg, MaxReadFan)
			errChan := pipelines.FanIn(ctx, fanOut...)

			for err := range errChan {
				select {
				case <-ctx.Done():
					return
				default:
					if err != nil {
						gs.logger.Error("failed to read client message", slog.Any("err", err))
					}
				}
			}
		}
	}
}
