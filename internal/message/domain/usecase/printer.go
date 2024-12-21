package usecase

import (
	"context"
	"fmt"
	"io"

	"main/internal/message/infra/postgres"
)

type Printer struct {
	messageReposytory postgres.MessageRepository
	userRepository    postgres.UserRepository
}

func NewPrinter(messageReposytory postgres.MessageRepository, userRepository postgres.UserRepository) *Printer {
	return &Printer{
		messageReposytory: messageReposytory,
		userRepository:    userRepository,
	}

}

func (p *Printer) PrintAll(ctx context.Context, w io.Writer) error {
	users, err := p.userRepository.GetUsers(ctx)
	if err != nil {
		return err
	}
	for _, u := range users {
		msgs, err := p.messageReposytory.FindByUserId(ctx, u.ID())
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
