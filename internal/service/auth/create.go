package auth

import (
	"context"
	"fmt"

	"github.com/siyoga/jwt-auth-boilerplate/internal/converter"
	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofrs/uuid"
)

func (s *service) Create(ctx context.Context, newUser domain.User) *errors.Error {
	tx, err := s.txRepo.StartTransaction(ctx)
	if err != nil {
		return s.log.ServiceTxError(err)
	}
	defer tx.Rollback()

	existUser, err := s.userRepo.CheckUserExist(ctx, tx, newUser.Email, newUser.Username)
	if err != nil {
		return errors.DatabaseError(err)
	}

	if !converter.UserDomainFromModel(existUser).IsEmpty() {
		return errors.WD(errors.ErrConflict, fmt.Errorf("user with this credentials already exists"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Pass), bcrypt.DefaultCost)
	if err != nil {
		return s.log.ServiceError(errors.WD(errors.ErrInternal, err))
	}

	newUser = domain.User{
		Username: newUser.Username,
		Pass:     string(hash),
		Email:    newUser.Email,
	}

	if _, err := s.userRepo.Create(ctx, tx, converter.UserModelFromDomain(newUser)); err != nil {
		return errors.DatabaseError(err)
	}

	if err := tx.Commit(); err != nil {
		return s.log.ServiceTxError(err)
	}

	return nil
}

func (s service) CreateTokens(ctx context.Context, ip string, userId string, number int64) (domain.Token, domain.Token, *errors.Error) {
	tx, err := s.txRepo.StartTransaction(ctx)
	if err != nil {
		return domain.Token{}, domain.Token{}, s.log.ServiceTxError(err)
	}
	defer tx.Rollback()

	id, err := uuid.FromString(userId)
	if err != nil {
		return domain.Token{}, domain.Token{}, s.log.ServiceError(errors.WD(errors.ErrInternal, err))
	}

	_, at, rt, e := s.createTokensTx(ctx, tx, id, ip, &number)
	if e != nil {
		return domain.Token{}, domain.Token{}, e
	}

	if err := tx.Commit(); err != nil {
		return domain.Token{}, domain.Token{}, s.log.ServiceTxError(err)
	}

	return at, rt, nil
}
