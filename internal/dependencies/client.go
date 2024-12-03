package dependencies

import (
	"github.com/siyoga/jwt-auth-boilerplate/internal/database"

	"go.uber.org/zap"
)

func (d *dependencies) PostgresClient() *database.PostgresClient {
	if d.psqlClient == nil {
		var err error
		if d.psqlClient, err = database.NewPostgresClient(d.cfg.Postgres.DSN, d.cfg.Postgres.CertLoc); err != nil {
			d.log.Zap().Panic("initialize postgres client", zap.Error(err))
		}

		d.shutdownCallbacks = append(d.shutdownCallbacks, func() {
			if err := d.psqlClient.Close(); err != nil {
				d.log.Zap().Warn("initialize postgres shutdown callback", zap.Error(err))
				return
			}
		})
	}

	return d.psqlClient
}
