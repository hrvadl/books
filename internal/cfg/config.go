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
	Host         string `env:"HOST,required,notEmpty"`
	Port         string `env:"PORT,required,notEmpty"`
	PostgresDSN  string `env:"POSTGRES_DSN,required,notEmpty"`
	GCPKeypath   string `env:"GCP_SERVICE_ACCOUNT_KEY,required,notEmpty"`
	GCPProjectID string `env:"GCP_PROJECT_ID,required,notEmpty"`
	FirestoreDB  string `env:"GCP_FIRESTORE_DB,required,notEmpty"`
}
