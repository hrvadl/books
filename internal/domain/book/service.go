package book

import (
	"context"
	"errors"
)

func NewService(bs BookSaver) *Service {
	return &Service{
		books: bs,
	}
}

type Service struct {
	books BookSaver
}

//go:generate mockgen -destination=./mocks/mock_saver.go -package=mocks . BookSaver
type BookSaver interface {
	Save(ctx context.Context, book Book) (int, error)
}

func (s *Service) Add(ctx context.Context, book Book) (int, error) {
	id, err := s.books.Save(ctx, book)
	if err != nil {
		return 0, errors.Join(ErrFailedToAdd, err)
	}

	return id, nil
}
