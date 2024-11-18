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

func (r *Repository) CreateQuoteAndNotify(ctx context.Context, params *model.CreateQuoteAndNotifyParams) (int64, error) {
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
	sqlcParams, err := convertToCreateTweetAsQuoteParams(params)
	if err != nil {
		tx.Rollback()
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "convert to create tweet as quote params",
				Err: err,
			},
		)
	}

	// Create tweet as quote
	tweetID, err := q.CreateTweetAsQuote(ctx, *sqlcParams)
	if err != nil {
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create tweet as quote",
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

	// Insert quote
	if err := q.CreateQuote(ctx, sqlcgen.CreateQuoteParams{
		QuotingAccountID: params.QuotingAccountID,
		OriginalTweetID:  params.OriginalTweetID,
		QuoteID:          tweetID,
	}); err != nil {
		tx.Rollback()
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create quote",
				Err: err,
			},
		)
	}

	if params.QuotingAccountID != params.QuotedAccountID {
		// Notify quoted account
		if err := q.CreateNotification(ctx, sqlcgen.CreateNotificationParams{
			SenderAccountID: sql.NullString{String: params.QuotingAccountID, Valid: true},
			RecipientAccountID: params.QuotedAccountID,
			Type: sqlcgen.NotificationTypeQuote,
			TweetID: sql.NullInt64{Int64: params.OriginalTweetID, Valid: true},
		}); err != nil {
			tx.Rollback()
			return 0, apperrors.WrapRepositoryError(
				&apperrors.ErrOperationFailed{
					Operation: "create notification",
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

func (r *Repository) GetQuotingAccountIDs(ctx context.Context, params *model.GetQuotingAccountIDsParams) ([]string, error) {
	// Get quoting account ids
	accountIDs, err := r.q.GetQuotingAccountIDs(ctx, sqlcgen.GetQuotingAccountIDsParams{
		OriginalTweetID: params.OriginalTweetID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get quoting account ids",
				Err: err,
			},
		)
	}

	return accountIDs, nil
}

func (r *Repository) GetQuotedTweetInfos(ctx context.Context, params *model.GetQuotedTweetInfosParams) ([]*model.QuotedTweetInfoInternal, error) {
	// Get quote relations
	quoteRelations, err := r.q.GetQuoteRelations(ctx, params.QuotingTweetIDs)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get quoted tweet ids",
				Err: err,
			},
		)
	}

	// Extract quoted tweet ids
	quotedTweetIDs := make([]int64, 0, len(quoteRelations))
	for _, relation := range quoteRelations {
		quotedTweetIDs = append(quotedTweetIDs, relation.OriginalTweetID)
	}

	// Get quoted tweets originals
	originalTweets, err := r.q.GetTweetInfosByIDs(ctx, sqlcgen.GetTweetInfosByIDsParams{
		TweetIds:		 quotedTweetIDs,
		ClientAccountID: params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get quoted tweet originals",
				Err: err,
			},
		)
	}

	// Check if all quoted tweets are found
	if len(quotedTweetIDs) != len(originalTweets) {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "tweet ids",
			},
		)
	}

	// Convert to quoted tweet infos
	ret, err := convertToQuotedTweetInfos(quoteRelations, originalTweets)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "convert to quoted tweet infos",
				Err: err,
			},
		)
	}

	return ret, nil
}

func convertToCreateTweetAsQuoteParams(params *model.CreateQuoteAndNotifyParams) (*sqlcgen.CreateTweetAsQuoteParams, error) {
	ret := &sqlcgen.CreateTweetAsQuoteParams{
		AccountID: params.QuotingAccountID,
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

func convertToQuotedTweetInfos(quoteRelations []sqlcgen.GetQuoteRelationsRow, originalTweets []sqlcgen.GetTweetInfosByIDsRow) ([]*model.QuotedTweetInfoInternal, error) {
	// Create map for original tweets
	originalTweetMap := make(map[int64]sqlcgen.GetTweetInfosByIDsRow, len(originalTweets))
	for _, tweet := range originalTweets {
		originalTweetMap[tweet.ID] = tweet
	}

	// Create quoted tweet infos
	ret := make([]*model.QuotedTweetInfoInternal, 0, len(quoteRelations))
	for _, relation := range quoteRelations {
		originalTweet, ok := originalTweetMap[relation.OriginalTweetID]
		if !ok {
			return nil, &apperrors.ErrRecordNotFound{
				Condition: "original tweet",
			}
		}

		tweet := model.TweetInfoInternal{
			TweetID:       originalTweet.ID,
			AccountID:     originalTweet.AccountID,
			LikesCount:    originalTweet.LikesCount,
			RetweetsCount: originalTweet.RetweetsCount,
			RepliesCount:  originalTweet.RepliesCount,
			IsQuote:       originalTweet.IsQuote,
			IsReply: 	   originalTweet.IsReply,
			IsPinned: 	   originalTweet.IsPinned,
			HasLiked:      originalTweet.HasLiked,
			HasRetweeted:  originalTweet.HasRetweeted,
			CreatedAt:     originalTweet.CreatedAt,
		}

		if originalTweet.Content.Valid {
			tweet.Content = &originalTweet.Content.String
		}

		if originalTweet.Code.Valid {
			tweet.Code = &originalTweet.Code.String
		}

		if originalTweet.Media.Valid {
			var media model.Media
			if err := json.Unmarshal(originalTweet.Media.RawMessage, &media); err != nil {
				return nil, &apperrors.ErrInvalidInput{
					Message: "failed to unmarshal media",
				}
			}
			tweet.Media = &media
		}

		ret = append(ret, &model.QuotedTweetInfoInternal{
			QuotedTweet: tweet,
			QuotingTweetID: relation.QuoteID,
		})
	}

	return ret, nil
}