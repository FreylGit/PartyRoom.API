package authService

import (
	"PartyRoom.API/internal/domain"
	"PartyRoom.API/internal/service/authService/tokenManager"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Repository interface {
	GetUserByEmail(email string) (domain.User, error)
	GetUserById(uuid uuid.UUID) (domain.User, error)
	GetRefreshToken(userId uuid.UUID) (domain.RefreshToken, error)
	UpdateRefreshToken(token domain.RefreshToken) error
}
type AuthService struct {
	Repository Repository
}

func New(repository Repository) AuthService {
	return AuthService{
		Repository: repository,
	}
}
func (as *AuthService) SignIn(email string, password string) (string, string) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {

	}
	userFind, err := as.Repository.GetUserByEmail(email)
	if err != nil {

	}
	if userFind.PasswordHash == string(passwordHash) {

	}
	accessToken, _ := tokenManager.GenerateAccessToken(userFind)
	if accessToken == "" {
		return "", ""
	}

	refreshToken := tokenManager.GenerateRefreshToken(userFind.ID.String())
	refreshTokenModel := domain.RefreshToken{
		Token: refreshToken,
		User:  &userFind,
	}
	err = as.Repository.UpdateRefreshToken(refreshTokenModel)
	if err != nil {
		return "", ""
	}
	return accessToken, refreshToken
}

func (as *AuthService) RefreshToken(userId uuid.UUID, token string) string {
	user, err := as.Repository.GetUserById(userId)
	if err != nil {
		return ""
	}
	refreshToken, err := as.Repository.GetRefreshToken(*user.ID)
	if err != nil {
		return ""
	}
	test := refreshToken.ExpirationDate.Sub(time.Now())
	_ = test
	if refreshToken.ExpirationDate.Sub(time.Now()) < 0 {
		return ""
	}
	if refreshToken.Token == token {
		accessToken, _ := tokenManager.GenerateAccessToken(user)
		return accessToken
	}
	return ""
}
