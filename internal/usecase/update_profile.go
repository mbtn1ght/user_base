package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/internal/domain"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/pkg/otel/tracer"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/pkg/transaction"
)

func (u *UseCase) UpdateProfile(ctx context.Context, input dto.UpdateProfileInput) error {
	ctx, span := tracer.Start(ctx, "usecase UpdateProfile")
	defer span.End()

	err := input.Validate()
	if err != nil {
		return fmt.Errorf("input.Validate: %w", err)
	}

	id, err := uuid.Parse(input.ID)
	if err != nil {
		return domain.ErrUUIDInvalid
	}

	err = transaction.Wrap(ctx, func(ctx context.Context) error {
		profile, err := u.postgres.GetProfile(ctx, id)
		if err != nil {
			return fmt.Errorf("postgres.GetProfile: %w", err)
		}

		if profile.IsDeleted() {
			return domain.ErrNotFound
		}

		newProfile := update(profile, input)

		if newProfile == profile {
			return nil
		}

		err = u.postgres.UpdateProfile(ctx, newProfile)
		if err != nil {
			return fmt.Errorf("postgres.UpdateProfile: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("transaction.Wrap: %w", err)
	}

	return nil
}

func update(profile domain.Profile, input dto.UpdateProfileInput) domain.Profile {
	if input.Name != nil {
		profile.Name = domain.Name(*input.Name)
	}

	if input.Age != nil {
		profile.Age = domain.Age(*input.Age)
	}

	if input.Email != nil {
		profile.Contacts.Email = *input.Email
	}

	if input.Phone != nil {
		profile.Contacts.Phone = *input.Phone
	}

	return profile
}
