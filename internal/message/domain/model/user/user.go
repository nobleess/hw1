package user

import (
	"sync/atomic"

	"github.com/google/uuid"
)

type Login string
type ID uuid.UUID

var userGID atomic.Uint64 = atomic.Uint64{}

type User struct {
	login Login
	id    ID
}

// func GenerateID() ID {
// 	return ID(userGID.Add(1))
// }

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
