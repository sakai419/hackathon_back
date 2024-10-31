package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
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
	var ret []*model.TweetInfoInternal
	for _, tweetInfo := range tweetInfos {
		info, err := convertToTweetInfoInternalFromGetTweetInfosByIDsRow(tweetInfo)
		if err != nil {
			return nil, apperrors.WrapRepositoryError(
				&apperrors.ErrOperationFailed{
					Operation: "convert to tweet infos internal",
					Err: err,
				},
			)
		}

		ret = append(ret, &info)
	}

	return ret, nil
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

func convertToTweetInfoInternal(row []sqlcgen.GetTweetInfosByAccountIDRow) ([]*model.TweetInfoInternal, error) {
	infos := make([]*model.TweetInfoInternal, 0, len(row))
	for _, r := range row {
		info := model.TweetInfoInternal{
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
			info.Code = &r.Code.String
		}

		if r.Media.Valid {
			var media model.Media
			if err := json.Unmarshal(r.Media.RawMessage, &media); err != nil {
				return nil, &apperrors.ErrInvalidInput{
					Message: "failed to unmarshal media",
				}
			}
		}

		infos = append(infos, &info)
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
		info.Code = &row.Code.String
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