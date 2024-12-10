package dto

import (
	"main/internal/message/domain/model/message/text"
	"time"

	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

type Message struct {
	ID        pgxUUID.UUID `db:"id"`
	UserId    pgxUUID.UUID `db:"user_id"`
	ChannelId pgxUUID.UUID `db:"channel_id"`
	Text      string       `db:"text"`
	CreateAt  time.Time    `db:"create_at"`
	UpdateAt  time.Time    `db:"update_at"`
	DeleteAt  time.Time    `db:"delete_at"`
}

func MessageAdapter(msgs []Message) text.Message {
	bmodel := make([]text.Message, 0)
	for _, m := range msgs {
		bmodel = append(bmodel, *text.NewMessage())
	}
}

// type SystemMessage struct {
// 	ID        pgxUUID.UUID `db:"id"`
// 	UserId    pgxUUID.UUID `db:"user_id"`
// 	ChannelId pgxUUID.UUID `db:"from"`
// 	Text      string       `db:"text"`
// 	CreateAt  time.Time    `db:"create_at"`
// 	UpdateAt  time.Time    `db:"update_at"`
// 	DeleteAt  time.Time    `db:"delete_at"`
// }
