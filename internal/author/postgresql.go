package author

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/client/postgresql"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
	logger logging.Logger
}

func formatQuery(query string) string {
	return strings.ReplaceAll(strings.ReplaceAll(query, "\t", " "), "\n", " ")
}

// Create implements author.Repository
func (r *repository) Create(ctx context.Context, author *Author) error {
	query := `
	INSERT INTO public.author 
		(name) 
	VALUES 
		($1) 
	RETURNING id
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	err := r.client.QueryRow(ctx, query, author.Name).Scan(&author.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState:= %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}

	return nil
}

// FindAll implements author.Repository
func (r *repository) FindAll(ctx context.Context) (authors []Author, err error) {
	query := `
		SELECT 
			id, name 
		FROM 
			public.author
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a Author
		err = rows.Scan(&a.ID, &a.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, a)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return authors, nil
}

// FindOne implements author.Repository
func (r *repository) FindOne(ctx context.Context, id string) (a Author, err error) {
	query := `
		SELECT 
			id, name 
		FROM 
			public.author
		WHERE 
			id = $1
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	err = r.client.QueryRow(ctx, query, id).Scan(&a.ID, &a.Name)
	if err != nil {
		return Author{}, err
	}

	return a, nil
}

// Update implements author.Repository
func (r *repository) Update(ctx context.Context, author Author) error {
	query := `
		UPDATE 
			public.author
		SET 
			name = ($1)
		WHERE 
			id = ($2)
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	_, err := r.client.Query(ctx, query, author.Name)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements author.Repository
func (r *repository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM 
			public.author
		WHERE 
			id = ($1)
	`

	r.logger.Tracef("SQL Query: %v", formatQuery(query))

	_, err := r.client.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func NewRepository(clnt postgresql.Client, lgr *logging.Logger) Repository {
	return &repository{
		client: clnt,
		logger: *lgr,
	}
}
