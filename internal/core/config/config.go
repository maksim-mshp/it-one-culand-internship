package config

import (
	"fmt"
	"go-simpler.org/env"
)

type Database struct {
	Host     string `env:"POSTGRES_HOST,required"`
	Port     int    `env:"POSTGRES_PORT,required"`
	User     string `env:"POSTGRES_USER,required"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	Database string `env:"POSTGRES_DATABASE,required"`
}

type Config struct {
	Port          int `env:"PORT,required"`
	Database      Database
	JWTToken      string `env:"JWT_TOKEN,required"`
	InternalToken string `env:"INTERNAL_TOKEN,required"`
}

func Load() (*Config, error) {
	cfg := Config{}
	if err := env.Load(&cfg, nil); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return &cfg, nil
}
