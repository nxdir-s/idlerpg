package ports

import (
	"context"

	pb "github.com/nxdir-s/IdleRpg/protobuf"
)

type KafkaPort interface {
	SendPlayerEvent(ctx context.Context, event *pb.PlayerEvent) error
	CloseProducer() error
}

type DatabasePort interface {
}

type DatabaseTxPort interface {
	DatabasePort
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
