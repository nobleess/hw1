package text

import (
	"fmt"
	"main/internal/channel/channel"
	"main/internal/message/domain/model"
	user "main/internal/user/domain/model"
	"time"
)

const Type = "text"

// Message
type Message struct {
	id        model.ID   // `db:"id"`
	userId    user.ID    // `db:"user_id"`
	channelId channel.ID // `db:"channel_id"`
	text      string     // `db:"text"`
	createAt  time.Time  // `db:"create_at"`
	deleteAt  time.Time
	updateAt  time.Time // `db:"update_at"`
}

func NewMessage(id model.ID, userId user.ID, channelId channel.ID, createAt time.Time, updateAt time.Time, at time.Time, text string) *Message {
	return &Message{
		id:        id,
		userId:    userId,
		channelId: channelId,
		createAt:  createAt,
		updateAt:  updateAt,
		text:      text,
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

func (m *Message) Text() string {
	return m.text
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
	return fmt.Sprintf("TextMsg[ID: %d, From: %s, To: %s, Time: %s, Text: %s]", m.id, m.text)
}
