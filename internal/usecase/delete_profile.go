package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/internal/domain"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/pkg/otel/tracer"
)

func (u *UseCase) DeleteProfile(ctx context.Context, input dto.DeleteProfileInput) error {
	ctx, span := tracer.Start(ctx, "usecase DeleteProfile")
	defer span.End()

	id, err := uuid.Parse(input.ID)
	if err != nil {
		return domain.ErrUUIDInvalid
	}

	err = u.postgres.DeleteProfile(ctx, id)
	if err != nil {
		return fmt.Errorf("postgres.DeleteProfile: %w", err)
	}

	return nil
}
