package author

import (
	"context"
	"fmt"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
)

type Service struct {
	repository Repository
	logger     *logging.Logger
}

func NewService(repository Repository, logger *logging.Logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]Author, error) {
	authors, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all authors due to error: %v", err)
	}
	return authors, nil
}

func (s *Service) GetOne(ctx context.Context, id string) (Author, error) {
	ath, err := s.repository.FindOne(ctx, id)
	if err != nil {
		return ath, fmt.Errorf("failed to get author due to error: %v", err)
	}
	return ath, nil
}

func (s *Service) CreateAuthor(ctx context.Context, name string) error {
	ath := Author{
		Name: name,
	}
	err := s.repository.Create(ctx, &ath)
	if err != nil {
		return fmt.Errorf("failed to create author due to error: %v", err)
	}
	return nil
}

func (s *Service) UpdateAuthour(ctx context.Context, author Author) error {
	err := s.repository.Update(ctx, author)
	if err != nil {
		return fmt.Errorf("failed to update author due to error: %v", err)
	}
	return nil
}

func (s *Service) DeleteAuthor(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete author due to error: %v", err)
	}
	return nil
}
