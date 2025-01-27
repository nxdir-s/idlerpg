package ports

import (
	"context"

	"github.com/nxdir-s/idlerpg/protobuf"
)

type KafkaPort interface {
	SendUserEvent(ctx context.Context, event *protobuf.UserEvent) error
	SendUserUpdate(ctx context.Context, update *protobuf.UserUpdate) error
	CloseProducer() error
	ConsumeUserEvents(ctx context.Context)
	CloseConsumer() error
}

type DatabasePort interface {
	NewTransactionAdapter(ctx context.Context) (DatabaseTxPort, error)
	CreateUser(ctx context.Context, email string) (int, error)
	RemoveUser(ctx context.Context, id int) error
	GetUserID(ctx context.Context, email string) (int, error)
	UserExists(ctx context.Context, id int) (bool, error)
	EmailExists(ctx context.Context, email string) (bool, error)
	UpdateRefreshToken(ctx context.Context, id int, token string) error
}

type DatabaseTxPort interface {
	DatabasePort
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
