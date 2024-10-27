package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"

	"github.com/lib/pq"
)

func (r *Repository) PostRetweet(ctx context.Context, params *model.PostRetweetParams) error {
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
	tweetID, err := q.CreateTweetAsRetweet(ctx, sqlcgen.CreateTweetAsRetweetParams{
		AccountID:       params.AccountID,
		OriginalTweetID: sql.NullInt64{Int64: params.OriginalTweetID, Valid: true},
	})
	if err != nil {
		tx.Rollback()
		if err.(*pq.Error).Code == ErrCodeForeignKey {
			return apperrors.WrapRepositoryError(
				&apperrors.ErrRecordNotFound{
					Condition: "original tweet id",
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

	// Insert into retweets table
	if err := q.CreateRetweetOrQuote(ctx, sqlcgen.CreateRetweetOrQuoteParams{
		RetweetID:             tweetID,
		RetweetingAccountID: params.AccountID,
		OriginalTweetID:     params.OriginalTweetID,
	}); err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create retweet or quote",
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