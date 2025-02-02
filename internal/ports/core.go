package ports

import (
	"context"
	"time"

	"github.com/nxdir-s/idlerpg/internal/core/entity"
	"github.com/nxdir-s/idlerpg/internal/core/valobj"
	"github.com/nxdir-s/idlerpg/protobuf"
)

type Events interface {
	HandleUserEvent(ctx context.Context, event *protobuf.UserEvent) error
}

type Users interface {
	GetUser(ctx context.Context, id int) (*entity.User, error)
	Login(ctx context.Context, email string) (*valobj.Token, error)
	Register(ctx context.Context, email string) (*valobj.Token, error)
}

type UserService interface {
	GetUser(ctx context.Context, id int) (*entity.User, error)
	GetUserID(ctx context.Context, email string) (int, error)
	EmailExists(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, email string) (int, error)
}

type UserTxService interface {
	UserService
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type Characters interface {
	GetCharacter(ctx context.Context, userId int) (*entity.Character, error)
	HandleCombatEvent(ctx context.Context, character *entity.Character, event *protobuf.UserEvent) error
	HandleCraftEvent(ctx context.Context, character *entity.Character, event *protobuf.UserEvent) error
	HandleGatherEvent(ctx context.Context, character *entity.Character, event *protobuf.UserEvent) error
}

type CharacterService interface {
	GetCharacter(ctx context.Context, userId int) (*entity.Character, error)
}

type CharacterTxService interface {
	CharacterService
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type JWT interface {
	IssueToken(ctx context.Context, id int, email string) (*valobj.Token, error)
	IssueAccessToken(ctx context.Context, userId int, email string, expires time.Time) (string, error)
	IssueRefreshToken(ctx context.Context, userId int, email string, expires time.Time) (string, error)
	ValidAccessToken(ctx context.Context, tokenStr string) error
	ValidRefreshToken(ctx context.Context, tokenStr string) error
}
