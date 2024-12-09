package text

import (
	"fmt"
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/user"
	"time"
)

// Message
type Message struct {
	id   message.ID
	from user.Login
	to   []user.Login
	time time.Time
	text string
}

func NewMessage(from user.Login, to []user.Login, time time.Time, text string) *Message {
	return &Message{
		id:   message.GenerateID(),
		from: from,
		to:   to,
		time: time,
		text: text,
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
