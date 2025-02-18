package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Hostname     string        `env:"AVIAPI_HOSTNAME" env-default:"localhost"`
	Port         string        `env:"AVIAPI_PORT" env-default:"8000"`
	DatabaseDSN  string        `env:"AVIAPI_DATABASE_DSN" env-default:"postgres://postgres:postgres@localhost:5432/finapi?sslmode=disable"`
	ReadTimeout  time.Duration `env:"AVIAPI__READ_TIMEOUT" env-default:"5s"`
	WriteTimeout time.Duration `env:"AVIAPI_WRITE_TIMEPUT" env-default:"5s"`
	IdleTimeout  time.Duration `env:"AVIAPI_IDLE_TIMEOUT" env-default:"60s"`
}

func Load() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read env for config: %w", err)
	}

	return &cfg, nil
}
