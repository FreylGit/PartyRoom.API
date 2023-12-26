package profile

import (
	"PartyRoom.API/internal/domain"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

type Response struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

func New(userRepository domain.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userIdString, ok := ctx.Value("userId").(string)
		if !ok {
			http.Error(w, "Failed to get user from context", http.StatusInternalServerError)
			return
		}
		userId, _ := uuid.Parse(userIdString)
		user, _ := userRepository.GetUserById(userId)
		response := Response{
			Email: user.Email,
			Name:  user.Name,
			Photo: user.Photo,
		}
		render.JSON(w, r, response)
	}
}
