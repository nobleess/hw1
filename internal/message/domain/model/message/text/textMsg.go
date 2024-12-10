package text

import (
	"fmt"
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/user"
	"time"
)

// Message
type Message struct {
	id        message.ID // `db:"id"`
	userId    user.ID    // `db:"user_id"`
	channelId message.ID // `db:"channel_id"`
	text      string     // `db:"text"`
	CreateAt  time.Time  // `db:"create_at"`
	UpdateAt  time.Time  // `db:"update_at"`
}

func NewMessage(id message.ID, userId user.ID, channelId channel.ID, createAt time.Time, updateAt time.Time, text string) Message {
	return Message{
		id:        id,
		userId:    userId,
		channelId: channelId,
		CreateAt:  createAt,
		UpdateAt:  updateAt,
		text:      text,
	}
}

func (m *Message) From() user.Login {
	return m.from
}

func (m *Message) To() []user.Login {
	return m.to
}

func (m *Message) Time() time.Time {
	return m.time
}

func (m *Message) Id() message.ID {
	return m.id
}

func (m *Message) Text() string {
	return m.text
}

func (m *Message) String() string {
	return fmt.Sprintf("TextMsg[ID: %d, From: %s, To: %s, Time: %s, Text: %s]", m.id, m.from, m.to, m.time, m.text)
}

func (m *Message) Redact(new message.Message) {
	m.text = new.Text()
}
