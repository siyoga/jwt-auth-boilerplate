package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/siyoga/jwt-auth-boilerplate/internal/adapter"
	"github.com/siyoga/jwt-auth-boilerplate/internal/config"
	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"
	"github.com/siyoga/jwt-auth-boilerplate/internal/service"

	"github.com/gorilla/mux"
)

type (
	authHandler struct {
		cfg      *config.Base
		timeouts *config.Timeouts

		authService service.AuthService

		timeAdapter adapter.TimeAdapter

		reqHandler RequestHandler
		middleware Middleware
	}
)

func NewAuthHandler(
	cfg config.Base,
	timeouts config.Timeouts,

	authSvc service.AuthService,

	timeAdpt adapter.TimeAdapter,

	reqHandler RequestHandler,
	middlewares Middleware,
) Handler {
	return &authHandler{
		cfg:      &cfg,
		timeouts: &timeouts,

		authService: authSvc,

		timeAdapter: timeAdpt,

		reqHandler: reqHandler,
		middleware: middlewares,
	}
}

func (h *authHandler) FillHandlers(router *mux.Router) {
	base := "/auth"
	r := router.PathPrefix(base).Subrouter()
	h.reqHandler.HandleJsonRequest(r, base, "/create", http.MethodPost, h.create)
	h.reqHandler.HandleJsonRequestWithMiddleware(r, base, "/whoami", http.MethodGet, h.whoAmI, h.middleware.JwtAuthMiddleware())
	h.reqHandler.HandleJsonRequestWithMiddleware(r, base, "/refresh", http.MethodGet, h.refresh, h.middleware.IpMiddleware())
	h.reqHandler.HandleJsonRequestWithMiddleware(r, base, "/", http.MethodPost, h.auth, h.middleware.IpMiddleware())
}

func (h *authHandler) auth(ctx context.Context, acc *domain.Account, r *http.Request) jsonResponse {
	ctx, cancel := context.WithTimeout(ctx, h.timeouts.RequestTimeout)
	defer cancel()

	var req AuthData
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errorResponse(errors.WD(errors.ErrInternal, err))
	}
	ip := ctx.Value(IpCtxKey).(string)

	at, rt, err := h.authService.Auth(ctx, ip, nil, req.Email, req.Password)
	if err != nil {
		return errorResponse(err)
	}

	tokens := DomainTokensToResponseTokens(at, rt)
	return successResponse(
		tokens,
		http.StatusOK,
		nil,
	)
}

func (h *authHandler) create(ctx context.Context, acc *domain.Account, r *http.Request) jsonResponse {
	ctx, cancel := context.WithTimeout(ctx, h.timeouts.RequestTimeout)
	defer cancel()

	var req CreateData
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errorResponse(errors.WD(errors.ErrInternal, err))
	}

	if e := h.authService.Create(
		ctx,
		domain.User{
			Username: req.Username,
			Pass:     req.Password,
			Email:    req.Email,
		}); e != nil {
		return errorResponse(e)
	}

	return successResponse(
		nil,
		http.StatusCreated,
		nil,
	)
}

func (h *authHandler) refresh(ctx context.Context, acc *domain.Account, r *http.Request) jsonResponse {
	ctx, cancel := context.WithTimeout(ctx, h.timeouts.RequestTimeout)
	defer cancel()

	isIpDifferent := false
	a := r.Header.Get(AuthHeader)
	isToken := strings.HasPrefix(a, TokenStart)
	if !isToken {
		return errorResponse(errors.ErrAuthInvalidToken)
	}
	ip := ctx.Value(IpCtxKey).(string)

	account, num, e := h.authService.Verify(ctx, a[TokenStartInd:], domain.PurposeRefresh)
	if e != nil {
		return errorResponse(e)
	}

	if account.Ip != ip {
		// здесь вызывать сервис связанный с почтой, который отправит уведомление
		isIpDifferent = true
	}

	at, rt, e := h.authService.CreateTokens(ctx, ip, account.UserId, num)
	if e != nil {
		return errorResponse(e)
	}

	return successResponse(
		Refresh{
			Tokens:        DomainTokensToResponseTokens(at, rt),
			IsIpDifferent: isIpDifferent,
		},
		http.StatusOK,
		nil,
	)
}

func (h *authHandler) whoAmI(ctx context.Context, acc *domain.Account, r *http.Request) jsonResponse {
	ctx, cancel := context.WithTimeout(ctx, h.timeouts.RequestTimeout)
	defer cancel()

	if acc != nil {
		user, e := h.authService.GetById(ctx, acc.UserId)
		if e != nil {
			return errorResponse(e)
		}

		usr := DomainUserToResponseUser(user, acc.Ip)
		accCookie, err := createCookie("acc", usr)
		if err != nil {
			return errorResponse(err)
		}

		return successResponse(
			DomainUserToResponseUser(user, acc.Ip),
			http.StatusOK,
			[]http.Cookie{accCookie},
		)
	}

	return successResponse(
		map[string]string{"message": "not logged in"},
		http.StatusOK,
		nil,
	)
}
