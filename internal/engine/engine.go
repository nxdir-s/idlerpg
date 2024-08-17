// SIGUSR1 toggles the engine pausing/running
package engine

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
	"github.com/nxdir-s/IdleRpg/internal/pool"
	"github.com/nxdir-s/IdleRpg/internal/ports"
	"github.com/nxdir-s/pipelines"
)

const (
	MaxFan int = 3
)

type GameEngine struct {
	ticker      *time.Ticker
	sigusr1     chan os.Signal
	isPaused    bool
	connections *pool.Pool
	kafka       ports.KafkaPort
}

func New(pool *pool.Pool, kafka ports.KafkaPort) *GameEngine {
	return &GameEngine{
		ticker:      time.NewTicker(time.Second * 5),
		sigusr1:     make(chan os.Signal, 1),
		connections: pool,
		kafka:       kafka,
	}
}

func (ngin *GameEngine) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-ngin.sigusr1:
			ngin.isPaused = !ngin.isPaused

			switch ngin.isPaused {
			case true:
				fmt.Fprint(os.Stdout, "server is paused...\n")
			case false:
				fmt.Fprint(os.Stdout, "server is resumed...\n")
			}
		case t := <-ngin.ticker.C:
			if ngin.isPaused {
				break
			}

			fmt.Fprintf(os.Stdout, "server tick: %s\n", t.UTC().String())

			reply := make(chan *pool.Snapshot)
			ngin.connections.Snapshot <- reply

			snapshot := <-reply

			// process/simulate actions
			go ngin.process(ctx, snapshot.Clients)

			snapshot.Processed <- true

			event := &valobj.Event{
				Body: &valobj.Message{
					Type: 1,
					Body: "server tick: " + t.UTC().String(),
				},
				Consumed: make(chan bool),
			}

			ngin.connections.Broadcast <- event
			<-event.Consumed
		}
	}
}

func (ngin *GameEngine) process(ctx context.Context, players map[*pool.Client]bool) {
	// 1. Loop through players
	// 2. Simulate their actions
	//      1. For ex. player is fighting mob, their action is Fight
	//      2. The Fight action generates exp and random loot
	//      3. The random loot and other unknowns need to be simulated
	// 3. Run each simulation in go routine and send results to channel for batch save to db
	if len(players) == 0 {
		fmt.Fprintf(os.Stdout, "%d players connected...\n", len(players))
		return
	}

	fmt.Fprintf(os.Stdout, "processings %d player actions...\n", len(players))

	stream := pipelines.StreamMap(ctx, players)
	fanOutChannels := pipelines.FanOut(ctx, stream, ngin.Simulate, MaxFan)
	playerEvents := pipelines.FanIn(ctx, fanOutChannels...)

	for event := range playerEvents {
		select {
		case <-ctx.Done():
			return
		default:
			err := ngin.kafka.SendPlayerEvent(ctx, event)
			if err != nil {
				fmt.Fprintf(os.Stdout, "error sending event: %s\n", err.Error())
			}
		}
	}
}

func (ngin *GameEngine) Simulate(ctx context.Context, client *pool.Client) *valobj.PlayerEvent {
	fmt.Fprint(os.Stdout, "runing simulation...\n")

	return &valobj.PlayerEvent{
		Action: int(client.Player.Action),
	}
}
