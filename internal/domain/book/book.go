package book

import (
	"github.com/hrvadl/book-service/internal/domain/author"
	"github.com/hrvadl/book-service/internal/domain/genre"
)

type Book struct {
	ID      int
	Authors []author.Author
	Title   string
	Genres  []genre.Genre
}
