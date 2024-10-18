package repository

import (
	"context"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) UpdateSettings(ctx context.Context, arg *model.UpdateSettingsParams) error {
	// Update settings
	query := sqlcgen.UpdateSettingsParams{
		AccountID: arg.AccountID,
	}
	if arg.IsPrivate != nil {
		query.IsPrivate = *arg.IsPrivate
	}
	res, err := r.q.UpdateSettings(ctx, query)
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