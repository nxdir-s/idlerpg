package users

type Action uint8

const (
	Unknown Action = iota
	Fight
	Craft
	Gather
)

type State struct {
	Action Action
}
