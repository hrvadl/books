package db

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const driver = "postgres"

func NewSQL(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		return nil, errors.Join(ErrFailedToConnect, err)
	}

	return db, nil
}
