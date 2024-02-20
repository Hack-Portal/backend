package config

import (
	"github.com/caarlos0/env"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
)

func LoadEnv(l logger.Logger) {
	config := &config{}
	if err := env.Parse(&config.Server); err != nil {
		l.Panic(err)
	}

	if err := env.Parse(&config.Postgres); err != nil {
		l.Panic(err)
	}

	if err := env.Parse(&config.NewRelic); err != nil {
		l.Panic(err)
	}

	if err := env.Parse(&config.Firebase); err != nil {
		l.Panic(err)
	}
}
