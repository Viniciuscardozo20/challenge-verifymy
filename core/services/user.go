package services

import (
	"challenge-verifymy/core/models"
	"challenge-verifymy/core/ports"
	"context"
)

type UserService struct {
	repository ports.UserRepository
}

func (u *UserService) Create(ctx context.Context, user *models.UserReq) error {
	if err := user.Validate(); err != nil {
		return err
	}

	return u.Create(ctx, user)
}

func (u *UserService) Read(ctx context.Context, ID string) (*models.UserRes, error) {
	return u.Read(ctx, ID)
}

func (u *UserService) ReadAll(ctx context.Context) ([]models.UserRes, error) {
	return u.ReadAll(ctx)
}

func (u *UserService) Delete(ctx context.Context, ID string) error {
	return u.Delete(ctx, ID)
}

func (u *UserService) Update(ctx context.Context, ID string, user *models.UserReq) (*models.UserRes, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return u.Update(ctx, ID, user)
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &UserService{
		repository: repo,
	}
}
