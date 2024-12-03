package converter

import (
	"time"

	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository/models"

	"github.com/gofrs/uuid"
)

func UserModelFromDomain(u domain.User) models.User {
	return models.User{
		Id:        u.Id.String(),
		Username:  u.Username,
		Password:  u.Pass,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.UnixMilli(),
	}
}

func UserDomainFromModel(u models.User) domain.User {
	id, _ := uuid.FromString(u.Id)
	return domain.User{
		Id:        id,
		Username:  u.Username,
		Pass:      u.Password,
		Email:     u.Email,
		CreatedAt: time.UnixMilli(u.CreatedAt),
	}
}
