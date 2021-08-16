package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	PSQL       *PSQL
	Whitelists *WhitelistProcessor
}

func Init() (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Init(cfg); err != nil {
		return nil, fmt.Errorf("config init error: %w", err)
	}

	return cfg, nil
}
