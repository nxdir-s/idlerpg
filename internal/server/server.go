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

	ngin := engine.New(pool.Broadcast)
	defer func() {
		go ngin.Start(ctx)
	}()

	gs := &GameServer{
		engine:      ngin,
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
	fmt.Fprint(os.Stdout, "recieved new client connection...\n")
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
			fmt.Fprintf(os.Stdout, "recieved message, getting connection writer...\n")

			wr, err := c.conn.Writer(ctx, websocket.MessageText)
			if err != nil {
				fmt.Fprintf(os.Stdout, "error getting connection writer: %s\n", err.Error())
				return err
			}

			fmt.Fprintf(os.Stdout, "preparing to send message...\n")

			err = json.NewEncoder(wr).Encode(msg)
			if err != nil {
				fmt.Fprintf(os.Stdout, "error encoding message: %s\n", err.Error())
				return err
			}

			err = wr.Close()
			if err != nil {
				fmt.Fprintf(os.Stdout, "error closing connection writer: %s\n", err.Error())
				return err
			}
		}
	}
}
