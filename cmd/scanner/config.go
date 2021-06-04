package main

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	AllowedOrigins []string `envconfig:"CORS_ALLOWED_ORIGINS" default:"http://localhost:4200"`
	DBFile         string   `envconfig:"DB_FILE" default:"./scans.db"`
	Port           int      `envconfig:"PORT" default:"80"`
}

func loadConfigFromEnv() (*config, error) {
	var cfg config
	err := envconfig.Process("", &cfg)
	return &cfg, err
}
