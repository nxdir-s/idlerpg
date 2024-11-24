package entity

import (
	"github.com/nxdir-s/idlerpg/internal/core/entity/players"
)

type Player struct {
	Plid   int32
	Action players.Action
}

func NewPlayer() *Player {
	return &Player{}
}
