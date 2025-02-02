package primary

import (
	"context"

	"github.com/nxdir-s/idlerpg/protobuf"
)

type ConsumerAdapter struct{}

func NewConsumerAdapter() (*ConsumerAdapter, error) {
	return &ConsumerAdapter{}, nil
}

func (a *ConsumerAdapter) ProcessUserEvent(ctx context.Context, event *protobuf.UserEvent) error {
	return nil
}
