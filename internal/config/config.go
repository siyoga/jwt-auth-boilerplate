package config

import (
	"fmt"
	"time"

	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Base     Base
		Postgres Postgres
		Timeouts Timeouts
		Auth     Auth
	}

	Base struct {
		Mode       domain.Mode
		Locale     int64
		ServerPort int
	}

	Timeouts struct {
		RequestTimeout time.Duration
		AuthTimeout    time.Duration
	}

	Postgres struct {
		DSN     string
		CertLoc string
	}

	Auth struct {
		JwtKey              string
		CookieTimeout       time.Duration
		AccessTokenTimeout  time.Duration
		RefreshTokenTimeout time.Duration
	}
)

func NewConfig(path string) (*Config, error) {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("config")
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	mode := domain.Mode(v.GetString("mode"))
	fmt.Println(mode)
	jwt, err := loadJwtKey(v, mode)
	if err != nil {
		fmt.Println("jwt")
		return nil, err
	}

	psql, err := loadPostgresCreds(v, mode)
	if err != nil {
		fmt.Println("psql")
		return nil, err
	}

	return &Config{
		Postgres: psql,
		Base: Base{
			ServerPort: v.GetInt("server.port"),
			Locale:     v.GetInt64("server.locale"),
			Mode:       mode,
		},
		Timeouts: Timeouts{
			RequestTimeout: v.GetDuration("timeouts.request"),
			AuthTimeout:    v.GetDuration("timeouts.auth"),
		},
		Auth: Auth{
			JwtKey:              jwt,
			AccessTokenTimeout:  getTimeoutFromInt(v, "auth.at"),
			RefreshTokenTimeout: getTimeoutFromInt(v, "auth.rt"),
			CookieTimeout:       getTimeoutFromInt(v, "auth.cookie"),
		},
	}, nil
}
