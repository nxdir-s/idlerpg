package entity

type User struct {
	Id           int32
	Email        string
	RefreshToken string
	Character    *Character
}

func NewUser() *User {
	return &User{
		Character: &Character{},
	}
}
