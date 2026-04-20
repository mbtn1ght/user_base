package v1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"user_base/internal/domain"
	"user_base/internal/dto"
)

func (h *Handlers) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "profileID")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	input := dto.DeleteProfileInput{ID: id}

	if err := h.usecase.DeleteProfile(r.Context(), input); err != nil {
		switch {
		case errors.Is(err, domain.ErrUUIDInvalid):
			http.Error(w, "invalid profile id", http.StatusBadRequest)
			return
		case errors.Is(err, domain.ErrNotFound):
			http.Error(w, "profile not found", http.StatusNotFound)
			return
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
