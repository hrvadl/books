package author

import (
	"context"
	"errors"
)

func NewService(as AuthorSaver) *Service {
	return &Service{
		authors: as,
	}
}

//go:generate mockgen -destination=./mocks/mock_saver.go -package=mocks . AuthorSaver
type AuthorSaver interface {
	Save(ctx context.Context, author Author) (int, error)
}

type Service struct {
	authors AuthorSaver
}

func (s *Service) Add(ctx context.Context, author Author) (int, error) {
	id, err := s.authors.Save(ctx, author)
	if err != nil {
		return 0, errors.Join(ErrFailedToAdd, err)
	}

	return id, err
}
