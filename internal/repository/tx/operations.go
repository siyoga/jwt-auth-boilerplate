package tx

import (
	"context"

	"github.com/siyoga/jwt-auth-boilerplate/internal/database"
	def "github.com/siyoga/jwt-auth-boilerplate/internal/repository"
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository/models"
)

var _ def.TxRepository = (*repository)(nil)

type (
	repository struct {
		client *database.PostgresClient
	}
)

func NewRepository(client *database.PostgresClient) *repository {
	return &repository{
		client: client,
	}
}

func (r repository) StartTransaction(ctx context.Context) (*models.Tx, error) {
	tx, err := r.client.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &models.Tx{
		Tx: tx,
	}, nil
}
