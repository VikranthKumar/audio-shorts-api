package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT" default:"8080"`
		Host string `envconfig:"SERVER_HOST"`
	}
	Postgres struct {
		Username string `envconfig:"POSTGRES_USER"`
		Password string `envconfig:"POSTGRES_PASSWORD"`
		Database string `envconfig:"POSTGRES_DB"`
		Port     string `envconfig:"POSTGRES_PORT" default:"5432"`
		Host     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	}
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := readEnv(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func readEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}
