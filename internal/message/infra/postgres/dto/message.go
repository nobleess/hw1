package dto

// убрать пользователя в другую папку intrernal
import (
	"fmt"
	"main/internal/channel/channel"
	"main/internal/message/domain/model"
	"main/internal/message/domain/model/media"
	"main/internal/message/domain/model/system"
	"main/internal/message/domain/model/text"
	user "main/internal/user/domain/model"

	"time"

	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

type Message struct {
	ID        pgxUUID.UUID `db:"id"`
	UserId    pgxUUID.UUID `db:"user_id"`
	ChannelId pgxUUID.UUID `db:"channel_id"`
	Text      string       `db:"text"`
	MediaType string       `db:"media_type"`
	URL       string       `db:"url"`
	Type      string       `db:"type"`
	CreateAt  time.Time    `db:"create_at"`
	UpdateAt  time.Time    `db:"update_at"`
	DeleteAt  time.Time    `db:"delete_at"`
}

func NewMessage(m model.Message) Message {
	var dto Message

	switch msg := m.(type) {
	case *text.Message:
		dto = Message{
			UserId:    pgxUUID.UUID(msg.UserId()),
			ChannelId: pgxUUID.UUID(msg.ChannelId()),
			Text:      msg.Text(),
			Type:      text.Type,
		}
	case *system.Message:
		dto = Message{
			UserId:    pgxUUID.UUID(msg.UserId()),
			ChannelId: pgxUUID.UUID(msg.ChannelId()),
			Text:      msg.Text(),
			Type:      system.Type,
		}
	case *media.Message:
		dto = Message{
			UserId:    pgxUUID.UUID(msg.UserId()),
			ChannelId: pgxUUID.UUID(msg.ChannelId()),
			Type:      media.Type,
			URL:       msg.URL(),
			MediaType: msg.MediaType(),
		}
	default:
		return dto
	}
	return dto
}

func MessagesAdapter(mm []Message) ([]model.Message, error) {
	bmodels := make([]model.Message, len(mm))
	for _, m := range mm {
		tmp, err := MessageAdapter(m)
		if err != nil {
			return nil, err
		}
		bmodels = append(bmodels, tmp)
	}
	return bmodels, nil
}

func MessageAdapter(m Message) (model.Message, error) {
	var bmodel model.Message
	switch m.Type {
	case "text":
		bmodel = text.NewMessage(
			model.ID(m.ID),
			user.ID(m.UserId),
			channel.ID(m.ChannelId),
			m.CreateAt,
			m.UpdateAt,
			m.DeleteAt,
			m.Text)
	case "system":
		bmodel = system.NewMessage(
			model.ID(m.ID),
			user.ID(m.UserId),
			channel.ID(m.ChannelId),
			m.CreateAt,
			m.UpdateAt,
			m.Text)
	case "media":
		bmodel = media.NewMessage(
			model.ID(m.ID),
			user.ID(m.UserId),
			channel.ID(m.ChannelId),
			m.CreateAt,
			m.UpdateAt,
			m.MediaType,
			m.URL,
		)
	default:
		return bmodel, fmt.Errorf("unknown type: %s", m.Type)
	}
	return bmodel, nil
}

// type SystemMessage struct {
// 	ID        pgxUUID.UUID `db:"id"`
// 	UserId    pgxUUID.UUID `db:"user_id"`
// 	ChannelId pgxUUID.UUID `db:"from"`
// 	Text      string       `db:"text"`
// 	createAt  time.Time    `db:"create_at"`
// 	updateAt  time.Time    `db:"update_at"`
// 	DeleteAt  time.Time    `db:"delete_at"`
// }
