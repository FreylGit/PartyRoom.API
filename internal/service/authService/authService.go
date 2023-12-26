package authService

import (
	"PartyRoom.API/internal/domain"
	"PartyRoom.API/internal/service/authService/tokenManager"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(userRepository domain.UserRepository, tokenRepository domain.RefreshTokenRepository, email string, password string) (string, string) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {

	}
	userFind, err := userRepository.GetUserByEmail(email)
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
		User:  userFind,
	}
	err = tokenRepository.UpdateRefreshToken(refreshTokenModel)
	if err != nil {
		return "", ""
	}
	return accessToken, refreshToken
}

func RefreshToken(repository domain.RefreshTokenRepository, user *domain.User, token string) string {
	refreshToken, err := repository.GetRefreshToken(*user.ID)
	if err != nil {
		return ""
	}
	if refreshToken.Token == token {
		accessToken, _ := tokenManager.GenerateAccessToken(user)
		return accessToken
	}
	return ""
}
