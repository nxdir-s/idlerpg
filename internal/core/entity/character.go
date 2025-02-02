package entity

import (
	"github.com/nxdir-s/idlerpg/internal/core/entity/characters/stats"
)

type Character struct {
	Name   string
	Level  stats.Level
	Combat stats.Combat
	Craft  stats.Craft
	Gather stats.Gather
}
