package handler

import "github.com/siyoga/jwt-auth-boilerplate/internal/domain"

func DomainTokensToResponseTokens(at, rt domain.Token) Tokens {
	return Tokens{
		AccessToken: Token{
			Token:     at.Token,
			ExpiresAt: at.ExpiresAt,
		},
		RefreshToken: Token{
			Token:     rt.Token,
			ExpiresAt: rt.ExpiresAt,
		},
	}
}

func DomainUserToResponseUser(user domain.User, ip string) User {
	return User{
		Id:       user.Id.String(),
		Username: user.Username,
		Ip:       ip,
		Email:    user.Email,
	}
}
