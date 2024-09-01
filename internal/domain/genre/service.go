package genre

import (
	"context"
	"errors"
)

func NewService(gs GenreSource) *Service {
	return &Service{
		genres: gs,
	}
}

type GenreSource interface {
	GetByNames(ctx context.Context, names []string) ([]Genre, error)
}

type Service struct {
	genres GenreSource
}

func (s *Service) GetByNames(ctx context.Context, names []string) ([]Genre, error) {
	genres, err := s.genres.GetByNames(ctx, names)
	if err != nil {
		return nil, errors.Join(ErrFailedToGet, err)
	}

	return genres, nil
}
