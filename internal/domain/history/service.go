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
