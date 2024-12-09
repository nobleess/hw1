package usecase

import (
	"main/internal/message/domain/model/message"
)

type Storage interface {
	Create(message.Message) error
}

type Sender struct {
	storage Storage
}

func NewSender(storage Storage) *Sender {
	return &Sender{
		storage: storage,
	}
}

func (svc *Sender) Send(m message.Message) error {
	if err := svc.storage.Create(m); err != nil {
		return err
	}
	return nil
}
