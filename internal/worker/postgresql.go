package worker

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/client/postgresql"
	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
	"github.com/jackc/pgconn"
)

// Create(ctx context.Context, wokrer *Workrer) error
// 	FindAll(ctx context.Context) (wokrers []Workrer, err error)
// 	FindOne(ctx context.Context, id string) (Workrer, error)
// 	Update(ctx context.Context, wokrer Workrer) error
// 	Delete(ctx context.Context, id string) error

type repository struct {
	client postgresql.Client
	logger logging.Logger
}

func NewRepository(cl postgresql.Client, lg *logging.Logger) Repository {
	return &repository{
		client: cl,
		logger: *lg,
	}
}

func formatQuery(query string) string {
	return strings.ReplaceAll(strings.ReplaceAll(query, "\t", " "), "\n", " ")
}

func (r *repository) Create(ctx context.Context, worker *Workrer) error {
	query := `
	INSERT INTO public.worker
		(firstname, lastname, age, experiens, number, address, email, password_hash)	
	VALUES
		($1 ,$2, $3, $4, $5, $6, $7, $8)
	RETURNING id	
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	err := r.client.QueryRow(ctx, query, worker.Firstname, worker.Lastname, worker.Age, worker.Experieons, worker.Number, worker.Address, worker.Email, worker.PasswordHash).Scan(&worker.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState:= %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}
	return nil
}
func (r *repository) FindAll(ctx context.Context) (workers []Workrer, err error) {
	query := `
	SELECT 
		id, firstname, lastname, age, experiens, number, address, email, password_hash
	FROM 
		public.worker
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var worker Workrer
		err = rows.Scan(&worker.Id, &worker.Firstname, &worker.Lastname, &worker.Age, &worker.Experieons, &worker.Number, &worker.Address, &worker.Email, &worker.PasswordHash)
		if err != nil {
			return nil, err
		}
		workers = append(workers, worker)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return workers, nil
}
func (r *repository) FindOne(ctx context.Context, id string) (Workrer, error) {
	query := `
	SELECT 
		id, firstname, lastname, age, experiens, number, address, email, password_hash
	FROM 
		public.worker
	WHERE 
		id = $1
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	var worker Workrer

	err := r.client.QueryRow(ctx, query, id).Scan(&worker.Id, &worker.Firstname, &worker.Lastname, &worker.Age, &worker.Experieons, &worker.Number, &worker.Address, &worker.Email, &worker.PasswordHash)
	if err != nil {
		return Workrer{}, err
	}
	return worker, nil
}
func (r *repository) Update(ctx context.Context, worker *Workrer) error {
	query := `
	UPDATE INTO public.worker
	SET 
		firstname = $1, lastname = $2, age = $3, experiens = $4, number = $5, address = $6, email = $7, password_hash = $8
	WHERE 
		id = $9	
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	_, err := r.client.Query(ctx, query, &worker.Firstname, &worker.Lastname, &worker.Age, &worker.Experieons, &worker.Number, &worker.Address, &worker.Email, &worker.PasswordHash)
	if err != nil {
		return err
	}
	return nil

}
func (r *repository) Delete(ctx context.Context, id string) error {
	query := `
	DELETE FROM
		public.worker
	WHERE
		id = $1
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	_, err := r.client.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
