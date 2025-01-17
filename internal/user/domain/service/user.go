package service

import (
	"context"
	"main/internal/user/domain/model"
	"main/internal/user/infra/postgres"
)

type UserService struct {
	repository postgres.UserRepository
}

func NewService(repository postgres.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (u UserService) Create(ctx context.Context, user model.User) error {
	if err := u.repository.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u UserService) Update(ctx context.Context, user model.User) error {
	if err := u.repository.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u UserService) Read(ctx context.Context, id model.ID) (model.User, error) {
	return u.repository.GetById(ctx, id)
}

func (u UserService) Delete(ctx context.Context, id model.ID) error {
	if err := u.repository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
