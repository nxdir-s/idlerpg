package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/coder/websocket"
)

const (
	ClientMessageBuffer int = 32
)

type Client struct {
	conn *websocket.Conn
	pool *Pool
	msgs chan []byte
}

func (c *Client) Read(ctx context.Context) {
	defer func() {
		c.pool.Unregister <- c
		c.conn.CloseNow()
	}()

	for {
		_, r, err := c.conn.Reader(ctx)
		if err != nil {
			fmt.Fprintf(os.Stdout, "error getting reader from connection: %s\n", err.Error())
			return
		}

		var msg []byte
		msg, err = io.ReadAll(r)
		if err != nil {
			fmt.Fprintf(os.Stdout, "error reading message: %s\n", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "Recieved Message: %s\n", string(msg))
	}
}

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan []byte
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
	}
}

func (p *Pool) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case client := <-p.Register:
			p.Clients[client] = true
		case client := <-p.Unregister:
			delete(p.Clients, client)
		case msg := <-p.Broadcast:
			for client := range p.Clients {
				go func(c *Client) {
					select {
					case <-ctx.Done():
						return
					case c.msgs <- msg:
					}
				}(client)
			}
		}
	}
}

type GameServer struct {
	serveMux    http.ServeMux
	connections *Pool
}

func NewGameServer(ctx context.Context) *GameServer {
	pool := NewPool()
	go pool.Start(ctx)

	gs := &GameServer{
		connections: pool,
	}

	gs.serveMux.Handle("/", http.FileServer(http.Dir(".")))
	gs.serveMux.HandleFunc("/ws", gs.serveWS)

	return gs
}

func (gs *GameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gs.serveMux.ServeHTTP(w, r)
}

func (gs *GameServer) serveWS(w http.ResponseWriter, r *http.Request) {
	err := gs.registerClient(r.Context(), w, r)

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
func (gs *GameServer) registerClient(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}

	c := &Client{
		conn: conn,
		pool: gs.connections,
		msgs: make(chan []byte),
	}

	defer func() {
		gs.connections.Unregister <- c
		c.conn.CloseNow()
	}()

	gs.connections.Register <- c

	go c.Read(ctx)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-c.msgs:
			if err := c.conn.Write(ctx, websocket.MessageBinary, msg); err != nil {
				fmt.Fprintf(os.Stdout, "error sending message to client: %s\n", err.Error())
				return err
			}
		}
	}
}
