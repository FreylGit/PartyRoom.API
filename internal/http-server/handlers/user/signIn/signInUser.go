package signIn

import (
	"PartyRoom.API/internal/lib/api/response"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

type Request struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Response struct {
	AccessToken string `json:"accessToken"`
}
type AuthService interface {
	SignIn(email string, password string) (string, string)
}

func New(authService AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			render.JSON(w, r, response.Error("invalid request"))
			return
		}

		accessToken, refreshToken := authService.SignIn(req.Email, req.Password)
		response := Response{AccessToken: accessToken}
		cookie := http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken,
			HttpOnly: true,
			Path:     "/",
			Expires:  time.Now().Add(22 * time.Hour),
		}
		http.SetCookie(w, &cookie)
		render.JSON(w, r, response)
	}
}
