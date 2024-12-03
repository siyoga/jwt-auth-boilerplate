package errors

import "errors"

var (
	// Postgres
	ErrPostgresScanRaw            = errors.New("postgresql scan error")
	ErrPostgresQueryRaw           = errors.New("postgresql query error")
	ErrPostgresGetRaw             = errors.New("postgresql get error")
	ErrPostgresQueryRowRaw        = errors.New("postgresql query row error")
	ErrPostgresExecRaw            = errors.New("postgresql exec error")
	ErrPostgresRowsAffectedRaw    = errors.New("postgresql rows affected error")
	ErrPostgresLastInsertIdRaw    = errors.New("postgresql last insert id error")
	ErrPostgresRebindRaw          = errors.New("postgresql rebind error")
	ErrPostgresRowsCloseRaw       = errors.New("postgresql rows close error")
	ErrPostgresNoRowsWereAffected = errors.New("no rows were affected")
	ErrPostgresTx                 = errors.New("postgresql transaction error")
	ErrPostgresNotExists          = errors.New("struct not exists")

	// Auth
	ErrAuthTokenNotExistRaw          = errors.New("token does not exists")
	ErrAuthNumberAssignmentFailedRaw = errors.New("number assignment failed")
	ErrAuthParseTokenRaw             = errors.New("parse token failed")
	ErrAuthHashTokenRaw              = errors.New("hashing token error")
	ErrAuthHashPasswordRaw           = errors.New("hashing password error")

	ErrInvalidUserIdRaw = errors.New("invalid user id ")
)

var (
	// Auth
	ErrAuthInvalidTokenPurpose = &Error{Code: 400, Reason: "invalid token purpose"}
	ErrAuthHashPassword        = &Error{Code: 400, Reason: ErrAuthHashPasswordRaw.Error()}
	ErrAuthHashToken           = &Error{Code: 400, Reason: ErrAuthHashTokenRaw.Error()}
	ErrAuthExpiredToken        = &Error{Code: 400, Reason: "expired token"}
	ErrAuthInvalidToken        = &Error{Code: 400, Reason: "invalid token"}
	ErrAuthParseToken          = &Error{Code: 400, Reason: ErrAuthParseTokenRaw.Error()}
	ErrAuthCreateAccessToken   = &Error{Code: 400, Reason: "create access token error"}
	ErrAuthFailed              = &Error{Code: 401, Reason: "auth failed"}

	// General
	ErrPermissionDenied   = &Error{Code: 403, Reason: "permission denied"}
	ErrParse              = &Error{Code: 400, Reason: "parse failed"}
	ErrInternal           = &Error{Code: 500, Reason: "internal error"}
	ErrConflict           = &Error{Code: 409, Reason: "conflict"}
	ErrInvalidCredentials = &Error{Code: 401, Reason: "invalid credentials"}
	ErrValidation         = &Error{Code: 400, Reason: "validation failed"}
	ErrTimeout            = &Error{Code: 504, Reason: "timeout"}
)
