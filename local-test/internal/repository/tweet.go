package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
	"time"

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

func (r *Repository) SetTweetAsPinned(ctx context.Context, params *model.SetTweetAsPinnedParams) error {
	// Pin tweet
	res, err := r.q.SetTweetAsPinned(ctx, sqlcgen.SetTweetAsPinnedParams{
		ID:         params.TweetID,
		AccountID: params.ClientAccountID,
	})
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "pin tweet",
				Err: err,
			},
		)
	}

	// Check if tweet pinned
	num, err := res.RowsAffected()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get rows affected",
				Err: err,
			},
		)
	}
	if num == 0 {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "tweet id",
			},
		)
	}

	return nil
}

func (r *Repository) UnsetTweetAsPinned(ctx context.Context, params *model.UnsetTweetAsPinnedParams) error {
	// Unpin tweet
	res, err := r.q.UnsetTweetAsPinned(ctx, sqlcgen.UnsetTweetAsPinnedParams{
		ID:         params.TweetID,
		AccountID: params.ClientAccountID,
	})
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "unpin tweet",
				Err: err,
			},
		)
	}

	// Check if tweet unpinned
	num, err := res.RowsAffected()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get rows affected",
				Err: err,
			},
		)
	}
	if num == 0 {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "tweet id",
			},
		)
	}

	return nil
}

func (r *Repository) GetPinnedTweetID(ctx context.Context, accountID string) (*int64, error) {
	// Get pinned tweet id
	tweetID, err := r.q.GetPinnedTweetID(ctx, accountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get pinned tweet id",
				Err: err,
			},
		)
	}

	return &tweetID, nil
}

func (r *Repository) GetAccountIDByTweetID(ctx context.Context, tweetID int64) (string, error) {
	// Get account id by tweet id
	accountID, err := r.q.GetAccountIDByTweetID(ctx, tweetID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", apperrors.WrapRepositoryError(
				&apperrors.ErrRecordNotFound{
					Condition: "tweet id",
				},
			)
		}

		return "", apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get account id by tweet id",
				Err: err,
			},
		)
	}

	return accountID, nil
}

func (r *Repository) GetRecentTweetMetadatas(ctx context.Context, params *model.GetRecentTweetMetadatasParams) ([]*model.TweetMetadata, error) {
	// Get recent tweet metadatas
	tweetMetadatas, err := r.q.GetRecentTweetMetadatas(ctx, sqlcgen.GetRecentTweetMetadatasParams{
		Limit: params.Limit,
		Offset: params.Offset,
		ClientAccountID: params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get recent tweet metadatas",
				Err: err,
			},
		)
	}

	// Convert to model
	var ret []*model.TweetMetadata
	for _, tweetMetadata := range tweetMetadatas {
		metadata := convertToTweetMetadata(tweetMetadata)
		ret = append(ret, metadata)
	}

	return ret, nil
}

func (r *Repository) GetTweetInfosByAccountID(ctx context.Context, params *model.GetTweetInfosByAccountIDParams) ([]*model.TweetInfoInternal, error) {
	// Get tweet infos by account id
	tweetInfos, err := r.q.GetTweetInfosByAccountID(ctx, sqlcgen.GetTweetInfosByAccountIDParams{
		ClientAccountID: params.ClientAccountID,
		TargetAccountID: params.TargetAccountID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get tweet infos by account id",
				Err: err,
			},
		)
	}

	// Convert to model
	ret, err := convertToTweetInfoInternal(tweetInfos)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "convert to tweet infos internal",
				Err: err,
			},
		)
	}

	return ret, nil
}

func (r *Repository) GetTweetInfosByIDs(ctx context.Context, params *model.GetTweetInfosByIDsParams) ([]*model.TweetInfoInternal, error) {
	// Get tweet infos by ids
	tweetInfos, err := r.q.GetTweetInfosByIDs(ctx, sqlcgen.GetTweetInfosByIDsParams{
		ClientAccountID: params.ClientAccountID,
		TweetIds:        params.TweetIDs,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get tweet infos by ids",
				Err: err,
			},
		)
	}

	// Convert to model
	ret, err := convertToTweetInfoInternal(tweetInfos)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "convert to tweet infos internal",
				Err: err,
			},
		)
	}

	return ret, nil
}

