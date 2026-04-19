package usecase

import (
	"context"

	"github.com/google/uuid"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/internal/domain"
)

//go:generate mockery

type Redis interface {
	IsExists(ctx context.Context, idempotencyKey string) bool
}

type Postgres interface {
	CreateProfile(ctx context.Context, profile domain.Profile) error
	CreateProperty(ctx context.Context, property domain.Property) error
	GetProfile(ctx context.Context, profileID uuid.UUID) (domain.Profile, error)
	UpdateProfile(ctx context.Context, profile domain.Profile) error
	DeleteProfile(ctx context.Context, id uuid.UUID) error

	ReadOutboxKafka(ctx context.Context, limit int) ([]domain.Event, error)
	SaveOutboxKafka(ctx context.Context, events ...domain.Event) error
}

type Kafka interface {
	Produce(ctx context.Context, events ...domain.Event) error
}

type UseCase struct {
	postgres Postgres
	redis    Redis
	kafka    Kafka
}

func New(postgres Postgres, redis Redis, kafka Kafka) *UseCase {
	return &UseCase{
		postgres: postgres,
		redis:    redis,
		kafka:    kafka,
	}
}
