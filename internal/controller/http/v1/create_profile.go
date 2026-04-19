package v1

import (
	"context"

	http_server "gitlab.golang-school.ru/potok-2/lessons/lesson-22/gen/http/profile_v1/server"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/pkg/render"
)

func (h *Handlers) CreateProfile(ctx context.Context, request http_server.CreateProfileRequestObject,
) (http_server.CreateProfileResponseObject, error) {
	input := dto.CreateProfileInput{
		Name:  request.Body.Name,
		Age:   request.Body.Age,
		Email: request.Body.Email,
		Phone: request.Body.Phone,
	}

	output, err := h.usecase.CreateProfile(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		return http_server.CreateProfile400JSONResponse{Error: err.Error()}, nil //nolint:nilerr
	}

	return http_server.CreateProfile200JSONResponse{
		ID: output.ID,
	}, nil
}