func (r *Repository) GetTweetCountByAccountID(ctx context.Context, accountID string) (int64, error) {
	// Get tweet count by account id
	count, err := r.q.GetTweetCountByAccountID(ctx, accountID)
	if err != nil {
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get tweet count by account id",
				Err: err,
			},
		)
	}

	return count, nil
}

func (r *Repository) SearchTweetsOrderByCreatedAt(ctx context.Context, params *model.SearchTweetsOrderByCreatedAtParams) ([]*model.TweetInfoInternal, error) {
	// Search tweets order by created at
	tweetInfos, err := r.q.SearchTweetsOrderByCreatedAt(ctx, sqlcgen.SearchTweetsOrderByCreatedAtParams{
		Keyword: params.Keyword,
		Offset:  params.Offset,
		Limit:   params.Limit,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "search tweets order by created at",
				Err: err,
			},
		)
	}

	// Convert to model
	ret, err := convertToTweetInfoInternal(tweetInfos)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "convert to tweet infos internal",
				Err: err,
			},
		)
	}

	return ret, nil
}

func (r *Repository) DeleteTweet(ctx context.Context, tweetID int64) error {
	// Delete tweet
	res, err := r.q.DeleteTweet(ctx, tweetID)
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "delete tweet",
				Err: err,
			},
		)
	}

	// Check if tweet deleted
	num, err := res.RowsAffected()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get rows affected",
				Err: err,
			},
		)
	}
	if num == 0 {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "tweet id",
			},
		)
	}

	return nil
}

