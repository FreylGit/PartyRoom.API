package deleteTag

import (
	"PartyRoom.API/internal/lib/api/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

type DeleterTag interface {
	DeleteTag(tagID uuid.UUID, userID uuid.UUID) error
}

func New(deleter DeleterTag) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userIdString, ok := ctx.Value("userId").(string)
		if !ok {
			http.Error(w, "Failed to get user from context", http.StatusInternalServerError)
			return
		}
		userId, _ := uuid.Parse(userIdString)
		// Извлекаем параметры из URL
		tagIDString := chi.URLParam(r, "tagID")
		if tagIDString == "" {
			http.Error(w, "Tag ID not found in URL", http.StatusBadRequest)
			return
		}
		tagID, err := uuid.Parse(tagIDString)
		if err != nil {
			http.Error(w, "Invalid Tag ID in URL", http.StatusBadRequest)
			return
		}
		err = deleter.DeleteTag(tagID, userId)
		if err != nil {
			http.Error(w, "Failed to delete tag: "+err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, response.Ok())
	}
}
