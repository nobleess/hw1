package usecase

import (
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/user"
)

type Storage interface {
	GetUserMessage(user.Login) ([]message.Message, error)
	UpdateByID(message.ID, message.Message) error
	Update(message.Message) error
}

type Redactor struct {
	storage Storage
}

func NewRedactor(storage Storage) *Redactor {
	return &Redactor{storage: storage}
}

func (svc *Redactor) Redact(m message.Message) error {
	if err := svc.storage.Update(m); err != nil {
		return err
	}
	return nil
}
