package inmemory

import (
	"errors"
	model2 "main/internal/message/domain/model"
	"main/internal/user/domain/model"
)

type MessageStorage struct {
	storage map[model.Login][]model2.Message
}

func NewMessage(storage map[model.Login][]model2.Message) *MessageStorage {
	return &MessageStorage{
		storage: storage,
	}
}

func (s *MessageStorage) GetUserMessages(login model.Login) ([]model2.Message, error) {
	if msg := s.storage[login]; msg != nil {
		return msg, nil
	}
	return nil, errors.New("messages not found")
}

func (s *MessageStorage) Update(msg model2.Message) error {
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

func (s *MessageStorage) CreateMessage(msg model2.Message) error {
	for _, recipient := range msg.To() {
		if msgs, _ := s.storage[recipient]; msgs == nil {
			s.storage[recipient] = []model2.Message{msg}
		} else {
			msgs = append(msgs, msg)
		}
	}
	return nil
}
