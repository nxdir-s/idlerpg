package entity

import "github.com/nxdir-s/idlerpg/internal/core/entity/users"

type User struct {
	Id           int32
	Email        string
	RefreshToken string
	State        *users.State
	Character    *Character
}

func NewUser() *User {
	return &User{
		State: &users.State{
			Action: users.Fight,
		},
		Character: &Character{},
	}
}
