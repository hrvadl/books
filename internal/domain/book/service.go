package book

import (
	"context"
	"errors"
)

func NewService(bs BookSource) *Service {
	return &Service{
		books: bs,
	}
}

type Service struct {
	books BookSource
}

//go:generate mockgen -destination=./mocks/mock_saver.go -package=mocks . BookSource
type BookSource interface {
	Save(ctx context.Context, book Book) (int, error)
	GetAll(ctx context.Context) ([]Book, error)
}

func (s *Service) Add(ctx context.Context, book Book) (int, error) {
	id, err := s.books.Save(ctx, book)
	if err != nil {
		return 0, errors.Join(ErrFailedToAdd, err)
	}

	return id, nil
}

func (s *Service) GetAll(ctx context.Context) ([]Book, error) {
	b, err := s.books.GetAll(ctx)
	if err != nil {
		return nil, errors.Join(ErrFailedToGet, err)
	}

	return b, nil
}
