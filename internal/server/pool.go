package server

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
	"golang.org/x/sys/unix"
)

type EpollEvent struct {
	Events []unix.EpollEvent
	Resp   chan []*Client
}

type Pool struct {
	Register    chan *Client
	Unregister  chan int
	Broadcast   chan *valobj.Event
	Snapshot    chan chan *Snapshot
	EpollEvents chan *EpollEvent
	Connections map[int]*Client

	counter int32
}

type Snapshot struct {
	Connections map[int]*Client
	Processed   chan bool
}

func NewPool() *Pool {
	return &Pool{
		Register:    make(chan *Client),
		Unregister:  make(chan int),
		Broadcast:   make(chan *valobj.Event),
		Snapshot:    make(chan chan *Snapshot),
		EpollEvents: make(chan *EpollEvent),
		Connections: make(map[int]*Client),
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

			p.Connections[client.Fd] = client

			if len(p.Connections)%100 == 0 {
				fmt.Fprintf(os.Stdout, "total number of connections: %d\n", len(p.Connections))
			}
		case fd := <-p.Unregister:
			fmt.Fprint(os.Stdout, "removing client from pool...\n")

			delete(p.Connections, fd)

			if len(p.Connections)%100 == 0 {
				fmt.Fprintf(os.Stdout, "total number of connections: %d\n", len(p.Connections))
			}
		case event := <-p.EpollEvents:
			var connections []*Client
			for i := range event.Events {
				connections = append(connections, p.Connections[int(event.Events[i].Fd)])
			}

			event.Resp <- connections
		case req := <-p.Snapshot:
			s := &Snapshot{
				Connections: p.Connections,
				Processed:   make(chan bool),
			}

			req <- s
			<-s.Processed
		case event := <-p.Broadcast:
			fmt.Fprintf(os.Stdout, "recieved event to broadcast: %s\n", event.Body.Body)

			for _, client := range p.Connections {
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
