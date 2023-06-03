package worker

import (
	"context"
	"fmt"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
)

type Service struct {
	logger     *logging.Logger
	repository Repository
}

func NewService(logger *logging.Logger, repository Repository) *Service {
	return &Service{
		logger:     logger,
		repository: repository,
	}
}

func (s *Service) CreateWorker(ctx context.Context, worker *Workrer) error {
	err := s.repository.Create(ctx, worker)
	if err != nil {
		return fmt.Errorf("failed to create worker, due to err: %v", err)
	}
	return nil
}
func (s *Service) GetAll(ctx context.Context) (workers []Workrer, err error) {
	workers, err = s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all worker, due to err: %v", err)
	}
	return workers, nil
}
func (s *Service) GetOne(ctx context.Context, id string) (Workrer, error) {
	worker, err := s.repository.FindOne(ctx, id)
	if err != nil {
		return Workrer{}, fmt.Errorf("failed to get one worker, due to err: %v", err)
	}
	return worker, nil
}
func (s *Service) UpdateWorker(ctx context.Context, worker *Workrer) error {
	err := s.repository.Update(ctx, worker)
	if err != nil {
		return fmt.Errorf("failed to update worker, due to err: %v", err)
	}
	return nil
}
func (s *Service) DeleteWorker(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete worker, due to err: %v", err)
	}
	return nil
}
