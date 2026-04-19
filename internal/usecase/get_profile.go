package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/internal/domain"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/pkg/otel/tracer"
)

func (u *UseCase) GetProfile(ctx context.Context, input dto.GetProfileInput) (dto.GetProfileOutput, error) {
	ctx, span := tracer.Start(ctx, "usecase GetProfile")
	defer span.End()

	var output dto.GetProfileOutput

	id, err := uuid.Parse(input.ID)
	if err != nil {
		return output, domain.ErrUUIDInvalid
	}

	profile, err := u.postgres.GetProfile(ctx, id)
	if err != nil {
		return output, fmt.Errorf("postgres.GetProfile: %w", err)
	}

	if profile.IsDeleted() {
		return output, domain.ErrNotFound
	}

	return dto.GetProfileOutput{
		Profile: profile,
	}, nil
}
