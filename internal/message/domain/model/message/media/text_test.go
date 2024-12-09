package media

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/user"
	"testing"
	"time"
)

func TestNewMediaMessage(t *testing.T) {
	m := NewMessage(
		"testLogin",
		[]user.Login{"Recipient1", "Recipient2", "Recipient3"},
		time.Now(),
		"mp3",
		"https://domain/filepath",
	)

	assert.Equal(t, m.types, "mp3")
	assert.Equal(t, m.url, "https://domain/filepath")
	assert.Equal(t, m.id, message.ID(1))
	assert.Equal(t, m.time.Second(), time.Now().Second())
}

func TestMessageString(t *testing.T) {
	m := NewMessage(
		"testLogin",
		[]user.Login{"Recipient1", "Recipient2", "Recipient3"},
		time.Now(),
		"mp3",
		"https://domain/filepath",
	)
	assert.Equal(t, m.String(), fmt.Sprintf("Message[ID: %d, From: %s, To: %s, Time: %s, Type: %s, URL: %s]", m.id, m.from, m.to, m.time, m.types, m.url))
}
