package usecase

import (
	"context"
	"errors"
	"fmt"

	"user_base/internal/domain"
	"user_base/internal/dto"

	"github.com/google/uuid"
)

func (u *UseCase) UpdateProfile(ctx context.Context, profileID uuid.UUID, input dto.UpdateProfileInput) (domain.Profile, error) {
	if input.IsEmpty() {
		return domain.Profile{}, domain.ErrAllFieldsForUpdateEmpty
	}

	if err := u.postgres.UpdateProfile(ctx, profileID, input); err != nil {
		if errors.Is(err, domain.ErrNotFound) || errors.Is(err, domain.ErrAllFieldsForUpdateEmpty) {
			return domain.Profile{}, err
		}

		return domain.Profile{}, fmt.Errorf("postgres.UpdateProfile: %w", err)
	}

	profile, err := u.postgres.GetProfile(ctx, profileID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return domain.Profile{}, err
		}

		return domain.Profile{}, fmt.Errorf("postgres.GetProfile: %w", err)
	}

	return profile, nil
}
