package ports

import (
	"context"

	pb "github.com/nxdir-s/idlerpg/protobuf"
)

type KafkaPort interface {
	SendPlayerEvent(ctx context.Context, event *pb.PlayerEvent) error
	CloseProducer() error
}

type DatabasePort interface {
	NewTransactionAdapter(ctx context.Context) (DatabaseTxPort, error)
}

type DatabaseTxPort interface {
	DatabasePort
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
