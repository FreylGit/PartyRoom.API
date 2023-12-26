package storage

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user exists")
	ErrCreateUser   = errors.New("user do not created")
)
