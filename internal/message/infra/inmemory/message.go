package inmemory

import (
	"errors"
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/user"
)

type MessageStorage struct {
	storage map[user.Login][]message.Message
}

func NewMessage(storage map[user.Login][]message.Message) *MessageStorage {
	return &MessageStorage{
		storage: storage,
	}
}

func (s *MessageStorage) GetUserMessages(login user.Login) ([]message.Message, error) {
	if msg := s.storage[login]; msg != nil {
		return msg, nil
	}
	return nil, errors.New("messages not found")
}

func (s *MessageStorage) Update(msg message.Message) error {
	for _, recipient := range msg.To() {
		if msgs, _ := s.storage[recipient]; msgs != nil {
			for _, m := range msgs {
				if m.Id() == msg.Id() {
					m.Redact(msg)
					break
				}
			}
		} else {
			return errors.New("messages not found")
		}
	}
	return nil
}

func (s *MessageStorage) CreateMessage(msg message.Message) error {
	for _, recipient := range msg.To() {
		if msgs, _ := s.storage[recipient]; msgs == nil {
			s.storage[recipient] = []message.Message{msg}
		} else {
			msgs = append(msgs, msg)
		}
	}
	return nil
}
