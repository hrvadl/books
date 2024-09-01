package history

import "errors"

var (
	ErrFailedToGet = errors.New("failed to get reading history")
	ErrFailedToAdd = errors.New("failed to add history")
)
