package engine

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nxdir-s/IdleRpg/internal/core/valobj"
	"github.com/nxdir-s/IdleRpg/internal/ports"
	"github.com/nxdir-s/IdleRpg/internal/server"
	pb "github.com/nxdir-s/IdleRpg/protobuf"
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

func New(pool *server.Pool, kafka ports.KafkaPort) *GameEngine {
	return &GameEngine{
		kafka:   kafka,
		pool:    pool,
		ticker:  time.NewTicker(TickerInterval),
		sigusr1: make(chan os.Signal, 1),
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

			reply := make(chan *server.Snapshot)
			ngin.pool.Snapshot <- reply

			snapshot := <-reply

			ngin.process(ctx, snapshot.Connections)
			snapshot.Processed <- true

			event := ngin.buildEvent(t)

			ngin.pool.Broadcast <- event
			<-event.Consumed
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
		if err != nil {
			// TODO: figure out how to handle replays. Maybe dlq?
			fmt.Fprintf(os.Stdout, "error sending player event: %s\n", err.Error())
		}
	}
}

func (ngin *GameEngine) Simulate(ctx context.Context, client *server.Client) *pb.PlayerEvent {
	fmt.Fprint(os.Stdout, "running simulation...\n")

	return &pb.PlayerEvent{
		Plid:   client.Player.Plid,
		Action: pb.Action(int32(client.Player.Action)),
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
