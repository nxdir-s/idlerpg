package actions

type Action uint8

const (
	Fight Action = 1 << iota
	Craft
	Gather
)
