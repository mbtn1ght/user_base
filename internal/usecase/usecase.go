package usecase

import (
	"context"

	"user_base/internal/domain"
)

type PostgresRepository interface {
	CreateProfile(ctx context.Context, profile domain.Profile) error
	CreateProperty(ctx context.Context, property domain.Property) error
	GetProfile(ctx context.Context, profileID string) (domain.Profile, error)
}

type UseCase struct {
	postgres PostgresRepository
}

func New(postgres PostgresRepository) *UseCase {
	return &UseCase{postgres: postgres}
}
