package config

import (
	"fmt"
	"go-simpler.org/env"
)

type Database struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Database string `env:"DB_DATABASE,required"`
}

type Config struct {
	Port     int `env:"PORT,required"`
	Database Database
}

func Load() (*Config, error) {
	cfg := Config{}
	if err := env.Load(&cfg, nil); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return &cfg, nil
}
