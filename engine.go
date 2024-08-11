package main

// SIGUSR1 toggles the engine pausing/running
import (
	"context"
	"fmt"
	"os"
	"time"
)

type Engine struct {
	ticker   *time.Ticker
	sigusr1  chan os.Signal
	isPaused bool
}

func NewEngine() *Engine {
	return &Engine{
		ticker:  time.NewTicker(time.Second * 5),
		sigusr1: make(chan os.Signal, 1),
	}
}

func (e *Engine) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-e.sigusr1:
			e.isPaused = !e.isPaused

			switch e.isPaused {
			case true:
				fmt.Fprint(os.Stdout, "server is paused...\n")
			case false:
				fmt.Fprint(os.Stdout, "server is resumed...\n")
			}
		case t := <-e.ticker.C:
			if e.isPaused {
				break
			}

			fmt.Fprintf(os.Stdout, "server tick: %s\n", t.UTC().String())
		}
	}
}
