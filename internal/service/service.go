package service

import (
	"context"

	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"
)

type (
	AuthService interface {
		Create(ctx context.Context, user domain.User) *errors.Error
		Auth(ctx context.Context, ip string, number *int64, email string, password string) (domain.Token, domain.Token, *errors.Error)
		GetByEmail(ctx context.Context, email string) (domain.User, *errors.Error)
		GetById(ctx context.Context, id string) (domain.User, *errors.Error)
		Verify(ctx context.Context, token string, purpose domain.AuthPurpose) (domain.Account, int64, *errors.Error)
		CreateTokens(ctx context.Context, ip string, userId string, number int64) (domain.Token, domain.Token, *errors.Error)
	}
)
