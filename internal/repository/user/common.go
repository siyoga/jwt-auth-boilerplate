package user

import (
	"context"
	"fmt"

	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository/models"

	"github.com/jmoiron/sqlx"
)

func (r *repository) getUsersBy(ctx context.Context, executor *sqlx.Tx, params string, args ...interface{}) ([]models.User, error) {
	query := `SELECT id, username, password, email, created_at FROM users`
	query = fmt.Sprintf("%s %s", query, params)

	users := []models.User{}
	if err := sqlx.SelectContext(ctx, executor, &users, query, args...); err != nil {
		return nil, r.log.SqlError(err, errors.ErrPostgresGetRaw, query)
	}

	return users, nil
}

func (r *repository) getUserBy(ctx context.Context, executor *sqlx.Tx, params string, args ...interface{}) (models.User, error) {
	users, err := r.getUsersBy(ctx, executor, params, args...)

	if err != nil {
		return models.User{}, err
	}

	if len(users) == 0 {
		return models.User{}, nil
	}

	return users[0], err
}
