package dto

import (
	"main/internal/message/domain/model/user"
	"time"

	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

type User struct {
	ID       pgxUUID.UUID `db:"id"`
	Username string       `db:"username"`
	Password []byte       `db:"password"`
	Data     time.Time    `db:"data"`
}

func UserAdapter(users []User) []user.User {
	busers := make([]user.User, 0)
	for _, u := range users {
		busers = append(busers,
			*user.New(
				user.ID(u.ID),
				user.Login(u.Username),
			))
	}
	return busers
}
