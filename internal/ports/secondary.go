package ports

import (
	"context"

	"github.com/nxdir-s/idlerpg/protobuf"
)

type KafkaPort interface {
	SendUserEvent(ctx context.Context, event *protobuf.UserEvent) error
	CloseProducer() error
}

type DatabasePort interface {
	NewTransactionAdapter(ctx context.Context) (DatabaseTxPort, error)
	GetUserID(ctx context.Context, email string) (int, error)
}

type DatabaseTxPort interface {
	DatabasePort
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
