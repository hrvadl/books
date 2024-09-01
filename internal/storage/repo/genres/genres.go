package genres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/hrvadl/book-service/internal/domain/genre"
)

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

type Repo struct {
	db *sqlx.DB
}

func (r *Repo) GetByNames(ctx context.Context, names []string) ([]genre.Genre, error) {
	const query = `SELECT * FROM genres WHERE name IN (?)`

	q, args, err := sqlx.In(query, names)
	if err != nil {
		return nil, fmt.Errorf("failed to construct in query: %w", err)
	}

	q = r.db.Rebind(q)

	var genres []genre.Genre
	if err := r.db.SelectContext(ctx, &genres, q, args...); err != nil {
		return nil, fmt.Errorf("failed to select genres by name: %w", err)
	}

	return genres, nil
}
