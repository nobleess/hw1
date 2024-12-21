package dto

import (
	"main/internal/user/domain/model"
	"time"

	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

type User struct {
	ID       pgxUUID.UUID `db:"id"`
	Username string       `db:"username"`
	Password []byte       `db:"password"`
	Data     time.Time    `db:"data"`
}

func UserAdapter(users []User) []model.User {
	bmodel := make([]model.User, 0)
	for _, u := range users {
		bmodel = append(bmodel,
			*model.New(
				model.ID(u.ID),
				model.Login(u.Username),
			))
	}
	return bmodel
}
