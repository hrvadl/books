package cfg

import (
	"errors"

	"github.com/caarlos0/env/v11"
)

var ErrFailedToParse = errors.New("failed to parse config")

func NewFromEnv() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.Join(ErrFailedToParse, err)
	}

	return &cfg, nil
}

type Config struct {
	Host string `env:"HOST,required,notEmpty"`
	Port string `env:"PORT,required,notEmpty"`
}
