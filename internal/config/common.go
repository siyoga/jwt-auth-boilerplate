package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/siyoga/jwt-auth-boilerplate/internal/domain"
	"github.com/spf13/viper"
)

func loadJwtKey(v *viper.Viper, mode domain.Mode) (string, error) {
	if mode == domain.Local {
		key := v.GetString("jwt_key")
		if key == "" {
			return "", fmt.Errorf("provide jwt sign key with in jwt_key property")
		}

		return key, nil
	}

	type jwtCreds struct {
		Key string `json:"key"`
	}

	path := v.GetString("apis.jwt")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	var creds jwtCreds
	if err := json.Unmarshal(data, &creds); err != nil {
		return "", err
	}

	return creds.Key, nil
}

func loadPostgresCreds(v *viper.Viper, mode domain.Mode) (Postgres, error) {
	if mode == domain.Local {
		psqlSource := v.GetString("psql")
		if psqlSource == "" {
			return Postgres{}, fmt.Errorf("provide psql dsn by psql property")
		}

		return Postgres{
			DSN:     psqlSource,
			CertLoc: "",
		}, nil
	}

	type psqlCreds struct {
		Source  string `json:"source"`
		CertLoc string `json:"cert_loc"`
	}

	path := v.GetString("apis.psql")
	data, err := os.ReadFile(path)
	if err != nil {
		return Postgres{}, err
	}

	var creds psqlCreds
	if err := json.Unmarshal(data, &creds); err != nil {
		return Postgres{}, nil
	}

	return Postgres{
		DSN:     creds.Source,
		CertLoc: creds.CertLoc,
	}, nil
}

func getTimeoutFromInt(v *viper.Viper, field string) time.Duration {
	return time.Second * time.Duration(v.GetInt64(field))
}
