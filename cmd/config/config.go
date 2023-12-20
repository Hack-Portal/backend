package config

import (
	"log"

	"github.com/caarlos0/env"
)

func LoadEnv() {
	config := &config{}
	if err := env.Parse(&config.Server); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	if err := env.Parse(&config.Database); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	if err := env.Parse(&config.Buckets); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	Config = config
}
