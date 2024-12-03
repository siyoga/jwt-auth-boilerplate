package dependencies

import (
	"github.com/siyoga/jwt-auth-boilerplate/internal/handler"
)

func (d *dependencies) Middlewares() handler.Middleware {
	if d.middlewares == nil {
		d.middlewares = handler.NewMiddleware(
			d.log,
			d.cfg.Timeouts,
			d.AuthService(),
		)
	}

	return d.middlewares
}

func (d *dependencies) RequestHandler() handler.RequestHandler {
	if d.reqHandler == nil {
		d.reqHandler = handler.NewRequestHandler(
			d.log,
			d.cfg.Auth.CookieTimeout,
		)
	}

	return d.reqHandler
}

func (d *dependencies) AuthHandler() handler.Handler {
	if d.authHandler == nil {
		d.authHandler = handler.NewAuthHandler(
			d.cfg.Base,
			d.cfg.Timeouts,
			d.AuthService(),
			d.TimeAdapter(),
			d.RequestHandler(),
			d.Middlewares(),
		)
	}

	return d.authHandler
}
