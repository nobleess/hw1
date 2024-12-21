package model

import (
	"github.com/google/uuid"
)

type Login string
type ID uuid.UUID

type User struct {
	login Login
	id    ID
}

func New(id ID, login Login) *User {
	return &User{
		login: login,
		id:    id,
	}
}

func (u User) Login() Login {
	return u.login
}

func (u User) ID() ID {
	return u.id
}
