package engine

// SIGUSR1 toggles the engine pausing/running
import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
)

type GameEngine struct {
	ticker   *time.Ticker
	sigusr1  chan os.Signal
	isPaused bool
	emitter  chan<- *valobj.Event
}

func New(emit chan<- *valobj.Event) *GameEngine {
	return &GameEngine{
		ticker:  time.NewTicker(time.Second * 5),
		sigusr1: make(chan os.Signal, 1),
		emitter: emit,
	}
}

func (ge *GameEngine) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-ge.sigusr1:
			ge.isPaused = !ge.isPaused

			switch ge.isPaused {
			case true:
				fmt.Fprint(os.Stdout, "server is paused...\n")
			case false:
				fmt.Fprint(os.Stdout, "server is resumed...\n")
			}
		case t := <-ge.ticker.C:
			if ge.isPaused {
				break
			}

			fmt.Fprintf(os.Stdout, "server tick: %s\n", t.UTC().String())

			event := &valobj.Event{
				Body: &valobj.Message{
					Type: 1,
					Body: "server tick: " + t.UTC().String(),
				},
				Consumed: make(chan bool),
			}

			ge.emitter <- event
			<-event.Consumed
		}
	}
}
