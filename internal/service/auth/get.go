package auth

import (
	"context"

	"github.com/siyoga/jwt-auth-boilerplate/internal/converter"
	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"
)

func (s *service) GetByEmail(ctx context.Context, email string) (domain.User, *errors.Error) {
	tx, err := s.txRepo.StartTransaction(ctx)
	if err != nil {
		return domain.User{}, s.log.ServiceTxError(err)
	}
	defer tx.Rollback()

	existUser, err := s.userRepo.GetByEmail(ctx, tx, email)
	if err != nil {
		return domain.User{}, errors.DatabaseError(err)
	}

	if err := tx.Commit(); err != nil {
		return domain.User{}, s.log.ServiceTxError(err)
	}

	return converter.UserDomainFromModel(existUser), nil
}

func (s *service) GetById(ctx context.Context, id string) (domain.User, *errors.Error) {
	tx, err := s.txRepo.StartTransaction(ctx)
	if err != nil {
		return domain.User{}, s.log.ServiceTxError(err)
	}
	defer tx.Rollback()

	user, err := s.userRepo.GetById(ctx, tx, id)
	if err != nil {
		return domain.User{}, errors.DatabaseError(err)
	}

	if err := tx.Commit(); err != nil {
		return domain.User{}, s.log.ServiceTxError(err)
	}

	return converter.UserDomainFromModel(user), nil
}
