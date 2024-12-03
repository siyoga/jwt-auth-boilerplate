package auth

import (
	"time"

	"github.com/siyoga/jwt-auth-boilerplate/internal/adapter"
	"github.com/siyoga/jwt-auth-boilerplate/internal/config"
	"github.com/siyoga/jwt-auth-boilerplate/internal/log"
	r "github.com/siyoga/jwt-auth-boilerplate/internal/repository"
	def "github.com/siyoga/jwt-auth-boilerplate/internal/service"
)

var _ def.AuthService = (*service)(nil)

type (
	service struct {
		cfg config.Auth
		log log.Logger

		txRepo   r.TxRepository
		userRepo r.UserRepository
		jwtRepo  r.JwtRepository

		timeAdapter   adapter.TimeAdapter
		randomAdapter adapter.RandomAdapter

		key                                     string
		accessTokenTimeout, refreshTokenTimeout time.Duration
	}
)

func NewService(
	log log.Logger,
	cfg config.Auth,

	txRepo r.TxRepository,
	userRepo r.UserRepository,
	jwtRepo r.JwtRepository,

	timeAdpt adapter.TimeAdapter,
	randAdpt adapter.RandomAdapter,
) *service {
	return &service{
		log: log,
		cfg: cfg,

		txRepo:   txRepo,
		userRepo: userRepo,
		jwtRepo:  jwtRepo,

		timeAdapter:   timeAdpt,
		randomAdapter: randAdpt,

		key:                 cfg.JwtKey,
		accessTokenTimeout:  cfg.AccessTokenTimeout,
		refreshTokenTimeout: cfg.RefreshTokenTimeout,
	}
}
