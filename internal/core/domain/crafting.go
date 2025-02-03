package domain

type Crafting struct{}

func NewCrafting() *Crafting {
	return &Crafting{}
}

func (d *Crafting) HandleUserEvent() error {
	return nil
}
