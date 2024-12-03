package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/siyoga/jwt-auth-boilerplate/internal/config"
	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/siyoga/jwt-auth-boilerplate/internal/log"
	"github.com/siyoga/jwt-auth-boilerplate/internal/service"
)

type (
	Middleware interface {
		JwtAuthMiddleware() func(http.Handler) http.Handler
		IpMiddleware() func(http.Handler) http.Handler
	}

	middleware struct {
		log log.Logger

		timeouts    config.Timeouts
		authService service.AuthService
	}
)

func NewMiddleware(
	log log.Logger,
	timeouts config.Timeouts,
	authService service.AuthService,
) Middleware {
	return &middleware{
		log:         log,
		timeouts:    timeouts,
		authService: authService,
	}
}

func (m *middleware) JwtAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			a := r.Header.Get(AuthHeader)
			isToken := strings.HasPrefix(a, TokenStart)
			if !isToken {
				next.ServeHTTP(w, r)
				return
			}

			ctx, cancel := context.WithTimeout(r.Context(), m.timeouts.AuthTimeout)
			defer cancel()

			acc, num, err := m.authService.Verify(
				ctx, a[TokenStartInd:], domain.PurposeAccess,
			)
			if err != nil {
				details := ""
				if err.Details != nil {
					details = err.Details.Error()
				}
				m.log.Zap().Error(fmt.Sprintf("auth_failed_log_jwt err=%s, token=%s", fmt.Sprintf("(reason=%s, details=%s)", err.Reason, details), a[TokenStartInd:]))
				sendJSON(w, 200, err)
				return
			}
			ctx = context.WithValue(r.Context(), AccountCtxKey, &acc)
			ctx = context.WithValue(ctx, TokenNumberCtxKey, num)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (m *middleware) IpMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), IpCtxKey, strings.Split(r.RemoteAddr, ":")[0])
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
