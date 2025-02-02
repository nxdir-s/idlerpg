package domain

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/ports"
	"github.com/nxdir-s/idlerpg/protobuf"
)

type Events struct {
	users ports.Users
}

func NewEvents(users ports.Users) *Events {
	return &Events{
		users: users,
	}
}

func (d *Events) HandleUserEvent(ctx context.Context, event *protobuf.UserEvent) error {
	user, err := d.users.GetUser(ctx, int(event.Id))
	if err != nil {
		return err
	}

	// get user action/validate

	return nil
}
