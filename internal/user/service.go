package user

import (
	"context"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
)

type Service struct {
	loger      logging.Logger
	repository Repository
}

func NewService(logger logging.Logger, repository Repository) *Service {
	return &Service{
		loger:      logger,
		repository: repository,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]User, error) {

}

func (s *Service) GetOne(ctx context.Context, id string) (User, error) {

}

func (s *Service) CreateAuthor(ctx context.Context, name string) error {

}

func (s *Service) UpdateAuthour(ctx context.Context, user User) error {

}

func (s *Service) DeleteAuthor(ctx context.Context, id string) error {

}
