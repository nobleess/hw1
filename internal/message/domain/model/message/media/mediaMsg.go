package media

import (
	"fmt"
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/user"
	"time"
)

// Message
type Message struct {
	id    message.ID
	from  user.Login
	to    []user.Login
	time  time.Time
	types string
	url   string
}

func NewMessage(from user.Login, to []user.Login, time time.Time, types, url string) *Message {
	return &Message{
		id:    message.GenerateID(),
		from:  from,
		to:    to,
		time:  time,
		types: types,
		url:   url,
	}
}

func (m *Message) From() user.Login {
	return m.from
}

func (m *Message) To() []user.Login {
	return m.to
}

func (m *Message) Id() message.ID {
	return m.id
}

func (m *Message) Text() string {
	return m.url
}

func (m *Message) Time() time.Time {
	return m.time
}

func (m *Message) String() string {
	return fmt.Sprintf("Message[ID: %d, From: %s, To: %s, Time: %s, Type: %s, URL: %s]", m.id, m.from, m.to, m.time, m.types, m.url)
}

func (m *Message) Redact(new message.Message) {
	m.url = new.Text() // todo return error if not allowed field change
}
