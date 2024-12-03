package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

type Role int32 // enum для роли
const (
	RoleUser Role = iota
	RoleAdmin
)

type (
	User struct {
		Id        uuid.UUID
		Username  string
		Pass      string
		Email     string
		CreatedAt time.Time
	}

	Account struct {
		UserId string
		Ip     string
	}

	AccountInfo struct {
		Account
		Firstname string
		Lastname  string
		Email     string
	}
)

func (a User) IsEmpty() bool {
	return a.Email == "" && a.Username == ""
}
