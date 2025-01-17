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
func (u UseCase) Create(ctx context.Context, user model.User) error {
	return u.svc.Create(ctx, user)
}

func (u UseCase) Read(ctx context.Context, id model.ID) (model.User, error) {
	return u.svc.Read(ctx, id)
}

func (u UseCase) Update(ctx context.Context, user model.User) error {
	return u.svc.Update(ctx, user)
}

func (u UseCase) Delete(ctx context.Context, id model.ID) error {
	return u.svc.Delete(ctx, id)
}
