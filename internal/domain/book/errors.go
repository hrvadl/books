package book

import "errors"

var (
	ErrFailedToAdd = errors.New("failed to add book")
	ErrFailedToGet = errors.New("failed to get books")
)
