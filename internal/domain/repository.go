package domain

import "github.com/google/uuid"

type UserRepository interface {
	GetUserByEmail(email string) (User, error)
	GetUserById(uuid uuid.UUID) (User, error)
}

type RefreshTokenRepository interface {
	GetRefreshToken(userId uuid.UUID) (RefreshToken, error)
	UpdateRefreshToken(token RefreshToken) error
}
