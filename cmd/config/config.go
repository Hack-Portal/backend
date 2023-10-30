package config

import (
	"temp/pkg/logger"

	"github.com/caarlos0/env"
)

func LoadEnv(l logger.Logger) {
	config := &config{}
	if err := env.Parse(&config.Server); err != nil {
		l.Panic(err)
	}

	if err := env.Parse(&config.Cockroach); err != nil {
		l.Panic(err)
	}

	if err := env.Parse(&config.NewRelic); err != nil {
		l.Panic(err)
	}

	if err := env.Parse(&config.Firebase); err != nil {
		l.Panic(err)
	}
	Config = config
}
