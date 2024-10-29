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

func (r *Repository) CreateReplyAndNotify(ctx context.Context, params *model.CreateReplyAndNotifyParams) (int64, error) {
	// Begin transaction
	tx, err := r.db.BeginTx(ctx, nil)
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
	sqlcParams, err := convertToCreateTweetAsReplyParams(params)
	if err != nil {
		tx.Rollback()
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "convert to create tweet as reply params",
				Err: err,
			},
		)
	}

	// Create tweet as reply
	tweetID, err := q.CreateTweetAsReply(ctx, *sqlcParams)
	if err != nil {
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create tweet as reply",
				Err: err,
			},
		)
	}

	// Create tweet hashtag
	if len(params.HashtagIDs) > 0 {
		if err := q.AssociateTweetWithHashtags(ctx, sqlcgen.AssociateTweetWithHashtagsParams{
			TweetID:    tweetID,
			HashtagIds: params.HashtagIDs,
		}); err != nil {
			tx.Rollback()
			return 0, apperrors.WrapRepositoryError(
				&apperrors.ErrOperationFailed{
					Operation: "associate tweet with hashtags",
					Err: err,
				},
			)
		}
	}

	// Insert reply
	if err := q.CreateReply(ctx, sqlcgen.CreateReplyParams{
		ReplyID:           tweetID,
		OriginalTweetID:   params.OriginalTweetID,
		ReplyingAccountID: params.ReplyingAccountID,
	}); err != nil {
		tx.Rollback()
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create reply",
				Err: err,
			},
		)
	}

	// Notify replied account
	if err := q.CreateNotification(ctx, sqlcgen.CreateNotificationParams{
		SenderAccountID: sql.NullString{String: params.ReplyingAccountID, Valid: true},
		RecipientAccountID: params.RepliedAccountID,
		Type: sqlcgen.NotificationTypeReply,
		TweetID: sql.NullInt64{Int64: tweetID, Valid: true},
	}); err != nil {
		tx.Rollback()
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create notification",
				Err: err,
			},
		)
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

func convertToCreateTweetAsReplyParams(params *model.CreateReplyAndNotifyParams) (*sqlcgen.CreateTweetAsReplyParams, error) {
	ret := &sqlcgen.CreateTweetAsReplyParams{
		AccountID: params.ReplyingAccountID,
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