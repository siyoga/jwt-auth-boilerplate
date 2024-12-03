package models

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type (
	Tx struct {
		*sqlx.Tx
		ctx context.Context
	}

	Transaction interface {
		Commit() error
		Rollback()
		Txm() *sqlx.Tx
	}
)

func (t *Tx) Txm() *sqlx.Tx {
	return t.Tx
}

func (t *Tx) Rollback() {
	_ = t.Tx.Rollback()
}
