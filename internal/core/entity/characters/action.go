package characters

type Action uint8

const (
	Unknown Action = iota
	Fight
	Craft
	Gather
)

func NewAction(action int) Action {
	return Action(action)
}
