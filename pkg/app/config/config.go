package config

import (
	"github.com/eser/go-service/pkg/bliss"
)

type AppConfig struct {
	bliss.BaseConfig

	AppName   string `conf:"NAME" default:"karma-points"`
	JWTSecret string `conf:"JWT"`
	Postgres  struct {
		Dsn string `conf:"DSN" default:"postgres://localhost:5432"`
	} `conf:"POSTGRES"`
}
