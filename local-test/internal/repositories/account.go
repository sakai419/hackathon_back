package repositories

import (
	"context"
	"fmt"
	"local-test/internal/models"
	sqlcgen "local-test/internal/sqlc/generated"
	"local-test/pkg/utils"
)

func (r *Repository) CreateAccount(ctx context.Context, arg *models.CreateAccountParams) error {
    // Begin transaction
    tx, err := r.db.Begin()
    if err != nil {
        return utils.WrapRepositoryError(err, "failed to begin transaction")
    }

    // Create query object with transaction
    q := r.q.WithTx(tx)

    // Create account
    accountParams := sqlcgen.CreateAccountParams{
        ID:       arg.ID,
        UserID:   arg.UserID,
        UserName: arg.UserName,
    }
    if err := q.CreateAccount(ctx, accountParams); err != nil {
        tx.Rollback()
        return utils.WrapRepositoryError(err, "failed to create account")
    }

    // Create empty profile
    if err := q.CreateProfilesWithDefaultValues(ctx, arg.ID); err != nil {
        tx.Rollback()
        return utils.WrapRepositoryError(err, "failed to create profile")
    }

    // Create empty setting
    if err := q.CreateSettingsWithDefaultValues(ctx, arg.ID); err != nil {
        tx.Rollback()
        return utils.WrapRepositoryError(err, "failed to create setting")
    }

	// Create empty interest
	if err := q.CreateInterestsWithDefaultValues(ctx, arg.ID); err != nil {
		tx.Rollback()
		return utils.WrapRepositoryError(err, "failed to create interest")
	}

    // Commit transaction
    if err := tx.Commit(); err != nil {
        tx.Rollback()
        return utils.WrapRepositoryError(err, "failed to commit transaction")
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