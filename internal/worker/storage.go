package worker

import "context"

type Repository interface {
	Create(ctx context.Context, worker *Workrer) error
	FindAll(ctx context.Context) (workers []Workrer, err error)
	FindOne(ctx context.Context, id string) (Workrer, error)
	Update(ctx context.Context, worker *Workrer) error
	Delete(ctx context.Context, id string) error
}
