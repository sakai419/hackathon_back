package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"

	"github.com/lib/pq"
)

func (r *Repository) CreateLikeAndNotify(ctx context.Context, params *model.CreateLikeAndNotifyParams) error {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Create query object with transaction
	q := r.q.WithTx(tx)

	// Create like
	if err := q.CreateLike(ctx, sqlcgen.CreateLikeParams{
		LikingAccountID: params.LikingAccountID,
		OriginalTweetID: params.OriginalTweetID,
	}); err != nil {
		tx.Rollback()
		if err.(*pq.Error).Code == ErrCodeDuplicateEntry {
			return apperrors.WrapRepositoryError(
				&apperrors.ErrDuplicateEntry{
					Entity: "liking account id/original tweet id",
					Err: err,
				},
			)
		}

		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create like",
				Err: err,
			},
		)
	}

	// Notify original tweet poster
	if err := q.CreateNotification(ctx, sqlcgen.CreateNotificationParams{
		SenderAccountID: sql.NullString{String: params.LikingAccountID, Valid: true},
		RecipientAccountID: params.LikedAccountID,
		Type: sqlcgen.NotificationTypeLike,
		TweetID: sql.NullInt64{Int64: params.OriginalTweetID, Valid: true},
	}); err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create notification",
				Err: err,
			},
		)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "commit transaction",
				Err: err,
			},
		)
	}

	return nil
}

func (r *Repository) UnlikeTweet(ctx context.Context, params *model.UnlikeTweetParams) error {
	// Delete like
	res, err := r.q.DeleteLike(ctx, sqlcgen.DeleteLikeParams{
		LikingAccountID: params.LikingAccountID,
		OriginalTweetID: params.OriginalTweetID,
	})
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "delete like",
				Err: err,
			},
		)
	}

	// Check if like exists
	num, err := res.RowsAffected()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "check if like exists",
				Err: err,
			},
		)
	}
	if num == 0 {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "like",
			},
		)
	}

	return nil
}

func (r *Repository) GetLikingAccountIDs(ctx context.Context, params *model.GetLikingAccountIDsParams) ([]string, error) {
	// Get liking account ids
	ids, err := r.q.GetLikingAccountIDs(ctx, sqlcgen.GetLikingAccountIDsParams{
		OriginalTweetID: params.OriginalTweetID,
		Limit          : params.Limit,
		Offset         : params.Offset,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get liking account ids",
				Err: err,
			},
		)
	}

	return ids, nil
}