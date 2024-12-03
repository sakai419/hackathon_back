package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"

	"github.com/lib/pq"
)

func (r *Repository) UpdateProfiles(ctx context.Context, params *model.UpdateProfilesParams) error {
	// begin transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "begin transaction",
				Err:       err,
			},
		)
	}

	// Create query object with transaction
	q := r.q.WithTx(tx)

	// Update profiles
	res, err := q.UpdateProfiles(ctx, convertToUpdateProfilesParams(params))
	if err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "update profiles",
				Err:       err,
			},
		)
	}

	// Check if profile is updated
	num, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get rows affected",
				Err:       err,
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

	// Update account infos
	res, err = q.UpdateAccountInfos(ctx, convertToUpdateAccountInfosParams(params))
	if err != nil {
		tx.Rollback()
		if err.(*pq.Error).Code == ErrCodeDuplicateEntry {
			return apperrors.WrapRepositoryError(
				&apperrors.ErrDuplicateEntry{
					Entity: "account id",
					Err:     err,
				},
			)
		}

		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "update account infos",
				Err:       err,
			},
		)
	}

	// Check if account info is updated
	num, err = res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get rows affected",
				Err:       err,
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
				Err:       err,
			},
		)
	}

	return nil
}

func convertToUpdateProfilesParams(params *model.UpdateProfilesParams) sqlcgen.UpdateProfilesParams {
	ret := sqlcgen.UpdateProfilesParams{
		AccountID: params.AccountID,
	}

	if params.Bio != nil {
		ret.Bio = sql.NullString{
			String: *params.Bio,
			Valid:  true,
		}
	}

	if params.ProfileImageURL != nil {
		ret.ProfileImageUrl = sql.NullString{
			String: *params.ProfileImageURL,
			Valid:  true,
		}
	}

	if params.BannerImageURL != nil {
		ret.BannerImageUrl = sql.NullString{
			String: *params.BannerImageURL,
			Valid:  true,
		}
	}

	return ret
}

func convertToUpdateAccountInfosParams(params *model.UpdateProfilesParams) sqlcgen.UpdateAccountInfosParams {
	ret := sqlcgen.UpdateAccountInfosParams{
		ID: params.AccountID,
	}

	if params.UserID != nil {
		ret.UserID = *params.UserID
	} else {
		ret.UserID = ""
	}

	if params.UserName != nil {
		ret.UserName = *params.UserName
	} else {
		ret.UserName = ""
	}

	return ret
}