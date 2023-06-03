package book

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

func (s *Service) GetAll(ctx context.Context) ([]Book, error) {
	books, err := s.repository.FindAllBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all books, due to err: %v", err)
	}
	return books, nil
}

func (s *Service) GetOne(ctx context.Context, id string) (Book, error) {
	book, err := s.repository.FindOne(ctx, id)
	if err != nil {
		return Book{}, fmt.Errorf("failed to get one book, due to err: %v", err)
	}
	return book, nil
}

func (s *Service) GetAllByAuthor(ctx context.Context, id string) (Author, error) {
	author, err := s.repository.FindAllBooksByAuthor(ctx, id)
	if err != nil {
		return Author{}, fmt.Errorf("failed to get all books by author id, due to err: %v", err)
	}
	return author, nil
}

func (s *Service) CreateBook(ctx context.Context, book *Book) error {
	err := s.repository.Create(ctx, book)
	if err != nil {
		return fmt.Errorf("failed to create book, due to err: %v", err)
	}
	return nil
}

func (s *Service) UpdateBook(ctx context.Context, book *Book) error {
	err := s.repository.Update(ctx, book)
	if err != nil {
		return fmt.Errorf("failed to update book, due to err: %v", err)
	}
	return nil
}

func (s *Service) DeleteBook(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to deleted book, due to err: %v", err)
	}
	return nil
}
