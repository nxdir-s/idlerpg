package pool

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
	Snapshot   chan chan *Snapshot

	counter int32
}

type Snapshot struct {
	Clients   map[*Client]bool
	Processed chan bool
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan *valobj.Event),
		Snapshot:   make(chan chan *Snapshot),
	}
}

func (p *Pool) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case client := <-p.Register:
			fmt.Fprint(os.Stdout, "adding client to pool...\n")

			// using counter for testing
			p.counter++
			client.Player.Plid = p.counter

			p.Clients[client] = true
		case client := <-p.Unregister:
			fmt.Fprint(os.Stdout, "removing client from pool...\n")
			delete(p.Clients, client)
		case req := <-p.Snapshot:
			s := &Snapshot{
				Clients:   p.Clients,
				Processed: make(chan bool),
			}

			req <- s
			<-s.Processed
		case event := <-p.Broadcast:
			fmt.Fprintf(os.Stdout, "recieved event to broadcast: %s\n", event.Body.Body)

			for client := range p.Clients {
				select {
				case <-ctx.Done():
					return
				case client.Msgs <- event.Body:
				case <-time.After(80 * time.Millisecond):
					fmt.Fprint(os.Stdout, "client too slow, dropping event...\n")
				}
			}

			event.Consumed <- true
		}
	}
}
