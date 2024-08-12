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
			fmt.Fprint(os.Stdout, "adding client to pool...\n")
			p.Clients[client] = true
		case client := <-p.Unregister:
			fmt.Fprint(os.Stdout, "removing client from pool...\n")
			delete(p.Clients, client)
		case event := <-p.Broadcast:
			fmt.Fprintf(os.Stdout, "recieved event to broadcast: %s\n", event.Body.Body)

			for client := range p.Clients {
				select {
				case <-ctx.Done():
					return
				case client.msgs <- event.Body:
				case <-time.After(5 * time.Second):
					fmt.Fprint(os.Stdout, "client too slow, dropping event...\n")
				}
				// go func(c *Client, e *valobj.Event) {
				// 	select {
				// 	case <-ctx.Done():
				// 		return
				// 	case c.msgs <- e.Body:
				// 	case <-time.After(5 * time.Second):
				// 		fmt.Fprint(os.Stdout, "client too slow, dropping event...\n")
				// 	}
				// }(client, event)
			}

			event.Consumed <- true
		}
	}
}
