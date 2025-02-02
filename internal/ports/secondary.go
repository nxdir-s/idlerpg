package ports

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/adapters/secondary/franz"
	"github.com/nxdir-s/idlerpg/internal/core/entity"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type KafkaPort interface {
	Send(ctx context.Context, record protoreflect.ProtoMessage) error
	Consume(ctx context.Context, consumer franz.Consumer)
	Close() error
}

type Database interface {
	NewTransactionAdapter(ctx context.Context) (DatabaseTx, error)
	CreateUser(ctx context.Context, email string) (int, error)
	RemoveUser(ctx context.Context, id int) error
	GetUser(ctx context.Context, id int) (*entity.User, error)
	GetCharacter(ctx context.Context, userId int) (*entity.Character, error)
	GetUserID(ctx context.Context, email string) (int, error)
	UserExists(ctx context.Context, id int) (bool, error)
	EmailExists(ctx context.Context, email string) (bool, error)
	UpdateRefreshToken(ctx context.Context, id int, token string) error
}

type DatabaseTx interface {
	Database
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
