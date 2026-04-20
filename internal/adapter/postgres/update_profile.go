package postgres

import (
	"context"
	"fmt"
	"strings"

	"user_base/internal/domain"
	"user_base/internal/dto"
	"user_base/pkg/transaction"

	"github.com/google/uuid"
)

func (p *Postgres) UpdateProfile(ctx context.Context, profileID uuid.UUID, input dto.UpdateProfileInput) error {
	if input.IsEmpty() {
		return domain.ErrAllFieldsForUpdateEmpty
	}

	setParts := make([]string, 0, 4)
	args := make([]any, 0, 5)
	paramIdx := 1

	if input.Name != nil {
		setParts = append(setParts, fmt.Sprintf("name = $%d", paramIdx))
		args = append(args, *input.Name)
		paramIdx++
	}

	if input.Age != nil {
		setParts = append(setParts, fmt.Sprintf("age = $%d", paramIdx))
		args = append(args, *input.Age)
		paramIdx++
	}

	if input.Email != nil {
		setParts = append(setParts, fmt.Sprintf("email = $%d", paramIdx))
		args = append(args, *input.Email)
		paramIdx++
	}

	if input.Phone != nil {
		setParts = append(setParts, fmt.Sprintf("phone = $%d", paramIdx))
		args = append(args, *input.Phone)
		paramIdx++
	}

	if len(setParts) == 0 {
		return domain.ErrAllFieldsForUpdateEmpty
	}

	args = append(args, profileID)

	sql := fmt.Sprintf(`UPDATE profiles SET %s WHERE id = $%d`, strings.Join(setParts, ", "), paramIdx)

	txOrPool := transaction.TryExtractTX(ctx)

	cmdTag, err := txOrPool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}

	return nil
}
