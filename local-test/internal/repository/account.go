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

func (r *Repository) CreateAccount(ctx context.Context, params *model.CreateAccountParams) error {
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
    if err := q.CreateAccount(ctx, sqlcgen.CreateAccountParams{
		ID:       params.ID,
		UserID:   params.UserID,
		UserName: params.UserName,
	}); err != nil {
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
    if err := q.CreateProfilesWithDefaultValues(ctx, params.ID); err != nil {
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
    if err := q.CreateSettingsWithDefaultValues(ctx, params.ID); err != nil {
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
	if err := q.CreateInterestsWithDefaultValues(ctx, params.ID); err != nil {
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
	// Delete account
	res, err := r.q.DeleteAccount(ctx, accountID)
	if err != nil {
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
        return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get rows affected",
				Err: err,
			},
		)
    }
    if num == 0 {
        return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "account id",
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

func (r *Repository) GetUserInfo(ctx context.Context, id string) (*model.UserInfoInternal, error) {
	// Get user info
	res, err := r.q.GetUserInfo(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.WrapRepositoryError(
				&apperrors.ErrRecordNotFound{
					Condition: "account id",
				},
			)
		}

		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get user info",
				Err: err,
			},
		)
	}

	// Convert to model
	userInfo := &model.UserInfoInternal{
		ID: res.ID,
		UserID: res.UserID,
		UserName: res.UserName,
		Bio: res.Bio.String,
		ProfileImageURL: res.ProfileImageUrl.String,
		BannerImageURL: res.BannerImageUrl.String,
		IsPrivate: res.IsPrivate.Bool,
		IsAdmin: res.IsAdmin,
		CreatedAt: res.CreatedAt,
	}

	return userInfo, nil
}

func (r *Repository) GetUserInfos(ctx context.Context, ids []string) ([]*model.UserInfoInternal, error) {
	// Get user infos
	res, err := r.q.GetUserInfos(ctx, ids)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get user and profile info by account ids",
				Err: err,
			},
		)
	}

	// Check if all accounts are found
	if len(res) != len(ids) {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "account id",
			},
		)
	}

	// Convert to model
	var userAndProfileInfos []*model.UserInfoInternal
	for _, r := range res {
		userAndProfileInfo := &model.UserInfoInternal{
			ID: r.ID,
			UserID: r.UserID,
			UserName: r.UserName,
			Bio: r.Bio.String,
			ProfileImageURL: r.ProfileImageUrl.String,
			BannerImageURL: r.BannerImageUrl.String,
			IsPrivate: r.IsPrivate.Bool,
			IsAdmin: r.IsAdmin,
			CreatedAt: r.CreatedAt,
		}
		userAndProfileInfos = append(userAndProfileInfos, userAndProfileInfo)
	}

	return userAndProfileInfos, nil
}

func (r *Repository) GetAccountInfo(ctx context.Context, accountID string) (*model.AccountInfo, error) {
	// Get account info
	res, err := r.q.GetAccountInfo(ctx, accountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.WrapRepositoryError(
				&apperrors.ErrRecordNotFound{
					Condition: "account id",
				},
			)
		}

		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get account info",
				Err: err,
			},
		)
	}

	// Convert to model
	accountInfo := &model.AccountInfo{
		IsAdmin : res.IsAdmin,
		IsSuspended : res.IsSuspended,
		IsPrivate : res.IsPrivate.Bool,
	}

	return accountInfo, nil
}