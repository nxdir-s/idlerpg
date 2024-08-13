package actions

type Action uint8

const (
	Fight Action = iota + 1
	Craft
	Gather
)
