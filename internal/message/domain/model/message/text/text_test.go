package text

import (
	"fmt"
	"main/internal/message/domain/model/message"
	"main/internal/user/domain/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTextMessage(t *testing.T) {
	m := NewMessage(
		"testLogin",
		[]model.Login{"Recipient1", "Recipient2", "Recipient3"},
		time.Now(),
		"TextMsg",
	)
	assert.Equal(t, m.text, "TextMsg")
	assert.Equal(t, m.id, message.ID(1))
	assert.Equal(t, m.time.Second(), time.Now().Second())
}

func TestMessageString(t *testing.T) {
	m := NewMessage(
		"testLogin",
		[]model.Login{"Recipient1", "Recipient2", "Recipient3"},
		time.Now(),
		"TextMsg",
	)
	assert.Equal(t, m.String(), fmt.Sprintf("TextMsg[ID: %d, From: %s, To: %s, Time: %s, Text: %s]", m.id, m.from, m.to, m.time, m.text))
}
