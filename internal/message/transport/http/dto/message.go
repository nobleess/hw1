package dto

import (
	"fmt"
	uuid "github.com/google/uuid"
	"main/internal/channel/channel"
	"main/internal/message/domain/model"
	"main/internal/message/domain/model/media"
	"main/internal/message/domain/model/system"
	"main/internal/message/domain/model/text"
	user "main/internal/user/domain/model"
)

type Message struct {
	UserId    uuid.UUID `json:"user_id"`
	ChannelId uuid.UUID `json:"channel_id"`
	Text      string    `json:"text"`
	MediaType string    `json:"media_type"`
	URL       string    `json:"url"`
	Type      string    `json:"type"`
}

func AdapterMessage(m Message) model.Message {

	var bmodel model.Message
	switch m.Type {
	case "text":
		bmodel = text.NewMessage(
			model.ID(m.ID),
			user.ID(m.UserId),
			channel.ID(m.ChannelId),
			"",
			"",
			"",
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

	return
}
