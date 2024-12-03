package user

import (
	"context"

	"github.com/siyoga/jwt-auth-boilerplate/internal/database"
	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"
	"github.com/siyoga/jwt-auth-boilerplate/internal/log"
	def "github.com/siyoga/jwt-auth-boilerplate/internal/repository"
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository/models"
)

var _ def.UserRepository = (*repository)(nil)

type (
	repository struct {
		log    log.Logger
		client *database.PostgresClient
	}
)

func NewRepository(
	log log.Logger,
	pgClient *database.PostgresClient,
) *repository {
	return &repository{
		log:    log,
		client: pgClient,
	}
}

func (r *repository) Create(ctx context.Context, tx models.Transaction, user models.User) (models.User, error) {
	query := `
		INSERT INTO users (username, email, password)
		VALUES(:username, :email, :password)
	`

	res, err := tx.Txm().NamedExecContext(ctx, query, user)
	if err != nil {
		return models.User{}, r.log.SqlError(err, errors.ErrPostgresExecRaw, query)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return models.User{}, r.log.SqlError(err, errors.ErrPostgresRowsAffectedRaw, query)
	}

	return user, err
}

func (r *repository) GetById(ctx context.Context, tx models.Transaction, id string) (models.User, error) {
	params := `WHERE id=$1`
	return r.getUserBy(ctx, tx.Txm(), params, id)
}

func (r *repository) GetByEmail(
	ctx context.Context,
	tx models.Transaction,
	email string,
) (models.User, error) {
	params := `WHERE email=$1`
	return r.getUserBy(ctx, tx.Txm(), params, email)
}

func (r *repository) CheckUserExist(
	ctx context.Context,
	tx models.Transaction,
	email, username string,
) (models.User, error) {
	params := `WHERE email=$1 OR username=$2`
	return r.getUserBy(ctx, tx.Txm(), params, email, username)
}
