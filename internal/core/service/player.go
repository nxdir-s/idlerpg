package service

type PlayerService struct{}

func NewPlayerService() (*PlayerService, error) {
	return &PlayerService{}, nil
}
