package config

import (
	"log"

	"github.com/caarlos0/env"
)

func init() {
	config := &config{}
	if err := env.Parse(&config.Server); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	if err := env.Parse(&config.Database); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	if err := env.Parse(&config.NewRelic); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	if err := env.Parse(&config.Firebase); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	Config = config
}
