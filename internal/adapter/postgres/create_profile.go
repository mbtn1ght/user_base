package postgres

import (
	"context"
	"fmt"

	"user_base/internal/domain"
	"user_base/pkg/transaction"
)

func (p *Postgres) CreateProfile(ctx context.Context, profile domain.Profile) error {
	const sql = `INSERT INTO profiles (id, name, age, email, phone)
                    VALUES ($1, $2, $3, $4, $5)`

	args := []any{
		profile.ID,
		string(profile.Name),
		int(profile.Age),
		profile.Contacts.Email,
		profile.Contacts.Phone,
	}

	txOrPool := transaction.TryExtractTX(ctx)

	if _, err := txOrPool.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}
