package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Credentials `envPrefix:"CREDENTIALS_"`
}

type Credentials struct {
	Path      string `env:"PATH"       envDefault:"credentials.json"`
	TokenPath string `env:"TOKEN_PATH" envDefault:"token.json"`
	Data      []byte
}

func Load() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse env: %w", err)
	}

	if err := cfg.Credentials.ReadData(); err != nil {
		return nil, fmt.Errorf("failed to read credentials: %w", err)
	}

	return cfg, nil
}

func (c *Credentials) ReadData() error {
	data, err := os.ReadFile(c.Path)
	if err != nil {
		return fmt.Errorf("failed to read credentials file: %w", err)
	}

	c.Data = data

	return nil
}
