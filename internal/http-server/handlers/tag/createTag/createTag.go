package createTag

import (
	"PartyRoom.API/internal/domain"
	"PartyRoom.API/internal/lib/api/response"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

type Request struct {
	Name   string `json:"name"`
	IsGood bool   `json:"isGood"`
}
type SaverTag interface {
	SaveTag(tag domain.Tag) error
}

func New(saver SaverTag) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tag domain.Tag
		ctx := r.Context()

		userIdString, ok := ctx.Value("userId").(string)
		if !ok {
			http.Error(w, "Failed to get user from context", http.StatusInternalServerError)
			return
		}
		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			http.Error(w, "Failed body", http.StatusBadRequest)
			return
		}
		userId, _ := uuid.Parse(userIdString)
		tag = domain.Tag{
			UserID: &userId,
			Name:   req.Name,
			IsGood: req.IsGood,
		}

		err = saver.SaveTag(tag)
		if err != nil {
			http.Error(w, "Failed save tag", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, response.Ok())
	}
}
