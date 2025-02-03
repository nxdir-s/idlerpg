package engine

import (
	"context"
	"log/slog"
	"math/rand"
	"os"
	"time"

	"github.com/nxdir-s/idlerpg/internal/core/valobj"
	"github.com/nxdir-s/idlerpg/internal/ports"
	"github.com/nxdir-s/idlerpg/internal/server"
	"github.com/nxdir-s/idlerpg/protobuf"
	"github.com/nxdir-s/pipelines"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	MaxRandExp int32 = 100

	SimulateMaxFan int = 3
	KafkaMaxFan    int = 3

	TickerInterval time.Duration = time.Second * 5
)

type GameEngine struct {
	kafka    ports.Kafka
	pool     *server.Pool
	ticker   *time.Ticker
	sigusr1  chan os.Signal
	isPaused bool
	tracer   trace.Tracer
	logger   *slog.Logger
}

// NewGameEngine creates a GameEngine
func NewGameEngine(pool *server.Pool, kafka ports.Kafka, logger *slog.Logger, tracer trace.Tracer) *GameEngine {
	return &GameEngine{
		kafka:   kafka,
		pool:    pool,
		ticker:  time.NewTicker(TickerInterval),
		sigusr1: make(chan os.Signal, 1),
		tracer:  tracer,
		logger:  logger,
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
				ngin.logger.Info("server is paused")
			case false:
				ngin.logger.Info("server is resumed")
			}
		case t := <-ngin.ticker.C:
			if ngin.isPaused {
				break
			}

			ctx, span := ngin.tracer.Start(ctx, "tick")

			ngin.logger.Info("server tick")

			reply := make(chan *server.Snapshot)
			ngin.pool.Snapshot <- reply

			snapshot := <-reply

			ngin.process(ctx, snapshot.Connections)
			snapshot.Processed <- struct{}{}

			event := ngin.buildEvent(ctx, t)

			ngin.pool.Broadcast <- event
			<-event.Consumed

			span.End()
		}
	}
}

func (ngin *GameEngine) process(ctx context.Context, users map[int32]*server.Client) {
	if len(users) == 0 {
		ngin.logger.Info("0 users connected")
		return
	}

	ctx, span := ngin.tracer.Start(ctx, "process")
	defer span.End()

	ngin.logger.Info("processing user actions", slog.Int("connections", len(users)))

	stream := pipelines.StreamMap[int32, *server.Client](ctx, users)

	fanOutChannels := pipelines.FanOut(ctx, stream, ngin.Simulate, SimulateMaxFan)
	playerEvents := pipelines.FanIn(ctx, fanOutChannels...)

	kafkaFanOut := pipelines.FanOut(ctx, playerEvents, ngin.kafka.Send, KafkaMaxFan)
	errChan := pipelines.FanIn(ctx, kafkaFanOut...)

	for err := range errChan {
		select {
		case <-ctx.Done():
			return
		default:
			if err != nil {
				// TODO: figure out how to handle replays. Maybe dlq?
				ngin.logger.Error("failed to send user event", slog.Any("err", err))
			}
		}
	}
}

// Simulate simulates player actions
func (ngin *GameEngine) Simulate(ctx context.Context, client *server.Client) protoreflect.ProtoMessage {
	return &protobuf.UserEvent{
		Id:     client.User.Id,
		Action: protobuf.Action(client.User.State.Action),
		Exp:    rand.Int31n(MaxRandExp),
	}
}

func (ngin *GameEngine) buildEvent(ctx context.Context, t time.Time) *valobj.Event {
	return &valobj.Event{
		Ctx: ctx,
		Msg: &valobj.Message{
			Value: "server tick: " + t.UTC().String(),
		},
		Consumed: make(chan struct{}),
	}
}
