package history

import (
	"context"
	"errors"
)

func NewService(hs HistorySource) *Service {
	return &Service{
		history: hs,
	}
}

//go:generate mockgen -destination=./mocks/mock_history.go -package=mocks . HistorySource
type HistorySource interface {
	GetByUserID(ctx context.Context, userID int) ([]ReadingHistory, error)
	Add(ctx context.Context, h ReadingHistory) (string, error)
}

type Service struct {
	history HistorySource
}

func (s *Service) GetByUserID(ctx context.Context, userID int) ([]ReadingHistory, error) {
	h, err := s.history.GetByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Join(ErrFailedToGet, err)
	}

	return h, nil
}

func (s *Service) Add(ctx context.Context, h ReadingHistory) (string, error) {
	id, err := s.history.Add(ctx, h)
	if err != nil {
		return "", errors.Join(ErrFailedToAdd, err)
	}

	return id, nil
}
