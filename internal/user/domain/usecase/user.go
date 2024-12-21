package usecase

import (
	"context"
	"main/internal/user/domain/model"
	"main/internal/user/domain/service"
)

type UseCase struct {
	svc service.UserService
}

func NewUseCase(svc service.UserService) *UseCase {
	return &UseCase{svc: svc}
}

func (u UseCase) Create(ctx context.Context, user model.User) (*model.User, error) {

	u.svc.Cre

}

func (u UseCase) Read(ctx context.Context, userID string) (*model.User, error) {

}

func (u UseCase) Update(ctx context.Context, userID string) (*model.User, error) {

}
func (u UseCase) Delete(ctx context.Context, userID string) (*model.User, error) {

}
