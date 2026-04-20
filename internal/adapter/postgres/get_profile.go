package postgres

import (
	"context"
	"errors"
	"fmt"

	"user_base/internal/domain"
	"user_base/pkg/transaction"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (p *Postgres) GetProfile(ctx context.Context, profileID uuid.UUID) (domain.Profile, error) {
	const sql = `SELECT created_at, name, age, email, phone
                    FROM profiles WHERE id = $1`

	dto := struct {
		CreatedAt pgtype.Timestamptz
		Name      pgtype.Text
		Age       pgtype.Int4
		Email     pgtype.Text
		Phone     pgtype.Text
	}{}

	dest := []any{
		&dto.CreatedAt,
		&dto.Name,
		&dto.Age,
		&dto.Email,
		&dto.Phone,
	}

	txOrPool := transaction.TryExtractTX(ctx)

	if err := txOrPool.QueryRow(ctx, sql, profileID).Scan(dest...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Profile{}, domain.ErrNotFound
		}

		return domain.Profile{}, fmt.Errorf("txOrPool.QueryRow: %w", err)
	}

	phone := ""
	if dto.Phone.Valid {
		phone = dto.Phone.String
	}

	profile := domain.Profile{
		ID:        profileID,
		CreatedAt: dto.CreatedAt.Time,
		Name:      domain.Name(dto.Name.String),
		Age:       domain.Age(dto.Age.Int32),
		Status:    domain.Pending,
		Verified:  false,
		Contacts: domain.Contacts{
			Email: dto.Email.String,
			Phone: phone,
		},
	}

	return profile, nil
}
