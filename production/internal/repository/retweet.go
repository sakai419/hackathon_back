package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"

	"github.com/lib/pq"
)

func (r *Repository) CreateRetweetAndNotify(ctx context.Context, params *model.CreateRetweetAndNotifyParams) error {
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

	// Create retweet
	if err := q.CreateRetweet(ctx, sqlcgen.CreateRetweetParams{
		RetweetingAccountID: params.RetweetingAccountID,
		OriginalTweetID:     params.OriginalTweetID,
	}); err != nil {
		tx.Rollback()
		if err.(*pq.Error).Code == ErrCodeDuplicateEntry {
			return apperrors.WrapRepositoryError(
				&apperrors.ErrDuplicateEntry{
					Entity: "retweeting account id/original tweet id",
					Err: err,
				},
			)
		}

		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create retweet",
				Err: err,
			},
		)
	}

	if params.RetweetedAccountID != params.RetweetingAccountID {
		// Notify original tweet poster
		if err := q.CreateNotification(ctx, sqlcgen.CreateNotificationParams{
			SenderAccountID: sql.NullString{String: params.RetweetingAccountID, Valid: true},
			RecipientAccountID: params.RetweetedAccountID,
			Type: sqlcgen.NotificationTypeRetweet,
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

func (r *Repository) Unretweet(ctx context.Context, params *model.UnretweetParams) error {
	// Delete retweet
	res, err := r.q.DeleteRetweet(ctx, sqlcgen.DeleteRetweetParams{
		RetweetingAccountID: params.RetweetingAccountID,
		OriginalTweetID: params.OriginalTweetID,
	})
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "delete retweet",
				Err: err,
			},
		)
	}

	// Check if retweet exists
	num, err := res.RowsAffected()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "check if retweet exists",
				Err: err,
			},
		)
	}
	if num == 0 {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "retweet",
			},
		)
	}

	return nil
}

func (r *Repository) GetRetweetingAccountIDs(ctx context.Context, params *model.GetRetweetingAccountIDsParams) ([]string, error) {
	// Get retweeting accounts
	retweetingAccounts, err := r.q.GetRetweetingAccountIDs(ctx, sqlcgen.GetRetweetingAccountIDsParams{
		OriginalTweetID: params.OriginalTweetID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get retweeting accounts",
				Err: err,
			},
		)
	}

	return retweetingAccounts, nil
}

func (r *Repository) GetRetweetedTweetIDsByAccountID(ctx context.Context, params *model.GetRetweetedTweetIDsByAccountIDParams) ([]int64, error) {
	// Get retweeted tweet ids
	retweetedTweetIDs, err := r.q.GetRetweetedTweetIDsByAccountID(ctx, sqlcgen.GetRetweetedTweetIDsByAccountIDParams{
		RetweetingAccountID: params.RetweetingAccountID,
		Limit:               params.Limit,
		Offset:              params.Offset,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get retweeted tweet ids",
				Err: err,
			},
		)
	}

	return retweetedTweetIDs, nil
}