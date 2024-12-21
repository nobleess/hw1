package dto

import pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"

type Channel struct {
	ID   pgxUUID.UUID `db:"id"`
	Name string       `db:"name"`
}

type ChannelMembers struct {
	ID        pgxUUID.UUID `db:"id"`
	ChannelID pgxUUID.UUID `db:"channel_id"`
	UserID    pgxUUID.UUID `db:"user_id"`
}
