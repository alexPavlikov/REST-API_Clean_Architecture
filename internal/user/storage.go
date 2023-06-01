package user

import "context"

type Repository interface {
	Create(ctx context.Context, user *User) error
	FindAll(ctx context.Context) (users []User, err error)
	FindOne(ctx context.Context, id uint64) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
