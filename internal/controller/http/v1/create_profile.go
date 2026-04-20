package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"user_base/internal/dto"
	"user_base/internal/usecase"
)

type Handlers struct {
	usecase *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{usecase: uc}
}

func (h *Handlers) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateProfileInput
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	output, err := h.usecase.CreateProfile(context.Background(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // Успешный ответ
	json.NewEncoder(w).Encode(output)
}
