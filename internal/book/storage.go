package book

import "context"

type Repository interface {
	Create(ctx context.Context, book *Book) error
	FindAllBooks(ctx context.Context) (books []Book, err error)
	FindOne(ctx context.Context, id string) (Book, error)
	FindAllBooksByAuthor(ctx context.Context, id string) (a Author, err error)
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, id string) error
}
