package ports

import (
	"context"

	"github.com/nxdir-s/idlerpg/protobuf"
)

type Auth interface {
	ValidateToken(ctx context.Context, token string, tokenType string) error
}

type Consumer interface {
	ProcessUserEvent(ctx context.Context, event *protobuf.UserEvent) error
}
