package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/siyoga/jwt-auth-boilerplate/internal/errors"
	"github.com/siyoga/jwt-auth-boilerplate/internal/log"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type (
	jsonHandler func(ctx context.Context, acc *domain.Account, r *http.Request) jsonResponse

	requestHandler struct {
		log           log.Logger
		cookieTimeout time.Duration
	}

	jsonResponse struct {
		Cookies []http.Cookie
		Data    interface{}   `json:"data;omitempty"`
		Code    int           `json:"code"`
		Error   *errors.Error `json:"error;omitempty"`
	}

	Handler interface {
		FillHandlers(router *mux.Router)
	}

	RequestHandler interface {
		HandleJsonRequest(router *mux.Router, main, path, method string, handler jsonHandler)
		HandleJsonRequestWithMiddleware(router *mux.Router, main, path, method string, handler jsonHandler, middleware func(http.Handler) http.Handler)
	}
)

func NewRequestHandler(
	log log.Logger,
	cookieTimeout time.Duration,
) RequestHandler {
	return &requestHandler{
		log:           log,
		cookieTimeout: cookieTimeout,
	}
}

func (h *requestHandler) HandleJsonRequestWithMiddleware(
	router *mux.Router,
	main, path, method string,
	handler jsonHandler,
	middleware func(http.Handler) http.Handler,
) {
	router.Handle(path, middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.requestWithJsonResult(main, path, method, w, r, handler)
	}))).Methods(method)
}

func (h *requestHandler) HandleJsonRequest(router *mux.Router, main, path, method string, handler jsonHandler) {
	router.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.requestWithJsonResult(main, path, method, w, r, handler)
	})).Methods(method)
}

func (h *requestHandler) requestWithJsonResult(
	main, path, method string,
	w http.ResponseWriter,
	r *http.Request,
	handler jsonHandler,
) {
	defer r.Body.Close()

	accPayload := r.Context().Value(AccountCtxKey)
	var acc *domain.Account

	if accPayload != nil {
		acc = accPayload.(*domain.Account)
	} else {
		acc = nil
	}

	res := handler(r.Context(), acc, r)
	var resBytes []byte
	if res.Error != nil {
		h.logRequest(main, path, method, r.Header.Get(IpHeader), false, res.Error, acc)

		sendJSON(w, res.Code, res.Error)
		return
	} else {
		resBytes, _ = json.Marshal(res.Data)
		h.logRequest(main, path, method, r.Header.Get(IpHeader), true, nil, acc)

		if res.Cookies != nil {
			for _, cookie := range res.Cookies {
				cookie.Expires = time.Now().Add(h.cookieTimeout)
				http.SetCookie(w, &cookie)
			}
		}

		sendBytes(w, res.Code, resBytes)
		return
	}
}

func (h *requestHandler) logRequest(main, path, method, ip string, success bool, err *errors.Error, acc *domain.Account) {
	fields := []zap.Field{
		zap.String("method", method), zap.String("path", main+path), zap.String("ip", ip),
	}

	if acc != nil {
		fields = append(fields, zap.String("acc", fmt.Sprintf("%+v", acc)))
	}

	if success {
		h.log.Info("request", fields...)
	} else {
		fields = append(fields, zap.String("error", fmt.Sprintf("details=%+v, reason=%+v, status=%+v", err.Details, err.Reason, err.Code)))
		h.log.Zap().Warn("request", fields...)
	}
}

func errorResponse(err *errors.Error) jsonResponse {
	return jsonResponse{
		Code:  int(err.Code),
		Error: err,
	}
}

func successResponse(data interface{}, status int, cookie []http.Cookie) jsonResponse {
	return jsonResponse{
		Data:    data,
		Code:    status,
		Cookies: cookie,
	}
}

func createCookie(name string, value interface{}) (http.Cookie, *errors.Error) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return http.Cookie{}, errors.WD(errors.ErrInternal, err)
	}

	return http.Cookie{
		Name:     name,
		Value:    string(jsonData),
		Path:     "/",
		MaxAge:   0,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}, nil
}
