package domain

type AuthPurpose int32 // enum для типа jwt-токена
const (
	PurposeAccess  = AuthPurpose(iota) // access токен
	PurposeRefresh                     // refresh токен
)
