package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) UpdateProfiles(ctx context.Context, params *model.UpdateProfilesParams) error {
	// Update profiles
	res, err := r.q.UpdateProfiles(ctx, convertToUpdateProfilesParams(params))
	if err != nil {
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