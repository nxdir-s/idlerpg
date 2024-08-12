package server

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan *valobj.Event
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan *valobj.Event),
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
		case event := <-p.Broadcast:
			for client := range p.Clients {
				go func(c *Client, e *valobj.Event) {
					select {
					case <-ctx.Done():
						return
					case c.msgs <- e.Body:
					case <-time.After(3 * time.Second):
						fmt.Fprintf(os.Stdout, "client too slow, dropping event...")
					}
				}(client, event)
			}

			event.Consumed <- true
		}
	}
}
