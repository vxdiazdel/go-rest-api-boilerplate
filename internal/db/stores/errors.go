package stores

import "errors"

var (
	ErrUserEmailTaken = errors.New("user email already exists")
	ErrUserNotFound   = errors.New("user not found")
)
