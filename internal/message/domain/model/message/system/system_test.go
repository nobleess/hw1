package system

import (
	"fmt"
	"main/internal/message/domain/model/message"
	"main/internal/user/domain/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSystemMessage(t *testing.T) {
	m := NewMessage(
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
		[]model.Login{"Recipient1", "Recipient2", "Recipient3"},
		time.Now(),
		"TextMsg",
	)
	assert.Equal(t, m.String(), fmt.Sprintf("SystemMsg[ID: %d, Time: %s, Text: %s]", m.id, m.time, m.text))
}
