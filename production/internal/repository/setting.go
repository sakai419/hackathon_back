package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) UpdateSettings(ctx context.Context, params *model.UpdateSettingsParams) error {
	// Update settings
	res, err := r.q.UpdateSettings(ctx, convertToUpdateSettingsParams(params))
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "update settings",
				Err:       err,
			},
		)
	}

	// Check if setting is updated
	num, err := res.RowsAffected()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get rows affected",
				Err:       err,
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

func convertToUpdateSettingsParams(params *model.UpdateSettingsParams) sqlcgen.UpdateSettingsParams {
	ret := sqlcgen.UpdateSettingsParams{
		AccountID: params.AccountID,
	}

	if params.IsPrivate != nil {
		ret.IsPrivate = sql.NullBool{
			Bool:  *params.IsPrivate,
			Valid: true,
		}
	}

	return ret
}