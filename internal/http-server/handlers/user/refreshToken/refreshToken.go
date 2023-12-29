package refreshToken

import (
	"PartyRoom.API/internal/config"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type Response struct {
	AccessToken string `json:"accessToken"`
}
type AuthService interface {
	RefreshToken(userId uuid.UUID, token string) string
}

func New(authService AuthService, cfg config.Config) http.HandlerFunc {
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
		accessToken := authService.RefreshToken(userId, refreshTokenCookie)
		if accessToken == "" {
			http.Error(w, "Invalid refresh token", http.StatusBadRequest)
			return
		}
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
