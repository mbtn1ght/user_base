package postgres

import (
	"context"
	"fmt"

	"user_base/internal/domain"
	"user_base/pkg/transaction"

	"github.com/google/uuid"
)

func (p *Postgres) DeleteProfile(ctx context.Context, profileID uuid.UUID) error {
	const sql = `DELETE FROM profiles WHERE id = $1`

	txOrPool := transaction.TryExtractTX(ctx)

	cmdTag, err := txOrPool.Exec(ctx, sql, profileID)
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}

	return nil
}
