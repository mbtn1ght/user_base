package v1

import (
	"context"

	http_server "gitlab.golang-school.ru/potok-2/lessons/lesson-22/gen/http/profile_v1/server"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/pkg/render"
)

func (h *Handlers) DeleteProfileByID(ctx context.Context, request http_server.DeleteProfileByIDRequestObject,
) (http_server.DeleteProfileByIDResponseObject, error) {
	input := dto.DeleteProfileInput{
		ID: request.ID.String(),
	}

	err := h.usecase.DeleteProfile(ctx, input)
	if err != nil {
		err = render.Error(ctx, err, "request failed")

		return http_server.DeleteProfileByID400JSONResponse{Error: err.Error()}, nil //nolint:nilerr
	}

	return http_server.DeleteProfileByID204Response{}, nil
}
