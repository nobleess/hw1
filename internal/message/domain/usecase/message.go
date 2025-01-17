package usecase

import (
	"context"
	"main/internal/channel/channel"
	"main/internal/message/domain/model"
	"main/internal/message/domain/service"
	user "main/internal/user/domain/model"
)

type Usecase struct {
	svc service.MessageService
}

func NewUsecase(svc service.MessageService) *Usecase {
	return &Usecase{svc: svc}
}

func (u Usecase) Send(ctx context.Context, message model.Message) error {
	// todo
	//func (u Usecase) SendBatch(ctx context.Context,message []model.Message) error{}

	return u.svc.Create(ctx, message)

}

func (u Usecase) Update(ctx context.Context, message model.Message) error {
	return u.svc.Update(ctx, message)
}

func (u Usecase) Delete(ctx context.Context, id model.ID) error {
	return u.svc.Delete(ctx, id)
}

func (u Usecase) GetChannelMessages(ctx context.Context, id channel.ID) ([]model.Message, error) {
	return u.svc.GetChannelMessages(ctx, id)
}

func (u Usecase) GetUserMessages(ctx context.Context, id user.ID) ([]model.Message, error) {
	return u.svc.GetUserMessages(ctx, id)
}