func convertToCreateTweetParams(params *model.CreateTweetParams) (*sqlcgen.CreateTweetParams, error) {
	ret := &sqlcgen.CreateTweetParams{
		AccountID: params.AccountID,
	}

	if params.Content != nil {
		ret.Content = sql.NullString{String: *params.Content, Valid: true}
	}

	if params.Code != nil {
		jsonb, err := json.Marshal(params.Code)
		if err != nil {
			return nil, &apperrors.ErrInvalidInput{
				Message: "failed to marshal code",
			}
		}

		ret.Code = pqtype.NullRawMessage{
			RawMessage: jsonb,
			Valid: true,
		}
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

func convertToTweetMetadata(row sqlcgen.GetRecentTweetMetadatasRow) (*model.TweetMetadata) {
	metadata := model.TweetMetadata{
		TweetID:       row.ID,
		AccountID:     row.AccountID,
		LikesCount:    row.LikesCount,
		RetweetsCount: row.RetweetsCount,
		RepliesCount:  row.RepliesCount,
	}

	if row.Label1.Valid {
		label1 := model.Label(row.Label1.TweetLabel)
		metadata.Label1 = &label1
	}

	if row.Label2.Valid {
		label2 := model.Label(row.Label2.TweetLabel)
		metadata.Label2 = &label2
	}

	if row.Label3.Valid {
		label3 := model.Label(row.Label3.TweetLabel)
		metadata.Label3 = &label3
	}

	return &metadata
}

func mapRowToTweetInfoInternal(row interface{}) (*model.TweetInfoInternal, error) {
	var r struct {
		ID            int64
		AccountID     string
		LikesCount    int32
		RetweetsCount int32
		RepliesCount  int32
		IsQuote       bool
		IsReply       bool
		IsPinned      bool
		HasLiked      bool
		HasRetweeted  bool
		CreatedAt     time.Time
		Content       sql.NullString
		Code          pqtype.NullRawMessage
		Media         pqtype.NullRawMessage
	}

	switch t := row.(type) {
	case sqlcgen.GetTweetInfosByAccountIDRow:
		r.ID = t.ID
		r.AccountID = t.AccountID
		r.LikesCount = t.LikesCount
		r.RetweetsCount = t.RetweetsCount
		r.RepliesCount = t.RepliesCount
		r.IsQuote = t.IsQuote
		r.IsReply = t.IsReply
		r.IsPinned = t.IsPinned
		r.HasLiked = t.HasLiked
		r.HasRetweeted = t.HasRetweeted
		r.CreatedAt = t.CreatedAt
		r.Content = t.Content
		r.Code = t.Code
		r.Media = t.Media
	case sqlcgen.GetTweetInfosByIDsRow:
		r.ID = t.ID
		r.AccountID = t.AccountID
		r.LikesCount = t.LikesCount
		r.RetweetsCount = t.RetweetsCount
		r.RepliesCount = t.RepliesCount
		r.IsQuote = t.IsQuote
		r.IsReply = t.IsReply
		r.IsPinned = t.IsPinned
		r.HasLiked = t.HasLiked
		r.HasRetweeted = t.HasRetweeted
		r.CreatedAt = t.CreatedAt
		r.Content = t.Content
		r.Code = t.Code
		r.Media = t.Media
	case sqlcgen.SearchTweetsOrderByCreatedAtRow:
		r.ID = t.ID
		r.AccountID = t.AccountID
		r.LikesCount = t.LikesCount
		r.RetweetsCount = t.RetweetsCount
		r.RepliesCount = t.RepliesCount
		r.IsQuote = t.IsQuote
		r.IsReply = t.IsReply
		r.IsPinned = t.IsPinned
		r.HasLiked = t.HasLiked
		r.HasRetweeted = t.HasRetweeted
		r.CreatedAt = t.CreatedAt
		r.Content = t.Content
		r.Code = t.Code
		r.Media = t.Media
	default:
		return nil, fmt.Errorf("invalid type: %T", t)
	}

	info := &model.TweetInfoInternal{
		TweetID:       r.ID,
		AccountID:     r.AccountID,
		LikesCount:    r.LikesCount,
		RetweetsCount: r.RetweetsCount,
		RepliesCount:  r.RepliesCount,
		IsQuote:       r.IsQuote,
		IsReply:       r.IsReply,
		IsPinned:      r.IsPinned,
		HasLiked:      r.HasLiked,
		HasRetweeted:  r.HasRetweeted,
		CreatedAt:     r.CreatedAt,
	}

	if r.Content.Valid {
		info.Content = &r.Content.String
	}

	if r.Code.Valid {
		var code model.Code
		if err := json.Unmarshal(r.Code.RawMessage, &code); err != nil {
			return nil, errors.New("failed to unmarshal code")
		}
		info.Code = &code
	}

	if r.Media.Valid {
		var m model.Media
		if err := json.Unmarshal(r.Media.RawMessage, &m); err != nil {
			return nil, errors.New("failed to unmarshal media")
		}
		info.Media = &m
	}

	return info, nil
}

func convertToTweetInfoInternal(rows interface{}) ([]*model.TweetInfoInternal, error) {
	infos := []*model.TweetInfoInternal{}

	switch typedRows := rows.(type) {
	case []sqlcgen.GetTweetInfosByAccountIDRow:
		for _, r := range typedRows {
			info, err := mapRowToTweetInfoInternal(r)
			if err != nil {
				return nil, err
			}
			infos = append(infos, info)
		}
	case []sqlcgen.GetTweetInfosByIDsRow:
		for _, r := range typedRows {
			info, err := mapRowToTweetInfoInternal(r)
			if err != nil {
				return nil, err
			}
			infos = append(infos, info)
		}
	case []sqlcgen.SearchTweetsOrderByCreatedAtRow:
		for _, r := range typedRows {
			info, err := mapRowToTweetInfoInternal(r)
			if err != nil {
				return nil, err
			}
			infos = append(infos, info)
		}
	default:
		return nil, errors.New("unsupported row type")
	}

	return infos, nil
}

func convertToTweetInfoInternalFromGetTweetInfosByIDsRow(row sqlcgen.GetTweetInfosByIDsRow) (model.TweetInfoInternal, error) {
	info := model.TweetInfoInternal{
		TweetID:       row.ID,
		AccountID:     row.AccountID,
		LikesCount:    row.LikesCount,
		RetweetsCount: row.RetweetsCount,
		RepliesCount:  row.RepliesCount,
		IsQuote:       row.IsQuote,
		IsReply:       row.IsReply,
		IsPinned:      row.IsPinned,
		HasLiked:      row.HasLiked,
		HasRetweeted:  row.HasRetweeted,
		CreatedAt:     row.CreatedAt,
	}

	if row.Content.Valid {
		info.Content = &row.Content.String
	}

	if row.Code.Valid {
		var code model.Code
		if err := json.Unmarshal(row.Code.RawMessage, &code); err != nil {
			return model.TweetInfoInternal{}, &apperrors.ErrInvalidInput{
				Message: "failed to unmarshal code",
			}
		}
		info.Code = &code
	}

	if row.Media.Valid {
		var media model.Media
		if err := json.Unmarshal(row.Media.RawMessage, &media); err != nil {
			return model.TweetInfoInternal{}, &apperrors.ErrInvalidInput{
				Message: "failed to unmarshal media",
			}
		}
	}

	return info, nil
}