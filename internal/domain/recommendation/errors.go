package recommendation

import "errors"

var (
	ErrFailedToGet      = errors.New("failed to get recommendation")
	ErrNoRecommendation = errors.New("no recommendation suitable")
)
