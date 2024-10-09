package repositories

import (
	"context"
	"local-test/internal/models"
	sqlcgen "local-test/internal/sqlc/generated"
	"local-test/pkg/utils"

	"github.com/go-sql-driver/mysql"
)

func (r *Repository) CreateAccount(ctx context.Context, arg *models.CreateAccountParams) error {
    // Begin transaction
    tx, err := r.db.Begin()
    if err != nil {
        return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "begin transaction", Err: err})
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
		if err.(*mysql.MySQLError).Number == 1062 {
			return utils.WrapRepositoryError(&utils.ErrDuplicateEntry{Entity: "account id", Err: err})
		}
        return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "create account", Err: err})
    }

    // Create empty profile
    if err := q.CreateProfilesWithDefaultValues(ctx, arg.ID); err != nil {
        tx.Rollback()
		if err.(*mysql.MySQLError).Number == 1062 {
			return utils.WrapRepositoryError(&utils.ErrDuplicateEntry{Entity: "account id", Err: err})
		}
        return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "create profile", Err: err})
    }

    // Create empty setting
    if err := q.CreateSettingsWithDefaultValues(ctx, arg.ID); err != nil {
        tx.Rollback()
        return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "create setting", Err: err})
    }

	// Create empty interest
	if err := q.CreateInterestsWithDefaultValues(ctx, arg.ID); err != nil {
		tx.Rollback()
		return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "create interest", Err: err})
	}

    // Commit transaction
    if err := tx.Commit(); err != nil {
        tx.Rollback()
        return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "commit transaction", Err: err})
    }

    return nil
}


func (r *Repository) DeleteMyAccount(ctx context.Context, id string) (error) {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "begin transaction", Err: err})
	}

	// Create query object with transaction
	q := r.q.WithTx(tx)

	// Delete account
	res, err := q.DeleteAccount(ctx, id)
	if err != nil {
		tx.Rollback()
		return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "delete account", Err: err})
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "commit transaction", Err: err})
	}

	num, err := res.RowsAffected()
	if err != nil {
		return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "get rows affected", Err: err})
	}
	if num == 0 {
		return utils.WrapRepositoryError(&utils.ErrRecordNotFound{Condition: "account id"})
	}
	return nil
}