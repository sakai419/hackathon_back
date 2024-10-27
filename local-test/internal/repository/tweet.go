package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"

	"github.com/sqlc-dev/pqtype"
)

func (r *Repository) CreateTweet(ctx context.Context, params *model.CreateTweetParams) (int64, error) {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "begin transaction",
				Err: err,
			},
		)
	}

	// Create query object with transaction
	q := r.q.WithTx(tx)

	// Convert params to sqlc params
	sqlcParams, err := convertToCreateTweetParams(params)
	if err != nil {
		tx.Rollback()
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "convert to create tweet params",
				Err: err,
			},
		)
	}

	// Create tweet
	tweetID, err := q.CreateTweet(ctx, *sqlcParams)
	if err != nil {
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create tweet",
				Err: err,
			},
		)
	}

	// Create tweet hashtag
	if len(params.HashtagIDs) > 0 {
		// Create tweet hashtag
		if err := q.AssociateTweetWithHashtags(ctx, sqlcgen.AssociateTweetWithHashtagsParams{
			TweetID:    tweetID,
			HashtagIds: params.HashtagIDs,
		}); err != nil {
			tx.Rollback()
			return 0, apperrors.WrapRepositoryError(
				&apperrors.ErrOperationFailed{
					Operation: "create tweet hashtag",
					Err: err,
				},
			)
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "commit transaction",
				Err: err,
			},
		)
	}

	return tweetID, nil
}

func convertToCreateTweetParams(params *model.CreateTweetParams) (*sqlcgen.CreateTweetParams, error) {
	ret := &sqlcgen.CreateTweetParams{
		AccountID: params.AccountID,
	}

	if params.Content != nil {
		ret.Content = sql.NullString{String: *params.Content, Valid: true}
	}

	if params.Code != nil {
		ret.Code = sql.NullString{String: *params.Code, Valid: true}
	}

	if params.Media != nil {
		jsonb, err := json.Marshal(params.Media)
		if err != nil {
			return nil, &apperrors.ErrInvalidInput{
				Message: "failed to marshal media",
			}
		}

		ret.Media = pqtype.NullRawMessage{
			RawMessage: jsonb,
			Valid: true,
		}
	}

	return ret, nil
}