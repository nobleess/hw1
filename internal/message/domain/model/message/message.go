package message

import (
	"main/internal/message/domain/model/user"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

var msgGID atomic.Uint64 = atomic.Uint64{} // Счётчик для уникальных ID

type ID uuid.UUID

type Message interface {
	From() user.Login
	To() []user.Login
	Id() ID
	Time() time.Time
	Text() string
	String() string
	Redact(Message)
}

// Функция для генерации уникального ID
// func GenerateID() ID {
// 	return ID(msgGID.Add(1))
// }
