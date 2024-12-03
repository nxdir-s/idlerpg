package engine

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nxdir-s/idlerpg/internal/core/valobj"
	"github.com/nxdir-s/idlerpg/internal/ports"
	"github.com/nxdir-s/idlerpg/internal/server"
	"github.com/nxdir-s/idlerpg/internal/util"
	"github.com/nxdir-s/idlerpg/protobuf"
	"github.com/nxdir-s/pipelines"
)

const (
	MaxRandExp int32 = 100

	SimulateMaxFan int = 3
	KafkaMaxFan    int = 3

	TickerInterval time.Duration = time.Second * 5
)

type GameEngine struct {
	kafka    ports.KafkaPort
	pool     *server.Pool
	ticker   *time.Ticker
	sigusr1  chan os.Signal
	isPaused bool
}

// NewGameEngine creates a GameEngine
func NewGameEngine(ctx context.Context, pool *server.Pool, kafka ports.KafkaPort) *GameEngine {
	return &GameEngine{
		kafka:   kafka,
		pool:    pool,
		ticker:  time.NewTicker(TickerInterval),
		sigusr1: make(chan os.Signal, 1),
	}
}

// Start runs the game server
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

			time := util.Timer("Server Tick")

			fmt.Fprintf(os.Stdout, "server tick: %s\n", t.UTC().String())

			reply := make(chan *server.Snapshot)
			ngin.pool.Snapshot <- reply

			snapshot := <-reply

			ngin.process(ctx, snapshot.Connections)
			snapshot.Processed <- true

			event := ngin.buildEvent(t)

			ngin.pool.Broadcast <- event
			<-event.Consumed
			time()
		}
	}
}

func (ngin *GameEngine) process(ctx context.Context, players map[int]*server.Client) {
	if len(players) == 0 {
		fmt.Fprint(os.Stdout, "0 players connected...\n")
		return
	}

	fmt.Fprintf(os.Stdout, "processings %d player actions...\n", len(players))

	stream := pipelines.StreamMap[int, *server.Client](ctx, players)

	fanOutChannels := pipelines.FanOut(ctx, stream, ngin.Simulate, SimulateMaxFan)
	playerEvents := pipelines.FanIn(ctx, fanOutChannels...)

	kafkaFanOut := pipelines.FanOut(ctx, playerEvents, ngin.kafka.SendPlayerEvent, KafkaMaxFan)
	errChan := pipelines.FanIn(ctx, kafkaFanOut...)

	for err := range errChan {
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "%s\n", ctx.Err().Error())
			return
		default:
			if err != nil {
				// TODO: figure out how to handle replays. Maybe dlq?
				fmt.Fprintf(os.Stdout, "error sending player event: %s\n", err.Error())
			}
		}
	}
}

// Simulate simulates player actions
func (ngin *GameEngine) Simulate(ctx context.Context, client *server.Client) *protobuf.PlayerEvent {
	fmt.Fprint(os.Stdout, "running simulation...\n")

	return &protobuf.PlayerEvent{
		Plid:   client.Player.Plid,
		Action: protobuf.Action(int32(client.Player.Action)),
		Exp:    rand.Int31n(MaxRandExp),
	}
}

func (ngin *GameEngine) buildEvent(t time.Time) *valobj.Event {
	return &valobj.Event{
		Body: &valobj.Message{
			Type: 1,
			Body: "server tick: " + t.UTC().String(),
		},
		Consumed: make(chan bool),
	}
}
