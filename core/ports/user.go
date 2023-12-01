package ports

import (
	"challenge-verifymy/core/models"
	"context"
)

type Database interface {
	Disconnected() <-chan struct{}
	GetRepository(name string) UserRepository
}

type UserRepository interface {
	Delete(ctx context.Context, id string) error
	Save(ctx context.Context, data any) error
	Update(ctx context.Context, id string, data any) error
	Read(ctx context.Context, id string, output any) error
	ReadAll(ctx context.Context, output any) error
}

type UserService interface {
	Create(ctx context.Context, user *models.UserReq) error
	ReadAll(ctx context.Context) ([]models.UserRes, error)
	Read(ctx context.Context, ID string) (*models.UserRes, error)
	Update(ctx context.Context, ID string, user *models.UserReq) (*models.UserRes, error)
	Delete(ctx context.Context, ID string) error
}
