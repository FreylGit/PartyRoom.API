package registrationUser

import (
	"PartyRoom.API/internal/domain"
	"PartyRoom.API/internal/lib/api/response"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Request struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}
type UserRepository interface {
	SaveUser(user domain.User) error
}

func New(repository UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.JSON(w, r, response.Error("invalid request"))
			return
		}
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		newUser := domain.User{
			Email:        req.Email,
			PasswordHash: string(passwordHash),
			Name:         req.Name,
			Phone:        req.Phone,
		}
		err = repository.SaveUser(newUser)
		if err != nil {
			render.JSON(w, r, response.Error(err.Error()))
			return
		}
		render.JSON(w, r, response.Ok())
	}
}
