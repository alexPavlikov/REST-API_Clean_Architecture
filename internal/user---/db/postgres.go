package db

import (
	"context"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/user---"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
)

type db struct {
	storage user.Storage
	logger  *logging.Logger
}

func (s *db) FindOne(ctx context.Context, id uint64) (us user.User, err error) {
	return us, err
}
