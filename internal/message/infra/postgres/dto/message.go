package dto

import (
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/user"
	"time"

	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

type TextMessage struct {
	ID   pgxUUID.UUID   `db:"id"`
	From pgxUUID.UUID   `db:"from"`
	To   []pgxUUID.UUID `db:"to"`
	Time time.Time      `db:time`
	Text string         `db:text`
}
