package user

import (
	"context"
	"fmt"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
)

type Service struct {
	loger      *logging.Logger
	repository Repository
}

func NewService(logger *logging.Logger, repository Repository) *Service {
	return &Service{
		loger:      logger,
		repository: repository,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]User, error) {
	users, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users, due to err: %v", err)
	}
	return users, nil
}

func (s *Service) GetOne(ctx context.Context, id string) (User, error) {
	user, err := s.repository.FindOne(ctx, id)
	if err != nil {
		return User{}, fmt.Errorf("failed to get one user, due to err: %v", err)
	}
	return user, nil
}

func (s *Service) Create(ctx context.Context, user *User) error {
	err := s.repository.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user, due to err: %v", err)
	}
	return nil
}

func (s *Service) Update(ctx context.Context, user *User) error {
	err := s.repository.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to update user, due to err: %v", err)
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user, due to err: %v", err)
	}
	return nil
}
