// SIGUSR1 toggles the engine pausing/running
package engine

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
	"github.com/nxdir-s/IdleRpg/internal/pool"
	"github.com/nxdir-s/pipelines"
)

type GameEngine struct {
	ticker      *time.Ticker
	sigusr1     chan os.Signal
	isPaused    bool
	connections *pool.Pool
}

func New(pool *pool.Pool) *GameEngine {
	return &GameEngine{
		ticker:      time.NewTicker(time.Second * 5),
		sigusr1:     make(chan os.Signal, 1),
		connections: pool,
	}
}

func (engine *GameEngine) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-engine.sigusr1:
			engine.isPaused = !engine.isPaused

			switch engine.isPaused {
			case true:
				fmt.Fprint(os.Stdout, "server is paused...\n")
			case false:
				fmt.Fprint(os.Stdout, "server is resumed...\n")
			}
		case t := <-engine.ticker.C:
			if engine.isPaused {
				break
			}

			fmt.Fprintf(os.Stdout, "server tick: %s\n", t.UTC().String())

			reply := make(chan *pool.Snapshot)
			engine.connections.Snapshot <- reply

			snapshot := <-reply

			// process/simulate actions

			event := &valobj.Event{
				Body: &valobj.Message{
					Type: 1,
					Body: "server tick: " + t.UTC().String(),
				},
				Consumed: make(chan bool),
			}

			engine.connections.Broadcast <- event
			<-event.Consumed
		}
	}
}

func (engine *GameEngine) process(players map[*pool.Client]bool) {
	// 1. Loop through players
	// 2. Simulate their actions
	//      1. For ex. player is fighting mob, their action is Fight
	//      2. The Fight action generates exp and random loot
	//      3. The random loot and other unknowns need to be simulated
	// 3. Run each simulation in go routine and send results to channel for batch save to db
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	playerStream := pipelines.StreamMap[*pool.Client](ctx, players)
	// fanOut := pipelines.FanOut(ctx, playerStream, )
}
