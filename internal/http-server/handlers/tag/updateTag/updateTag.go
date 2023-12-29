package updateTag

import (
	"PartyRoom.API/internal/domain"
	"PartyRoom.API/internal/lib/api/response"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

type Request struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	IsGood bool      `json:"isGood"`
}
type UpdaterTag interface {
	UpdateTag(tag domain.Tag) error
}

func New(updater UpdaterTag) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		ctx := r.Context()

		userIdString, ok := ctx.Value("userId").(string)
		if !ok {
			http.Error(w, "Failed to get user from context", http.StatusInternalServerError)
			return
		}
		userId, _ := uuid.Parse(userIdString)
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			http.Error(w, "Failed body", http.StatusBadRequest)
			return
		}
		tag := domain.Tag{
			ID:     &req.ID,
			UserID: &userId,
			Name:   req.Name,
			IsGood: req.IsGood,
		}

		err = updater.UpdateTag(tag)
		if err != nil {
			http.Error(w, "Failed update tag", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, response.Ok())
	}
}
