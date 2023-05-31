package user

import "context"

type Storage interface {
	FindOne(ctx context.Context, id uint64) (User, error)
}
