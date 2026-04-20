package usecase

import (
	"context"
	"errors"
	"fmt"

	"user_base/internal/domain"

	"github.com/google/uuid"
)

func (u *UseCase) DeleteProfile(ctx context.Context, profileID uuid.UUID) error {
	if err := u.postgres.DeleteProfile(ctx, profileID); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return err
		}

		return fmt.Errorf("postgres.DeleteProfile: %w", err)
	}

	return nil
}
