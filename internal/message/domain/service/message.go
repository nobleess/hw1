package service

import (
	"context"
	"main/internal/channel/channel"
	"main/internal/message/domain/model"
	"main/internal/message/infra/postgres"
	"main/internal/message/infra/postgres/dto"
	user "main/internal/user/domain/model"
)

type MessageService struct {
	repository postgres.MessageRepository
}

func NewMessageService(repository postgres.MessageRepository) *MessageService {
	return &MessageService{
		repository: repository,
	}
}

func (s MessageService) Create(ctx context.Context, message model.Message) error {
	return s.repository.Create(ctx, dto.NewMessage(message))
}

func (s MessageService) GetChannelMessages(ctx context.Context, id channel.ID) ([]model.Message, error) {
	return s.repository.FindByChannelId(ctx, id)
}

func (s MessageService) GetUserMessages(ctx context.Context, id user.ID) ([]model.Message, error) {
	return s.repository.FindByUserId(ctx, id)
}

func (s MessageService) Delete(ctx context.Context, id model.ID) error {
	return nil
}

//func (s MessageService) Read(ctx context.Context, id model.ID) (model.Message, error) {
//
//}

func (s MessageService) Update(ctx context.Context, message model.Message) error {
	return s.repository.Update(ctx, dto.NewMessage(message))
}
