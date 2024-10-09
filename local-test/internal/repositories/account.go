package repositories

import (
	"context"
	"fmt"
	"local-test/internal/sqlc/generated"
	"local-test/internal/models"
	"local-test/pkg/utils"
)

func (r *Repository) CreateAccount(ctx context.Context, arg *models.CreateAccountParams) (error) {
	params := sqlcgen.CreateAccountParams{
		ID:       arg.ID,
		UserID:   arg.UserID,
		UserName: arg.UserName,
	}

	if err := r.q.CreateAccount(ctx, params); err != nil {
		return utils.WrapRepositoryError(err, "failed to create account")
	}
	return nil
}

func (r *Repository) DeleteMyAccount(ctx context.Context, id string) (error) {
	res, err := r.q.DeleteAccount(ctx, id)
	if err != nil {
		return utils.WrapRepositoryError(err, "failed to delete account")
	}

	num, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if num == 0 {
		return ErrAccountNotFound
	}
	return nil
}