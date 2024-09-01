package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/hrvadl/book-service/internal/domain/genre"
)

func NewService(us UserSource, gs GenresSource) *Service {
	return &Service{
		users:  us,
		genres: gs,
	}
}

//go:generate mockgen -destination=./mocks/mock_users.go -package=mocks . UserSource
type UserSource interface {
	GetByID(ctx context.Context, id int) (*User, error)
	Create(ctx context.Context, u User) (int, error)
}

//go:generate mockgen -destination=./mocks/mock_users.go -package=mocks . UserSource
type GenresSource interface {
	GetByNames(ctx context.Context, names []string) ([]genre.Genre, error)
}

type Service struct {
	users  UserSource
	genres GenresSource
}

type CreateUserCmd struct {
	Name           string
	Email          string
	FavoriteGenres []string
}

func (s *Service) Create(ctx context.Context, cmd CreateUserCmd) (int, error) {
	genres, err := s.genres.GetByNames(ctx, cmd.FavoriteGenres)
	if err != nil {
		return 0, fmt.Errorf("%w: failed to get genres: %w", ErrFailedToCreate, err)
	}

	u := User{
		Name:            cmd.Name,
		Email:           cmd.Email,
		PreferredGenres: genres,
	}

	id, err := s.users.Create(ctx, u)
	if err != nil {
		return 0, errors.Join(ErrFailedToCreate, err)
	}

	return id, nil
}

func (s *Service) GetByID(ctx context.Context, id int) (*User, error) {
	u, err := s.users.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Join(ErrFailedToGet, err)
	}

	return u, nil
}
