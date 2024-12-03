package domain

import "github.com/gofrs/uuid"

type (
	Token struct {
		Token     string
		ExpiresAt int64
	}

	TokenPayload struct {
		UserId    uuid.UUID
		Number    int64
		Payload   string
		ExpiresAt int64
	}
)
