package entity

import "github.com/nxdir-s/idlerpg/internal/core/entity/users"

type User struct {
	Id           int32
	Email        string
	RefreshToken string
	Ign          string
	State        users.GameState
}

func NewUser() *User {
	return &User{
		State: users.GameState{
			Action: users.Fight,
		},
	}
}
