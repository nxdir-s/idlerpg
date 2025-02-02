package entity

import "github.com/nxdir-s/idlerpg/internal/core/entity/characters"

type Character struct {
	Name        string
	Level       uint16
	CombatLevel uint16
	CombatExp   uint64
	CraftLevel  uint16
	CraftExp    uint64
	GatherLevel uint16
	GatherExp   uint64
	State       *characters.State
}
