package repository

import (
	"context"
	"database/sql"
	"errors"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"

	"github.com/lib/pq"
)

func (r *Repository) CreateAccount(ctx context.Context, arg *model.CreateAccountParams) error {
    // Begin transaction
    tx, err := r.db.Begin()
    if err != nil {
        return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "begin transaction",
				Err: err,
			},
		)
    }

    // Create query object with transaction
    q := r.q.WithTx(tx)

    // Create account
    query := sqlcgen.CreateAccountParams{
        ID:       arg.ID,
        UserID:   arg.UserID,
        UserName: arg.UserName,
    }
    if err := q.CreateAccount(ctx, query); err != nil {
        tx.Rollback()
		if err.(*pq.Error).Code == ErrCodeDuplicateEntry {
			return apperrors.WrapRepositoryError(
				&apperrors.ErrDuplicateEntry{
					Entity: "account id",
					Err: err,
				},
			)
		}

        return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create account",
				Err: err,
			},
		)
    }

    // Create empty profile
    if err := q.CreateProfilesWithDefaultValues(ctx, arg.ID); err != nil {
        tx.Rollback()
		if err.(*pq.Error).Code == ErrCodeDuplicateEntry {
			return apperrors.WrapRepositoryError(
				&apperrors.ErrDuplicateEntry{
					Entity: "account id",
					Err: err,
				},
			)
		}

        return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create profile",
				Err: err,
			},
		)
    }

    // Create empty setting
    if err := q.CreateSettingsWithDefaultValues(ctx, arg.ID); err != nil {
        tx.Rollback()
		if err.(*pq.Error).Code == ErrCodeDuplicateEntry {
			return apperrors.WrapRepositoryError(
				&apperrors.ErrDuplicateEntry{
					Entity: "account id",
					Err: err,
				},
			)
		}

        return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create setting",
				Err: err,
			},
		)
    }

	// Create empty interest
	if err := q.CreateInterestsWithDefaultValues(ctx, arg.ID); err != nil {
		tx.Rollback()
		if err.(*pq.Error).Code == ErrCodeDuplicateEntry {
			return apperrors.WrapRepositoryError(
				&apperrors.ErrDuplicateEntry{
					Entity: "account id",
					Err: err,
				},
			)
		}

		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create interest",
				Err: err,
			},
		)
	}

    // Commit transaction
    if err := tx.Commit(); err != nil {
        tx.Rollback()
        return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "commit transaction",
				Err: err,
			},
		)
    }

    return nil
}


func (r *Repository) DeleteMyAccount(ctx context.Context, accountID string) (error) {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "begin transaction",
				Err: err,
			},
		)
	}

	// Create query object with transaction
	q := r.q.WithTx(tx)

	// Delete account
	res, err := q.DeleteAccount(ctx, accountID)
	if err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "delete account",
				Err: err,
			},
		)
	}

    // Check if account is deleted
    num, err := res.RowsAffected()
    if err != nil {
        tx.Rollback()
        return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get rows affected",
				Err: err,
			},
		)
    }
    if num == 0 {
        tx.Rollback()
        return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "account id",
			},
		)
    }

	// Commit transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "commit transaction",
				Err: err,
			},
		)
	}

	return nil
}

func (r *Repository) GetAccountIDByUserID(ctx context.Context, userId string) (string, error) {
	AccountID, err := r.q.GetAccountIDByUserID(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", apperrors.WrapRepositoryError(
				&apperrors.ErrRecordNotFound{
					Condition: "user id",
				},
			)
		}

		return "", apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get account id by user id",
				Err: err,
			},
		)
	}

	return AccountID, nil
}

func (r *Repository) GetUserAndProfileInfoByAccountIDs(ctx context.Context, arg *model.GetUserAndProfileInfosParams) ([]*model.UserAndProfileInfo, error) {
	// Get user and profile info
	query := sqlcgen.GetUserAndProfileInfosParams{
		Limit:  int32(arg.Limit),
		Offset: int32(arg.Offset),
		Ids: arg.IDs,
	}
	res, err := r.q.GetUserAndProfileInfos(ctx, query)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get user and profile info by account ids",
				Err: err,
			},
		)
	}

	// Convert to model
	var userAndProfileInfos []*model.UserAndProfileInfo
	for _, r := range res {
		userAndProfileInfo := &model.UserAndProfileInfo{
			UserID: r.UserID,
			UserName: r.UserName,
			Bio: r.Bio.String,
			ProfileImageURL: r.ProfileImageUrl.String,
		}
		userAndProfileInfos = append(userAndProfileInfos, userAndProfileInfo)
	}

	return userAndProfileInfos, nil
}

func (r *Repository) IsAdmin(ctx context.Context, accountID string) (bool, error) {
	isAdmin, err := r.q.IsAdmin(ctx, accountID)
	if err != nil {
		return false, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get admin status",
				Err: err,
			},
		)
	}

	return isAdmin, nil
}

func (r *Repository) IsSuspended(ctx context.Context, accountID string) (bool, error) {
	isSuspended, err := r.q.IsSuspended(ctx, accountID)
	if err != nil {
		return false, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get suspended status",
				Err: err,
			},
		)
	}

	return isSuspended, nil
}