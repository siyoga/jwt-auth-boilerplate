package models

import (
	"github.com/gofrs/uuid"
)

type Token struct {
	UserId    uuid.UUID `db:"user_id"`
	Number    int64     `db:"number"`
	Payload   string    `db:"payload"`
	ExpiresAt int64     `db:"expires_at"`
}
