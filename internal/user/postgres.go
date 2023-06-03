package user

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

// Create implements Repository.
func (r *repository) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO public.user
		(firstname, lastname, age, email, password_hash)
	VALUES
		($1, $2, $3, $4, $5)
	RETURNING id
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	err := r.client.QueryRow(ctx, query, &user.Firstname, &user.Lastname, &user.Age, &user.Email, &user.PasswordHash).Scan(&user.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}
	return nil
}

// FindAll implements Repository.
func (r *repository) FindAll(ctx context.Context) (users []User, err error) {
	query := `
	SELECT
		id, firstname, lastname, age, email, password_hash
	FROM
		public.user
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Age, &user.Email, &user.PasswordHash)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (us User, err error) {
	query := `
	SELECT
		id, firstname, lastname, age, email, password_hash
	FROM
		public.user
	WHERE id = $1	
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	err = r.client.QueryRow(ctx, query, id).Scan(&us.Id, &us.Firstname, &us.Lastname, &us.Age, &us.Email, &us.PasswordHash)
	if err != nil {
		return User{}, err
	}
	return us, nil
}

// Update implements Repository.
func (r *repository) Update(ctx context.Context, user *User) error {
	query := `
	UPDATE
		public.user
	SET
		firstname = $1, lastname = $2, age = $3, email = $4, password_hash = $5
	WHERE
		id = $6
	`

	r.logger.Tracef("SQL Query: %s", formatQuery(query))

	_, err := r.client.Query(ctx, query, &user.Firstname, &user.Lastname, &user.Age, &user.Email, &user.PasswordHash, &user.Id)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements Repository.
func (r *repository) Delete(ctx context.Context, id string) error {
	query := `
	DELETE FROM
		public.user
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

func NewRepository(client postgresql.Client, logger *logging.Logger) Repository {
	return &repository{
		client: client,
		logger: *logger,
	}
}
