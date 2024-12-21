package dto

// убрать пользователя в другую папку intrernal
import (
	"encoding/json"
	"main/internal/channel/channel"
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/message/text"
	"main/internal/user/domain/model"
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

func MessageAdapter(msgs []Message) []message.Message {
	bmodel := make([]message.Message, 0)
	for _, m := range msgs {
		bmodel = append(bmodel, text.NewMessage(
			message.ID(m.ID),
			model.ID(m.UserId),
			channel.ID(m.ChannelId),
			m.CreateAt,
			m.UpdateAt,
			m.Text))

	}
	return bmodel
	json.RawMessage{}
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
