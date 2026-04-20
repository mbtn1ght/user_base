package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"user_base/internal/domain"
	"user_base/internal/dto"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := chi.URLParam(r, "profileID")

	var request dto.UpdateProfileInput
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	if id != "" {
		if request.ID != "" && request.ID != id {
			http.Error(w, "profile id mismatch", http.StatusBadRequest)
			return
		}

		request.ID = id
	}

	if request.ID == "" {
		http.Error(w, "profile id is required", http.StatusBadRequest)
		return
	}

	profile, err := h.usecase.UpdateProfile(r.Context(), request)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrAllFieldsForUpdateEmpty):
			http.Error(w, "no fields to update", http.StatusBadRequest)
			return
		case errors.Is(err, domain.ErrNotFound):
			http.Error(w, "profile not found", http.StatusNotFound)
			return
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
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
