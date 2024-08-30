package review

import (
	"context"
	"errors"
)

func NewService(rs ReviewSource) *Service {
	return &Service{
		reviews: rs,
	}
}

//go:generate mockgen -destination=./mocks/mock_review.go -package=mocks . ReviewSource
type ReviewSource interface {
	GetByUserID(ctx context.Context, userID int) ([]Review, error)
}

type Service struct {
	reviews ReviewSource
}

func (s *Service) GetByUserID(ctx context.Context, userID int) ([]Review, error) {
	reviews, err := s.reviews.GetByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Join(ErrFailedToGet, err)
	}

	return reviews, nil
}
