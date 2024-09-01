package user

import "errors"

var (
	ErrFailedToGet    = errors.New("failed to get user")
	ErrFailedToCreate = errors.New("failed to create user")
)
