package converter

import (
	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository/models"
)

func TokenModelFromDomain(t domain.TokenPayload) models.Token {
	return models.Token{
		UserId:    t.UserId,
		Number:    t.Number,
		Payload:   t.Payload,
		ExpiresAt: t.ExpiresAt,
	}
}
