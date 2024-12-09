package user

import (
	"sync/atomic"
)

type Login string
type ID uint64

var userGID atomic.Uint64 = atomic.Uint64{}

type User struct {
	login Login
	id    ID
}

func GenerateID() ID {
	return ID(userGID.Add(1))
}

func New(login Login) *User {
	return &User{
		login: login,
		id:    GenerateID(),
	}
}

func (u User) Login() Login {
	return u.login
}

func (u User) ID() ID {
	return u.id
}
