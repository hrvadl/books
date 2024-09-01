package db

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const driver = "postgres"

func NewSQL(ctx context.Context, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, driver, dsn)
	if err != nil {
		return nil, errors.Join(ErrFailedToConnect, err)
	}

	return db, nil
}
