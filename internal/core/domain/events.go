package domain

import (
	"context"
	"strconv"

	"github.com/nxdir-s/idlerpg/internal/core/entity/users"
	"github.com/nxdir-s/idlerpg/internal/ports"
	"github.com/nxdir-s/idlerpg/protobuf"
)

type ErrUnknownAction struct {
	action int
}

func (e *ErrUnknownAction) Error() string {
	return "unknown action: " + strconv.Itoa(e.action)
}

type Events struct {
	users      ports.Users
	characters ports.Characters
}

func NewEvents(users ports.Users, characters ports.Characters) *Events {
	return &Events{
		users:      users,
		characters: characters,
	}
}

func (d *Events) HandleUserEvent(ctx context.Context, event *protobuf.UserEvent) error {
	character, err := d.characters.GetCharacter(ctx, int(event.Id))
	if err != nil {
		return err
	}

	switch users.Action(event.Action) {
	case users.Fight:
		return d.characters.HandleCombatEvent(ctx, character, event)
	case users.Craft:
		return d.characters.HandleCraftEvent(ctx, character, event)
	case users.Gather:
		return d.characters.HandleGatherEvent(ctx, character, event)
	default:
		return &ErrUnknownAction{int(event.Id)}
	}
}
