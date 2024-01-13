package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func LoadEnv(envPath ...string) {

	if envPath != nil {
		if err := godotenv.Load(envPath...); err != nil {
			log.Fatalf("Error loading .env file")
		}
		log.Println("load .env file")
	}

	config := &config{}
	if err := env.Parse(&config.Server); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	if err := env.Parse(&config.Database); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	if err := env.Parse(&config.Redis); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	if err := env.Parse(&config.Buckets); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	if err := env.Parse(&config.NewRelic); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	if err := env.Parse(&config.Discord); err != nil {
		log.Fatalf("env load error: %v", err)
	}

	Config = config
}
