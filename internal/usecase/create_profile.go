package usecase

import (
	"context"
	"fmt"

	"user_base/internal/domain"
	"user_base/internal/dto"
	"user_base/pkg/transaction"
)

func (u *UseCase) CreateProfile(ctx context.Context, input dto.CreateProfileInput) (dto.CreateProfileOutput, error) {
	var output dto.CreateProfileOutput

	profile, err := domain.NewProfile(input.Name, input.Age, input.Email, input.Phone)
	if err != nil {
		return output, fmt.Errorf("domain.NewProfile: %w", err)
	}

	property := domain.NewProperty(profile.ID, []string{"home", "primary"})

	err = transaction.Wrap(ctx, func(ctx context.Context) error {
		if err := u.postgres.CreateProfile(ctx, profile); err != nil {
			return fmt.Errorf("postgres.CreateProfile: %w", err)
		}

		if err := u.postgres.CreateProperty(ctx, property); err != nil {
			return fmt.Errorf("postgres.CreateProperty: %w", err)
		}

		return nil
	})
	if err != nil {
		return output, fmt.Errorf("transaction.Wrap: %w", err)
	}

	return dto.CreateProfileOutput{
		ID: profile.ID.String(),
	}, nil
}
