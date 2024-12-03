package repository

import (
	"context"

	"github.com/siyoga/jwt-auth-boilerplate/internal/repository/models"
)

type (
	TxRepository interface {
		StartTransaction(ctx context.Context) (*models.Tx, error)
	}

	JwtRepository interface {
		AddToken(ctx context.Context, tx models.Transaction, token models.Token) (models.Token, error)
		CheckToken(ctx context.Context, tx models.Transaction, number int64, userId string) (models.Token, error)
		DropToken(ctx context.Context, tx models.Transaction, userId string, number int64) error
		FindNumber(ctx context.Context, tx models.Transaction, userId string) (int64, error)
	}

	UserRepository interface {
		Create(ctx context.Context, tx models.Transaction, user models.User) (models.User, error)
		GetById(ctx context.Context, tx models.Transaction, userId string) (models.User, error)
		GetByEmail(ctx context.Context, tx models.Transaction, email string) (models.User, error)
		CheckUserExist(ctx context.Context, tx models.Transaction, email, username string) (models.User, error)
	}
)
