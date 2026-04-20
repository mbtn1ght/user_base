package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"user_base/internal/domain"
	"user_base/internal/dto"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handlers) GetProfile(w http.ResponseWriter, r *http.Request) {
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

	profile, err := h.usecase.GetProfile(r.Context(), profileID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, "profile not found", http.StatusNotFound)
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := dto.GetProfileOutput{
		ID:        profile.ID.String(),
		Name:      string(profile.Name),
		Age:       int(profile.Age),
		Email:     profile.Contacts.Email,
		Phone:     profile.Contacts.Phone,
		CreatedAt: profile.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
