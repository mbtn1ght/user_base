package usecase

import (
	"context"

	"user_base/internal/domain"
	"user_base/internal/dto"

	"github.com/google/uuid"
)

type Postgres interface {
	CreateProfile(ctx context.Context, profile domain.Profile) error
	CreateProperty(ctx context.Context, property domain.Property) error
	GetProfile(ctx context.Context, profileID uuid.UUID) (domain.Profile, error)
	UpdateProfile(ctx context.Context, profileID uuid.UUID, input dto.UpdateProfileInput) error
	DeleteProfile(ctx context.Context, profileID uuid.UUID) error
}

type UseCase struct {
	postgres Postgres
}

func New(postgres Postgres) *UseCase {
	return &UseCase{postgres: postgres}
}
