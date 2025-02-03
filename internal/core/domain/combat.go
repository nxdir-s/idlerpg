package domain

type Combat struct{}

func NewCombat() *Combat {
	return &Combat{}
}

func (d *Combat) HandleUserEvent() error {
	return nil
}
