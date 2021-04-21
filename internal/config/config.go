package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"

	"rebrainme/gotest/internal/processors/whitelists"
	"rebrainme/gotest/internal/system/database/psql"
)

type Config struct {
	PSQL       *psql.Config
	Whitelists *whitelists.Config
}

func Init() (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Init(cfg); err != nil {
		return nil, fmt.Errorf("config init error: %w", err)
	}

	return cfg, nil
}
