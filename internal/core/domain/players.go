package domain

type Players struct{}

func NewPlayers() (*Players, error) {
	return &Players{}, nil
}
