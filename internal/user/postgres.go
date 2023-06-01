package user

import (
	"context"
	"strings"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/client/postgresql"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger logging.Logger
}

func formatQuery(query string) string {
	return strings.ReplaceAll(strings.ReplaceAll(query, "\t", " "), "\n", " ")
}

// Create implements Repository.
func (*repository) Create(ctx context.Context, user *User) error {
	panic("unimplemented")
}

// Delete implements Repository.
func (*repository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindAll implements Repository.
func (*repository) FindAll(ctx context.Context) (users []User, err error) {
	panic("unimplemented")
}

func (r *repository) FindOne(ctx context.Context, id uint64) (us User, err error) {
	return us, err
}

// Update implements Repository.
func (*repository) Update(ctx context.Context, user User) error {
	panic("unimplemented")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) Repository {
	return &repository{
		client: client,
		logger: *logger,
	}
}
