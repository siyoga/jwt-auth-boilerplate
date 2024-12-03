package auth

import (
	"context"

	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"
)

func (s service) Verify(ctx context.Context, token string, purpose domain.AuthPurpose) (domain.Account, int64, *errors.Error) {
	tx, err := s.txRepo.StartTransaction(ctx)
	if err != nil {
		return domain.Account{}, 0, s.log.ServiceTxError(err)
	}
	defer tx.Rollback()

	t, e := s.parseToken(token, purpose)
	if e != nil {
		return domain.Account{}, 0, e
	}

	acc, number, e := s.checkToken(ctx, tx, t)
	if e != nil {
		return domain.Account{}, 0, e
	}

	if err := tx.Commit(); err != nil {
		return domain.Account{}, 0, s.log.ServiceTxError(err)
	}

	return acc, number, nil
}
