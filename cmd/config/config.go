package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func LoadEnv(envPath ...string) error {
	if len(envPath) > 0 {
		if err := godotenv.Load(envPath...); err != nil {
			return fmt.Errorf("Error loading .env file")
		}
	}

	config := &config{}
	if err := env.Parse(&config.Server); err != nil {
		return fmt.Errorf("env load error: %v", err)
	}

	if err := env.Parse(&config.Database); err != nil {
		return fmt.Errorf("env load error: %v", err)
	}

	if err := env.Parse(&config.Buckets); err != nil {
		return fmt.Errorf("env load error: %v", err)
	}

	if err := env.Parse(&config.NewRelic); err != nil {
		return fmt.Errorf("env load error: %v", err)
	}

	Config = config
	return nil
}
