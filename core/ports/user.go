package ports

import "context"

type Client interface {
	Close(context.Context) error
	GetRepository(name string) (UserRepository, error)
}

type UserRepository interface {
	Delete(ctx context.Context, id string) error
	Save(ctx context.Context, data any) error
	Update(ctx context.Context, id string, data any) (any, error)
	Read(ctx context.Context, id string) (any, error)
	ReadAll(ctx context.Context) (any, error)
}
