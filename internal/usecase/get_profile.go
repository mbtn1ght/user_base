package usecase

import (
	"context"
	"errors"
	"fmt"

	"user_base/internal/domain"

	"github.com/google/uuid"
)

func (u *UseCase) GetProfile(ctx context.Context, profileID uuid.UUID) (domain.Profile, error) {
	profile, err := u.postgres.GetProfile(ctx, profileID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return domain.Profile{}, err
		}

		return domain.Profile{}, fmt.Errorf("postgres.GetProfile: %w", err)
	}

	return profile, nil
}
