package signIn

import (
	"PartyRoom.API/internal/domain"
	"PartyRoom.API/internal/lib/api/response"
	"PartyRoom.API/internal/service/authService"
	"github.com/go-chi/render"
	"net/http"
)

type Request struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Response struct {
	AccessToken string `json:"accessToken"`
}

func New(userRepository domain.UserRepository, tokenRepository domain.RefreshTokenRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			render.JSON(w, r, response.Error("invalid request"))
			return
		}

		accessToken, refreshToken := authService.SignIn(userRepository, tokenRepository, req.Email, req.Password)
		response := Response{AccessToken: accessToken}
		cookie := http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, &cookie)
		render.JSON(w, r, response)
	}
}
