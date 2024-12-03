package dependencies

import (
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository"
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository/jwt"
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository/tx"
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository/user"
)

func (d *dependencies) TransactionRepo() repository.TxRepository {
	if d.txRepo == nil {
		d.txRepo = tx.NewRepository(d.PostgresClient())
	}

	return d.txRepo
}

func (d *dependencies) UserRepo() repository.UserRepository {
	if d.userRepo == nil {
		d.userRepo = user.NewRepository(d.log, d.PostgresClient())
	}

	return d.userRepo
}

func (d *dependencies) JwtRepo() repository.JwtRepository {
	if d.jwtRepo == nil {
		d.jwtRepo = jwt.NewRepository(d.log, d.PostgresClient())
	}

	return d.jwtRepo
}
