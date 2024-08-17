package player

import "github.com/nxdir-s/IdleRpg/internal/core/entity/player/actions"

type Player struct {
	Action actions.Action
}

func New() *Player {
	return &Player{}
}
