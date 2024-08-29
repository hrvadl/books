package book

import "github.com/hrvadl/book-service/internal/domain/author"

type Book struct {
	ID      int
	Authors []author.Author
	Title   string
	Genres  []string
}
