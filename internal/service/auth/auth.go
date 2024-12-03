package auth

import (
	"context"

	"github.com/siyoga/jwt-auth-boilerplate/internal/converter"
	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"
	"golang.org/x/crypto/bcrypt"
)

// Вход в систему
func (s service) Auth(ctx context.Context, ip string, number *int64, email string, password string) (domain.Token, domain.Token, *errors.Error) {
	tx, err := s.txRepo.StartTransaction(ctx)
	if err != nil {
		return domain.Token{}, domain.Token{}, s.log.ServiceTxError(err)
	}
	defer tx.Rollback()

	u, err := s.userRepo.GetByEmail(ctx, tx, email)
	if err != nil {
		return domain.Token{}, domain.Token{}, errors.DatabaseError(err)
	}
	user := converter.UserDomainFromModel(u)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(password)); err != nil {
		s.log.Error(errors.ErrAuthHashPasswordRaw, err)
		return domain.Token{}, domain.Token{}, errors.ErrInvalidCredentials
	}

	_, at, rt, e := s.createTokensTx(ctx, tx, user.Id, ip, number)
	if e != nil {
		return domain.Token{}, domain.Token{}, e
	}

	if err := tx.Commit(); err != nil {
		return domain.Token{}, domain.Token{}, s.log.ServiceTxError(err)
	}

	return at, rt, nil
}
