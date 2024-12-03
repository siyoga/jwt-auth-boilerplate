package log

import (
	"fmt"

	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"

	"go.uber.org/zap"
)

type (
	Logger interface {
		Named(name string) Logger

		Zap() *zap.Logger

		Error(err, reason error) error
		SqlError(err, reason error, query string) error

		ServiceError(err *errors.Error) *errors.Error
		ServiceTxError(err error) *errors.Error
		ServiceDatabaseError(err error) *errors.Error

		Info(msg string, fields ...zap.Field)
	}

	logger struct {
		moduleName string
		log        *zap.Logger
	}
)

func NewLogger(l *zap.Logger) Logger {
	return &logger{
		moduleName: "github.com/siyoga/jwt-auth-boilerplate",
		log:        l,
	}
}

func (l *logger) Named(name string) Logger {
	return &logger{
		log: l.log.Named(name),
	}
}

func (l *logger) Zap() *zap.Logger {
	return l.log
}

func (l *logger) Panic(text string, fields ...zap.Field) {
	l.log.Panic(text, fields...)
}
func (l *logger) Info(text string, fields ...zap.Field) {
	l.log.Info(text, fields...)
}
func (l *logger) Warn(text string, fields ...zap.Field) {
	l.log.Warn(text, fields...)
}
func (l *logger) Error(err, reason error) error {
	if reason == nil {
		reason = err
	}

	l.log.Error(l.genError(err, reason.Error()))
	return err
}
func (l *logger) SqlError(err, reason error, query string) error {
	if reason == nil {
		reason = err
	}

	l.log.Error(l.genError(fmt.Errorf("%s \n %s", err.Error(), query), reason.Error()))
	return err
}

func (l *logger) ServiceDatabaseError(err error) *errors.Error {
	return l.ServiceError(errors.DatabaseError(err))
}

func (l *logger) ServiceError(err *errors.Error) *errors.Error {
	e := fmt.Errorf(err.Reason)
	if err.Details != nil {
		e = err.Details
	}

	l.log.Error(l.genError(e, err.Reason))
	return err
}

func (l *logger) ServiceTxError(err error) *errors.Error {
	l.log.Error(l.genError(err, errors.ErrPostgresTx.Error()))
	return errors.DatabaseError(err)
}
