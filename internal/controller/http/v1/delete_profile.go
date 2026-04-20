package v1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"user_base/internal/domain"
)

func (h *Handlers) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "profileID")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	profileID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "invalid profile id", http.StatusBadRequest)
		return
	}

	if err := h.usecase.DeleteProfile(r.Context(), profileID); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, "profile not found", http.StatusNotFound)
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
