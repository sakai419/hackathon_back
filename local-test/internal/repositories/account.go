package repositories

import (
	"context"
	"database/sql"
	"errors"
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
    params := sqlcgen.CreateAccountParams{
        ID:       arg.ID,
        UserID:   arg.UserID,
        UserName: arg.UserName,
    }
    if err := q.CreateAccount(ctx, params); err != nil {
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
		if err.(*mysql.MySQLError).Number == 1062 {
			return utils.WrapRepositoryError(&utils.ErrDuplicateEntry{Entity: "account id", Err: err})
		}
        return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "create setting", Err: err})
    }

	// Create empty interest
	if err := q.CreateInterestsWithDefaultValues(ctx, arg.ID); err != nil {
		tx.Rollback()
		if err.(*mysql.MySQLError).Number == 1062 {
			return utils.WrapRepositoryError(&utils.ErrDuplicateEntry{Entity: "account id", Err: err})
		}
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

    // Check if account is deleted
    num, err := res.RowsAffected()
    if err != nil {
        tx.Rollback()
        return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "get rows affected", Err: err})
    }
    if num == 0 {
        tx.Rollback()
        return utils.WrapRepositoryError(&utils.ErrRecordNotFound{Condition: "account id"})
    }

	// Commit transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "commit transaction", Err: err})
	}

	return nil
}

func (r *Repository) GetAccountIDByUserId(ctx context.Context, userId string) (string, error) {
	AccountID, err := r.q.GetAccountIDByUserId(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", utils.WrapRepositoryError(&utils.ErrRecordNotFound{Condition: "user id"})
		}
		return "", utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "get account id by user id", Err: err})
	}

	return AccountID, nil
}