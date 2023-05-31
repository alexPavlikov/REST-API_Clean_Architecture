package book

import (
	"context"
	"strings"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/internal/author"
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
func (*repository) Create(ctx context.Context, book *Book) error {
	panic("unimplemented")
}

// Delete implements Repository.
func (*repository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindOne implements Repository.
func (*repository) FindOne(ctx context.Context, id string) (Book, error) {
	panic("unimplemented")
}

// Update implements Repository.
func (*repository) Update(ctx context.Context, book Book) error {
	panic("unimplemented")
}

// FindAll implements book.Repository MANY to MANY query
func (r *repository) FindAll(ctx context.Context) (books []Book, err error) {
	query := `
		SELECT id, name 
		FROM public.book
	`
	r.logger.Tracef("SQL Query: %s", formatQuery(query))
	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b Book
		err = rows.Scan(&b.ID, &b.Name)
		if err != nil {
			return nil, err
		}

		subQuery := `SELECT a.id, a.name
		 	FROM public.book_authors ba
		 	JOIN public.author a on a.id = ba.author_id
		 	WHERE ba.book_id = $1;`

		authorsRows, err := r.client.Query(ctx, subQuery, b.ID)
		if err != nil {
			return nil, err
		}
		authors := make([]author.Author, 0)
		for authorsRows.Next() {
			var a author.Author
			err = authorsRows.Scan(&a.ID, &a.Name)
			if err != nil {
				return nil, err
			}
			authors = append(authors, a)
		}
		b.Author = authors
		books = append(books, b)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func NewRepository(clnt postgresql.Client, lgr *logging.Logger) Repository {
	return &repository{
		client: clnt,
		logger: *lgr,
	}
}
