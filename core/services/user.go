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
	return u.repository.Save(ctx, user)
}

func (u *UserService) Read(ctx context.Context, ID string) (res *models.UserRes, err error) {
	res = &models.UserRes{}

	if err = u.repository.Read(ctx, ID, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserService) ReadAll(ctx context.Context) (res *[]models.UserRes, err error) {
	res = &[]models.UserRes{}

	if err = u.repository.ReadAll(ctx, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserService) Delete(ctx context.Context, ID string) error {
	return u.repository.Delete(ctx, ID)
}

func (u *UserService) Update(ctx context.Context, ID string, user *models.UserReq) (res *models.UserRes, err error) {
	res = &models.UserRes{}

	if err = u.repository.Update(ctx, ID, user, res); err != nil {
		return nil, err
	}

	return res, nil
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &UserService{
		repository: repo,
	}
}
