package system

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/user"
	"testing"
	"time"
)

func TestNewSystemMessage(t *testing.T) {
	m := NewMessage(
		[]user.Login{"Recipient1", "Recipient2", "Recipient3"},
		time.Now(),
		"TextMsg",
	)
	assert.Equal(t, m.text, "TextMsg")
	assert.Equal(t, m.id, message.ID(1))
	assert.Equal(t, m.time.Second(), time.Now().Second())
}

func TestMessageString(t *testing.T) {
	m := NewMessage(
		[]user.Login{"Recipient1", "Recipient2", "Recipient3"},
		time.Now(),
		"TextMsg",
	)
	assert.Equal(t, m.String(), fmt.Sprintf("SystemMsg[ID: %d, Time: %s, Text: %s]", m.id, m.time, m.text))
}
