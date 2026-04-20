package usecase

import (
	"context"
	"errors"
	"fmt"
	"user_base/internal/dto"

	"user_base/internal/domain"

	"github.com/google/uuid"
)

func (u *UseCase) DeleteProfile(ctx context.Context, input dto.DeleteProfileInput) error {
	profileID, err := uuid.Parse(input.ID)
	if err != nil {
		return domain.ErrUUIDInvalid
	}

	if err := u.postgres.DeleteProfile(ctx, profileID); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return err
		}

		return fmt.Errorf("postgres.DeleteProfile: %w", err)
	}

	return nil
}
