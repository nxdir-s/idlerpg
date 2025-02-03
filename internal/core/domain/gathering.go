package domain

type Gathering struct{}

func NewGathering() *Gathering {
	return &Gathering{}
}

func (d *Gathering) HandleUserEvent() error {
	return nil
}
