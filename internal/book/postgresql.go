package book

import (
	"context"
	"fmt"
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
func (r *repository) Create(ctx context.Context, book *Book) error {
	query := `
	INSERT INTO public.book
		(name, author)
	VALUES
		($1, $2)
	RETURNING id
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	err := r.client.QueryRow(ctx, query, &book.Name, &book.Author).Scan(&book.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements Repository.
func (r *repository) Delete(ctx context.Context, id string) error {
	query := `
	DELETE FROM 
		public.book
	WHERE 
		id = $1
	`
	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	_, err := r.client.Query(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// FindOne implements Repository.
func (r *repository) FindOne(ctx context.Context, id string) (Book, error) {
	query := `
	SELECT 
		id, name
	FROM 
		public.book
	WHERE 
		id = $1
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	var book Book

	rows, err := r.client.Query(ctx, query, id)
	if err != nil {
		return Book{}, err
	}

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Name)
		if err != nil {
			return Book{}, err
		}

		subQuery := `SELECT a.id, a.name
		 	FROM public.book_authors ba
		 	JOIN public.author a on a.id = ba.author_id
		 	WHERE ba.book_id = $1;`

		authorsRows, err := r.client.Query(ctx, subQuery, book.ID)
		if err != nil {
			return Book{}, err
		}
		authors := make([]author.Author, 0)
		for authorsRows.Next() {
			var a author.Author
			err = authorsRows.Scan(&a.ID, &a.Name)
			if err != nil {
				return Book{}, err
			}
			authors = append(authors, a)
		}
		book.Author = authors
	}
	return book, nil
}

// Update implements Repository.
func (r *repository) Update(ctx context.Context, book *Book) error {
	query := `
	UPDATE INTO 
		public.book
	SET 
		name = $1, author = $2
	WHERE id = $3
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	_, err := r.client.Query(ctx, query, &book.Name, &book.Author, &book.ID)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements book.Repository MANY to MANY query
func (r *repository) FindAllBooks(ctx context.Context) (books []Book, err error) {
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

func (r *repository) FindAllBooksByAuthor(ctx context.Context, id string) (a Author, err error) {
	query := `
		SELECT id, name 
		FROM public.author
		WHERE id = $1
	`
	r.logger.Tracef("SQL Query: %s", formatQuery(query))
	rows, err := r.client.Query(ctx, query, id)
	if err != nil {
		return Author{}, err
	}
	for rows.Next() {
		err = rows.Scan(&a.ID, &a.Name)
		if err != nil {
			return Author{}, err
		}

		subQuery := `
		SELECT 
			ba.id, b.name, b.id
		FROM 
			public.book_authors ba
		JOIN public.book 
			b on b.id = ba.book_id
		WHERE 
			ba.author_id = $1;`

		authorsRows, err := r.client.Query(ctx, subQuery, id)
		if err != nil {
			return Author{}, err
		}
		books := make([]Book, 0)
		for authorsRows.Next() {
			var b Book
			var bookId string
			err = authorsRows.Scan(&b.ID, &b.Name, &bookId)
			if err != nil {
				return Author{}, err
			}
			fmt.Println(b.ID)
			b.Author, err = r.getAllAuthorByBookId(ctx, bookId)
			if err != nil {
				return Author{}, err
			}
			books = append(books, b)
		}
		a.Books = books
	}
	err = rows.Err()
	if err != nil {
		return Author{}, err
	}
	return a, nil
}

func NewRepository(clnt postgresql.Client, lgr *logging.Logger) Repository {
	return &repository{
		client: clnt,
		logger: *lgr,
	}
}

func (r *repository) getAllAuthorByBookId(ctx context.Context, id string) (arr []author.Author, err error) {
	query := `
	SELECT 
			ba.id, b.name
		FROM 
			public.book_authors ba
		JOIN public.author 
			b on b.id = ba.author_id
		WHERE 
			ba.book_id = $1;
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	rows, err := r.client.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a author.Author
		err = rows.Scan(&a.ID, &a.Name)
		if err != nil {
			return nil, err
		}
		arr = append(arr, a)
	}
	return arr, nil
}
