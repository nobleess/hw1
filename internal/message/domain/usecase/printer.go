package usecase

import (
	"fmt"
	"io"
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/user"
)

type Storage interface {
	GetUserMessages(user.Login) ([]message.Message, error)
	GetUsers() []user.User
}

type Printer struct {
	storage Storage
}

func NewPrinter(storage Storage) *Printer {
	return &Printer{
		storage: storage,
	}

}

func (p *Printer) Print(w io.Writer) error {
	for _, u := range p.storage.GetUsers() {
		msgs, err := p.storage.GetUserMessages(u.Login())
		if err != nil {
			return err
		}
		for _, m := range msgs {
			if _, err := w.Write([]byte(m.String())); err != nil {
				return fmt.Errorf("cant write: %w", err)
			}
		}
	}
	return nil
}
