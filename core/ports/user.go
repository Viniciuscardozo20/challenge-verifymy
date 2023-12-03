package ports

import (
	"challenge-verifymy/core/models"
	"context"
)

type Database interface {
	Disconnected() <-chan struct{}
	GetRepository(ctx context.Context, name string) (UserRepository, error)
}

type UserRepository interface {
	Delete(ctx context.Context, id string) error
	Save(ctx context.Context, data, output any) error
	Update(ctx context.Context, id string, data, output any) error
	Read(ctx context.Context, id string, output any) error
	ReadAll(ctx context.Context, output any) error
}

type UserService interface {
	Create(ctx context.Context, user *models.UserReq) (res *models.UserRes, err error)
	ReadAll(ctx context.Context) (*[]models.UserRes, error)
	Read(ctx context.Context, ID string) (*models.UserRes, error)
	Update(ctx context.Context, ID string, user *models.UserReq) (*models.UserRes, error)
	Delete(ctx context.Context, ID string) error
}
