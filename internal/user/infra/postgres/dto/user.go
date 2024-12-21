package dto

import (
	"main/internal/user/domain/model"
	"time"

	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

type User struct {
	ID       pgxUUID.UUID `db:"id"`
	Username string       `db:"username"` // user see at
	Login    string       `db:"login"`    // only auth verbs mb = username
	Password []byte       `db:"password"`
	CreateAt time.Time    `db:"data"`
}

func UsersAdapter(users []User) []model.User {
	models := make([]model.User, 0)
	for _, u := range users {
		models = append(models,
			UserAdapter(u))
	}
	return models
}

func UserAdapter(u User) model.User {
	return *model.New(
		model.ID(u.ID),
		model.Login(u.Username),
	)
}
