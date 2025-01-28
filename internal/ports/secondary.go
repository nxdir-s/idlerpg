package ports

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/adapters/secondary/franz"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type KafkaPort interface {
	Send(ctx context.Context, record protoreflect.ProtoMessage) error
	Consume(ctx context.Context, consumer franz.Consumer)
	Close() error
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
