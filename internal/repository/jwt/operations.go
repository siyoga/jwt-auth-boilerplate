package jwt

import (
	"context"
	"database/sql"

	"github.com/siyoga/jwt-auth-boilerplate/internal/database"
	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"
	"github.com/siyoga/jwt-auth-boilerplate/internal/log"
	def "github.com/siyoga/jwt-auth-boilerplate/internal/repository"
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository/models"

	"github.com/jmoiron/sqlx"
)

var _ def.JwtRepository = (*repository)(nil)

type repository struct {
	log    log.Logger
	client *database.PostgresClient
}

func NewRepository(
	log log.Logger,
	pgClient *database.PostgresClient,
) *repository {
	return &repository{
		log:    log,
		client: pgClient,
	}
}

func (r *repository) AddToken(
	ctx context.Context, tx models.Transaction,
	token models.Token,
) (models.Token, error) {
	query := `
    INSERT INTO refresh_tokens (user_id, number, payload, expires_at)
    VALUES(:user_id, :number, :payload, :expires_at)
  `

	res, err := tx.Txm().NamedExecContext(ctx, query, token)
	if err != nil {
		return models.Token{}, r.log.SqlError(err, errors.ErrPostgresExecRaw, query)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return models.Token{}, r.log.SqlError(err, errors.ErrPostgresRowsAffectedRaw, query)
	}

	return token, nil
}

func (r *repository) DropToken(
	ctx context.Context,
	tx models.Transaction,
	userId string,
	number int64,
) error {
	query := `DELETE FROM refresh_tokens WHERE user_id=$1 AND number=$2`
	_, err := tx.Txm().ExecContext(ctx, query, userId, number)
	if err != nil {
		return r.log.SqlError(err, errors.ErrPostgresExecRaw, query)
	}

	return nil
}

func (r *repository) CheckToken(
	ctx context.Context,
	tx models.Transaction,
	number int64,
	userId string,
) (models.Token, error) {
	query := `
		SELECT user_id, number, expires_at
		FROM refresh_tokens
		WHERE number=$1 AND user_id=$2
	`

	tokens := []models.Token{}
	if err := sqlx.SelectContext(ctx, tx.Txm(), &tokens, query, number, userId); err != nil {
		return models.Token{}, r.log.SqlError(err, errors.ErrPostgresGetRaw, query)
	}

	if len(tokens) == 0 {
		return models.Token{}, errors.ErrPostgresNotExists
	}

	return tokens[0], nil
}

func (r *repository) FindNumber(ctx context.Context, tx models.Transaction, userId string) (int64, error) {
	var numbers []int64
	query := `
    SELECT number
    FROM refresh_tokens
    WHERE user_id = $1
    ORDER BY number
  `

	if err := tx.Txm().SelectContext(ctx, &numbers, query, userId); err != nil {
		if err == sql.ErrNoRows {
			numbers = []int64{}
		}

		return 0, r.log.SqlError(err, errors.ErrPostgresQueryRowRaw, query)
	}

	return r.findNumbers(numbers)
}
