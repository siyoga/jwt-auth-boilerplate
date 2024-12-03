package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository/models"
)

//go:generate  mockgen -source repository.go -package mocks -destination ../mocks/repository.go
type (
	Transaction interface {
		Commit() error
		Rollback()
		Txm() *sqlx.Tx
	}

	TxRepository interface {
		StartTransaction(ctx context.Context) (*models.Tx, error)
	}

	JwtRepository interface {
		AddToken(ctx context.Context, tx Transaction, token models.Token) (models.Token, error)
		CheckToken(ctx context.Context, tx Transaction, number int64, userId string) (models.Token, error)
		DropToken(ctx context.Context, tx Transaction, userId string, number int64) error
		FindNumber(ctx context.Context, tx Transaction, userId string) (int64, error)
	}

	UserRepository interface {
		Create(ctx context.Context, tx Transaction, user models.User) (models.User, error)
		GetById(ctx context.Context, tx Transaction, userId string) (models.User, error)
		GetByEmail(ctx context.Context, tx Transaction, email string) (models.User, error)
		CheckUserExist(ctx context.Context, tx Transaction, email, username string) (models.User, error)
	}
)
