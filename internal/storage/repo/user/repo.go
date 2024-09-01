package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/hrvadl/book-service/internal/domain/genre"
	"github.com/hrvadl/book-service/internal/domain/user"
)

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

type Repo struct {
	db *sqlx.DB
}

func (r *Repo) Create(ctx context.Context, u user.User) (int, error) {
	const userCmd = `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`

	var userID int
	if err := r.db.QueryRowContext(ctx, userCmd, u.Name, u.Email).Scan(&userID); err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	const genresCmd = `
		INSERT INTO user_favorite_genres (genre_id, user_id) VALUES (:genre_id, :user_id)
	`

	userPreferredGenres := make([]map[string]any, 0, len(u.PreferredGenres))
	for _, g := range u.PreferredGenres {
		userPreferredGenres = append(
			userPreferredGenres,
			map[string]any{"genre_id": g.ID, "user_id": userID},
		)
	}

	if _, err := r.db.NamedExecContext(ctx, genresCmd, userPreferredGenres); err != nil {
		return 0, fmt.Errorf("failed to create user preferred genres: %w", err)
	}

	return userID, nil
}

func (r *Repo) GetByID(ctx context.Context, id int) (*user.User, error) {
	const userQuery = `SELECT * FROM users WHERE id = $1`

	var u user.User
	if err := r.db.GetContext(ctx, &u, userQuery, id); err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	const genresQuery = `
		SELECT g.id, g.name
		FROM genres g
		JOIN user_favorite_genres ug ON ug.genre_id = g.id
		JOIN users u ON ug.user_id = u.id AND u.id = $1
	`

	var genres []genre.Genre
	if err := r.db.SelectContext(ctx, &genres, genresQuery, id); err != nil {
		return nil, fmt.Errorf("failed to get users preferred genres: %w", err)
	}

	u.PreferredGenres = genres

	return &u, nil
}
