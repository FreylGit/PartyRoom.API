package refreshToken

import (
	"PartyRoom.API/internal/config"
	"PartyRoom.API/internal/domain"
	"PartyRoom.API/internal/service/authService"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type Response struct {
	AccessToken string `json:"accessToken"`
}

func New(userRepository domain.UserRepository, tokenRepository domain.RefreshTokenRepository, cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refreshTokenCookie, err := getRefreshTokenFromCookie(r)
		if !strings.Contains(refreshTokenCookie, "_") {
			http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
			return
		}
		userIdString := strings.Split(refreshTokenCookie, "_")[1]
		userId, err := uuid.Parse(userIdString)
		if err != nil {
			http.Error(w, "Invalid user id", http.StatusUnauthorized)
			return
		}
		user, err := userRepository.GetUserById(userId)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		accessToken := authService.RefreshToken(tokenRepository, user, refreshTokenCookie)
		response := Response{AccessToken: accessToken}
		render.JSON(w, r, response)
	}
}

func getRefreshTokenFromCookie(r *http.Request) (string, error) {
	refreshTokenCookie, err := r.Cookie("refreshToken")
	if err != nil {
		return "", err
	}
	return refreshTokenCookie.Value, nil
}
