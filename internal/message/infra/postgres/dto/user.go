package dto

import "time"

type User struct {
	ID        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Data      time.Time `db:"data"`
}
