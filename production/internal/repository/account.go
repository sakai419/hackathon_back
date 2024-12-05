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

func (r *Repository) GetUserInfo(ctx context.Context, params *model.GetUserInfoParams) (*model.UserInfoInternal, error) {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "begin transaction",
				Err: err,
			},
		)
	}

	// Create query object with transaction
	q := r.q.WithTx(tx)

	// Get user info
	res, err := q.GetUserInfo(ctx, params.TargetAccountID)
	if err != nil {
		tx.Rollback()
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

	// Get follow status
	followStatus, err := q.CheckFollowStatus(ctx, sqlcgen.CheckFollowStatusParams{
		ClientAccountID: params.ClientAccountID,
		TargetAccountID: params.TargetAccountID,
	})
	if err != nil {
		tx.Rollback()
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "check follow status",
				Err: err,
			},
		)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "commit transaction",
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
		IsFollowing: followStatus.IsFollowing,
		IsFollowed: followStatus.IsFollowed,
		IsPending: followStatus.IsPending,
		CreatedAt: res.CreatedAt,
	}

	return userInfo, nil
}

func (r *Repository) GetUserInfos(ctx context.Context, params *model.GetUserInfosParams) ([]*model.UserInfoInternal, error) {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "begin transaction",
				Err: err,
			},
		)
	}

	// Create query object with transaction
	q := r.q.WithTx(tx)

	// Get user infos
	res, err := q.GetUserInfos(ctx, params.TargetAccountIDs)
	if err != nil {
		tx.Rollback()
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get user and profile info by account ids",
				Err: err,
			},
		)
	}

	// Check if all accounts are found
	if len(res) != len(params.TargetAccountIDs) {
		tx.Rollback()
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "account id",
			},
		)
	}

	// Get follow status
	followStatuses, err := q.CheckMultipleFollowStatus(ctx, sqlcgen.CheckMultipleFollowStatusParams{
		ClientAccountID: params.ClientAccountID,
		AccountIds: params.TargetAccountIDs,
	})
	if err != nil {
		tx.Rollback()
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "check multiple follow status",
				Err: err,
			},
		)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "commit transaction",
				Err: err,
			},
		)
	}

	// Convert to map
	followStatusMap := make(map[string]sqlcgen.CheckMultipleFollowStatusRow)
	for _, f := range followStatuses {
		followStatusMap[f.AccountID.(string)] = f
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
		userAndProfileInfo.IsFollowing = followStatusMap[r.ID].IsFollowing
		userAndProfileInfo.IsFollowed = followStatusMap[r.ID].IsFollowed
		userAndProfileInfo.IsPending = followStatusMap[r.ID].IsPending
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

func (r *Repository) SearchUsersOrderByCreatedAt(ctx context.Context, params *model.SearchUsersOrderByCreatedAtParams) ([]*model.UserInfoInternal, error) {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "begin transaction",
				Err: err,
			},
		)
	}

	// Create query object with transaction
	q := r.q.WithTx(tx)

	// Search users
	users, err := q.SearchAccountsOrderByCreatedAt(ctx, sqlcgen.SearchAccountsOrderByCreatedAtParams{
		Keyword: params.Keyword,
		Offset: params.Offset,
		Limit: params.Limit,
	})
	if err != nil {
		tx.Rollback()
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "search users",
				Err: err,
			},
		)
	}

	// extract accountIDs
	accountIDs := make([]string, 0)
	for _, u := range users {
		accountIDs = append(accountIDs, u.ID)
	}

	// Get follow status
	followStatuses, err := q.CheckMultipleFollowStatus(ctx, sqlcgen.CheckMultipleFollowStatusParams{
		ClientAccountID: params.ClientAccountID,
		AccountIds: accountIDs,
	})
	if err != nil {
		tx.Rollback()
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "check multiple follow status",
				Err: err,
			},
		)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "commit transaction",
				Err: err,
			},
		)
	}

	// Convert to map
	followStatusMap := make(map[string]sqlcgen.CheckMultipleFollowStatusRow)
	for _, f := range followStatuses {
		followStatusMap[f.AccountID.(string)] = f
	}

	// Convert to model
	var userInfos []*model.UserInfoInternal
	for _, u := range users {
		userInfo := &model.UserInfoInternal{
			ID: u.ID,
			UserID: u.UserID,
			UserName: u.UserName,
			Bio: u.Bio.String,
			ProfileImageURL: u.ProfileImageUrl.String,
			IsPrivate: u.IsPrivate.Bool,
			IsAdmin: u.IsAdmin,
			CreatedAt: u.CreatedAt,
		}
		userInfo.IsFollowing = followStatusMap[u.ID].IsFollowing
		userInfo.IsFollowed = followStatusMap[u.ID].IsFollowed
		userInfo.IsPending = followStatusMap[u.ID].IsPending
		userInfos = append(userInfos, userInfo)
	}

	return userInfos, nil
}

func (r *Repository) FilterAccessibleAccountIDs(ctx context.Context, params *model.FilterAccessibleAccountIDsParams) ([]string, error) {
	// Filter accessible account ids
	accessibleAccountIDs, err := r.q.FilterAccessibleAccountIDs(ctx, sqlcgen.FilterAccessibleAccountIDsParams{
		AccountIds: params.AccountIDs,
		ClientAccountID: params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "filter accessible account ids",
				Err: err,
			},
		)
	}

	return accessibleAccountIDs, nil
}

func (r *Repository) FilterAccessibleAccountIDsByBlockStatus(ctx context.Context, params *model.FilterAccessibleAccountIDsByBlockStatusParams) ([]string, error) {
	// Filter accessible account ids
	accessibleAccountIDs, err := r.q.FilterAccessibleAccountIDsByBlockStatus(ctx, sqlcgen.FilterAccessibleAccountIDsByBlockStatusParams{
		AccountIds: params.AccountIDs,
		ClientAccountID: params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "filter accessible account ids by block status",
				Err: err,
			},
		)
	}

	return accessibleAccountIDs, nil
}