package usecase

import (
	"context"
	"fmt"

	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/internal/domain"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/internal/dto"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/pkg/otel/tracer"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/pkg/transaction"
)

func (u *UseCase) CreateProfile(ctx context.Context, input dto.CreateProfileInput) (dto.CreateProfileOutput, error) {
	ctx, span := tracer.Start(ctx, "usecase CreateProfile")
	defer span.End()

	var output dto.CreateProfileOutput

	profile, err := domain.NewProfile(input.Name, input.Age, input.Email, input.Phone)
	if err != nil {
		return output, fmt.Errorf("domain.NewProfile: %w", err)
	}

	event, err := profile.ToEvent("awesome-topic")
	if err != nil {
		return output, fmt.Errorf("profile.ToEvent: %w", err)
	}

	property := domain.NewProperty(profile.ID, []string{"home", "primary"})

	err = transaction.Wrap(ctx, func(ctx context.Context) error {
		err = u.postgres.CreateProfile(ctx, profile)
		if err != nil {
			return fmt.Errorf("postgres.CreateProfile: %w", err)
		}

		err = u.postgres.CreateProperty(ctx, property)
		if err != nil {
			return fmt.Errorf("postgres.CreateProperty: %w", err)
		}

		// Запись в таблицу Outbox (из которой читает воркер и гарантировано отправляет в Кафку)
		err = u.postgres.SaveOutboxKafka(ctx, event)
		if err != nil {
			return fmt.Errorf("postgres.SaveOutboxKafka: %w", err)
		}

		return nil
	})
	if err != nil {
		return output, fmt.Errorf("transaction.Wrap: %w", err)
	}

	// Сразу отправляем в Кафку (но гарантии отправки нет)
	//err = u.kafka.Produce(ctx, event)
	//if err != nil {
	//	log.Error().Err(err).Msg("usecase CreateProfile: kafka.Produce")
	//}

	return dto.CreateProfileOutput{
		ID: profile.ID,
	}, nil
}
