package message

import (
	"github.com/google/uuid"
	"sync/atomic"
)

var msgGID atomic.Uint64 = atomic.Uint64{} // Счётчик для уникальных ID

type ID uuid.UUID

type Message interface {
	Id() ID
}

// Функция для генерации уникального ID
// func GenerateID() ID {
// 	return ID(msgGID.Add(1))
// }
