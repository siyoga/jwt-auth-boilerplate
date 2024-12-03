package dependencies

import (
	"github.com/siyoga/jwt-auth-boilerplate/internal/service"
	"github.com/siyoga/jwt-auth-boilerplate/internal/service/auth"
)

func (d *dependencies) AuthService() service.AuthService {
	if d.authService == nil {
		d.authService = auth.NewService(
			d.log,
			d.cfg.Auth,
			d.TransactionRepo(),
			d.UserRepo(),
			d.JwtRepo(),
			d.TimeAdapter(),
			d.RandomAdapter(),
		)
	}

	return d.authService
}
