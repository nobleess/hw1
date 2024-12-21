package system

import (
	"fmt"
	"main/internal/message/domain/model/message"
	"main/internal/user/domain/model"
	"time"
)

// Message
type Message struct {
	id   message.ID
	to   []model.Login
	time time.Time
	text string
}

func NewMessage(to []model.Login, time time.Time, text string) *Message {
	return &Message{
		id:   message.GenerateID(),
		to:   to,
		time: time,
		text: text,
	}
}

func (m *Message) From() model.Login {
	return ""
}

func (m *Message) To() []model.Login {
	return nil
}

func (m *Message) Id() message.ID {
	return m.id
}

func (m *Message) Text() string {
	return m.text
}

func (m *Message) Time() time.Time {
	return m.time
}

func (m *Message) String() string {
	return fmt.Sprintf("SystemMsg[ID: %d, Time: %s, Text: %s]", m.id, m.time, m.text)
}

func (m *Message) Redact(new message.Message) {
	m.text = new.Text()
}
