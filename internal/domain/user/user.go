package user

import "github.com/hrvadl/book-service/internal/domain/genre"

type User struct {
	ID              int           `db:"id"`
	Name            string        `db:"name"`
	Email           string        `db:"email"`
	PreferredGenres []genre.Genre `db:"genres"`
}
