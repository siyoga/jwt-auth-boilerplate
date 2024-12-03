package auth

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/siyoga/jwt-auth-boilerplate/internal/converter"
	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

// Регистрация

func (s *service) createTokensTx(
	ctx context.Context, tx models.Transaction, id uuid.UUID, ip string, number *int64,
) (int64, domain.Token, domain.Token, *errors.Error) {
	now := s.timeAdapter.Now()

	if number == nil {
		num, err := s.jwtRepo.FindNumber(ctx, tx, id.String())
		if err != nil {
			return 0, domain.Token{}, domain.Token{}, errors.DatabaseError(err)
		}
		number = &num
	}

	accessExpiresAt, refreshExpiresAt := now.Add(s.accessTokenTimeout), now.Add(s.refreshTokenTimeout)
	accessTokenHash, e := s.generateToken(ctx, tx, domain.PurposeAccess, *number, id, ip, accessExpiresAt)
	if e != nil {
		return 0, domain.Token{}, domain.Token{}, e
	}
	refreshTokenHash, e := s.generateToken(ctx, tx, domain.PurposeRefresh, *number, id, ip, refreshExpiresAt)
	if e != nil {
		return 0, domain.Token{}, domain.Token{}, e
	}

	accessToken := domain.Token{
		Token:     accessTokenHash,
		ExpiresAt: accessExpiresAt.UnixNano() / 1e+6,
	}
	refreshToken := domain.Token{
		Token:     refreshTokenHash,
		ExpiresAt: refreshExpiresAt.UnixNano() / 1e+6,
	}

	return *number, accessToken, refreshToken, nil
}

func (s *service) generateToken(
	ctx context.Context,
	tx models.Transaction,
	purpose domain.AuthPurpose,
	number int64,
	userId uuid.UUID,
	ip string,
	expire time.Time,
) (string, *errors.Error) {
	jti, err := uuid.NewV4()
	if err != nil {
		return "", s.log.ServiceError(errors.WD(errors.ErrInternal, err))
	}

	claims := jwt.MapClaims{
		"jti":     jti.String(),
		"user_id": userId.String(),
		"ip":      ip,
		"exp":     expire.Unix(),
		"number":  number,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	res, err := token.SignedString([]byte(s.key))
	if err != nil {
		return "", s.log.ServiceError(errors.WD(errors.ErrAuthParseToken, err))
	}

	if purpose == domain.PurposeRefresh {
		if err := s.jwtRepo.DropToken(ctx, tx, userId.String(), number); err != nil {
			return "", errors.DatabaseError(err)
		}

		tokenToSave := domain.TokenPayload{
			UserId:    userId,
			Number:    number,
			Payload:   res,
			ExpiresAt: expire.Unix(),
		}

		if _, err := s.jwtRepo.AddToken(ctx, tx, converter.TokenModelFromDomain(tokenToSave)); err != nil {
			return "", errors.DatabaseError(err)
		}

		res = base64.StdEncoding.EncodeToString([]byte(res))
	}

	return res, nil
}

func (s *service) parseToken(
	token string,
	purpose domain.AuthPurpose,
) (*jwt.Token, *errors.Error) {
	if purpose == domain.PurposeRefresh {
		t, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			return nil, s.log.ServiceError(errors.ErrAuthParseToken)
		}

		token = string(t)
	}

	res, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.key), nil
	})

	if err != nil && res == nil {
		return nil, s.log.ServiceError(errors.ErrAuthParseToken)
	}

	return res, nil
}

func (s *service) checkToken(
	ctx context.Context, tx models.Transaction, res *jwt.Token,
) (domain.Account, int64, *errors.Error) {
	claims, ok := res.Claims.(jwt.MapClaims)
	if !ok {
		return domain.Account{}, 0, errors.ErrAuthParseToken
	}

	if !claims.VerifyExpiresAt(s.timeAdapter.Now().Unix(), true) {
		return domain.Account{}, 0, errors.ErrAuthExpiredToken
	}

	if err := claims.Valid(); err != nil {
		return domain.Account{}, 0, errors.ErrAuthInvalidToken
	}

	user_id, err := s.parseTokenStringClaim(claims, "user_id")
	if err != nil {
		return domain.Account{}, 0, err
	}
	number, err := s.parseTokenIntClaim(claims, "number")
	if err != nil {
		return domain.Account{}, 0, err
	}
	ip, err := s.parseTokenStringClaim(claims, "ip")
	if err != nil {
		return domain.Account{}, 0, err
	}

	if _, err := s.jwtRepo.CheckToken(ctx, tx, number, user_id); err != nil {
		if err == errors.ErrPostgresNotExists {
			return domain.Account{}, 0, errors.ErrAuthFailed
		}
		return domain.Account{}, 0, errors.DatabaseError(err)
	}

	return domain.Account{
		UserId: user_id,
		Ip:     ip,
	}, number, nil
}

func (s *service) parseTokenIntClaim(claims jwt.MapClaims, key string) (int64, *errors.Error) {
	if parsedValue, ok := claims[key].(float64); !ok {
		return 0, errors.ErrAuthInvalidToken
	} else {
		return int64(parsedValue), nil
	}
}

func (s *service) parseTokenStringClaim(claims jwt.MapClaims, key string) (string, *errors.Error) {
	if stringValue, ok := claims[key].(string); !ok {
		return "", errors.ErrAuthInvalidToken
	} else {
		return stringValue, nil
	}
}
