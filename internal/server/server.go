package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/coder/websocket"
	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
	"github.com/nxdir-s/IdleRpg/internal/engine"
)

type GameServer struct {
	engine      *engine.GameEngine
	connections *Pool
	serveMux    http.ServeMux
}

func NewGameServer(ctx context.Context) *GameServer {
	pool := NewPool()
	go pool.Start(ctx)

	eng := engine.New(pool.Broadcast)
	defer eng.Start(ctx)

	gs := &GameServer{
		engine:      eng,
		connections: pool,
	}

	gs.serveMux.Handle("/", http.FileServer(http.Dir(".")))
	gs.serveMux.HandleFunc("/ws", gs.serveWS)

	return gs
}

func (server *GameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.serveMux.ServeHTTP(w, r)
}

func (server *GameServer) serveWS(w http.ResponseWriter, r *http.Request) {
	err := server.registerClient(r.Context(), w, r)

	if errors.Is(err, context.Canceled) {
		return
	}

	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}

	if err != nil {
		return
	}
}

// registerClient registers the incoming websocket connection to the game server
func (server *GameServer) registerClient(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}

	c := &Client{
		conn: conn,
		msgs: make(chan *valobj.Message),
	}

	defer func() {
		server.connections.Unregister <- c
		c.conn.CloseNow()
	}()

	server.connections.Register <- c

	go c.Read(ctx)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-c.msgs:
			wc, err := c.conn.Writer(ctx, websocket.MessageBinary)
			if err != nil {
				fmt.Fprintf(os.Stdout, "error getting connection writer: %s\n", err.Error())
				return err
			}

			defer wc.Close()

			err = json.NewEncoder(wc).Encode(msg)
			if err != nil {
				fmt.Fprintf(os.Stdout, "error sending message to client: %s\n", err.Error())
				return err
			}
		}
	}
}
