package user

import (
	"context"
	"errors"
)

func NewService(us UserSource) *Service {
	return &Service{
		users: us,
	}
}

//go:generate mockgen -destination=./mocks/mock_users.go -package=mocks . UserSource
type UserSource interface {
	GetByID(ctx context.Context, id int) (*User, error)
}

type Service struct {
	users UserSource
}

func (s *Service) GetByID(ctx context.Context, id int) (*User, error) {
	u, err := s.users.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Join(ErrFailedToGet, err)
	}

	return u, nil
}
