package domain

import (
	"context"

	"github.com/nxdir-s/idlerpg/internal/core/entity"
	"github.com/nxdir-s/idlerpg/internal/ports"
	"github.com/nxdir-s/idlerpg/protobuf"
)

type Characters struct {
	service ports.CharacterService
}

func NewCharacters(service ports.CharacterService) *Characters {
	return &Characters{
		service: service,
	}
}

func (d *Characters) GetCharacter(ctx context.Context, userId int) (*entity.Character, error) {
	return d.service.GetCharacter(ctx, userId)
}

func (d *Characters) HandleCombatEvent(ctx context.Context, character *entity.Character, event *protobuf.UserEvent) error {
	return nil
}

func (d *Characters) HandleCraftEvent(ctx context.Context, character *entity.Character, event *protobuf.UserEvent) error {
	return nil
}

func (d *Characters) HandleGatherEvent(ctx context.Context, character *entity.Character, event *protobuf.UserEvent) error {
	return nil
}
