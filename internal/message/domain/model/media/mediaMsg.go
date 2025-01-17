package media

import (
	"fmt"
	"main/internal/channel/channel"
	"main/internal/message/domain/model"
	user "main/internal/user/domain/model"
	"time"
)

const Type = "media"

// Message
type Message struct {
	id        model.ID
	userId    user.ID // `db:"user_id"`
	channelId channel.ID
	createAt  time.Time
	deleteAt  time.Time
	updateAt  time.Time
	mediaType string
	url       string
}

func NewMessage(id model.ID, userId user.ID, channelId channel.ID, UpdateAt, CreateAt time.Time, mediaType, url string) *Message {
	return &Message{
		id:        id,
		userId:    userId,
		channelId: channelId,
		createAt:  CreateAt,
		updateAt:  UpdateAt,
		mediaType: mediaType,
		url:       url,
	}
}

func (m *Message) Id() model.ID {
	return m.id
}

func (m *Message) UserId() user.ID {
	return m.userId
}

func (m *Message) ChannelId() channel.ID {
	return m.channelId
}

func (m *Message) URL() string {
	return m.url
}

func (m *Message) MediaType() string {
	return m.mediaType
}

func (m *Message) CreatedAt() time.Time {
	return m.createAt
}

func (m *Message) UpdatedAt() time.Time {
	return m.updateAt
}

func (m *Message) DeletedAt() time.Time {
	return m.deleteAt
}

func (m *Message) String() string {
	return fmt.Sprintf("Message[ID: %d, From: %s, To: %s, Time: %s, Type: %s, URL: %s]", m.id, m.mediaType, m.url)
}
