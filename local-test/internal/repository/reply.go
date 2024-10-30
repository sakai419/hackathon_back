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

func (r *Repository) GetRepliedTweetInfos(ctx context.Context, params *model.GetRepliedTweetInfosParams) ([]*model.RepliedTweetInfoInternal, error) {
	// Get reply relations
	replyRelations, err := r.q.GetReplyRelations(ctx, params.ReplyingTweetIDs)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get reply relations",
				Err: err,
			},
		)
	}

	// extract original and parent reply tweet ids
	originalTweetIDs := make([]int64, 0, len(replyRelations))
	parentReplyTweetIDs := make([]int64, 0, len(replyRelations))
	for _, relation := range replyRelations {
		originalTweetIDs = append(originalTweetIDs, relation.OriginalTweetID)
		if relation.ParentReplyID.Valid {
			parentReplyTweetIDs = append(parentReplyTweetIDs, relation.ParentReplyID.Int64)
		}
	}

	// Get original tweet infos
	originalTweetInfos, err := r.q.GetTweetInfosByIDs(ctx, sqlcgen.GetTweetInfosByIDsParams{
		TweetIds: originalTweetIDs,
		ClientAccountID: params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get original tweet infos",
				Err: err,
			},
		)
	}

	// Get parent reply tweet infos
	parentReplyTweetInfos, err := r.q.GetTweetInfosByIDs(ctx, sqlcgen.GetTweetInfosByIDsParams{
		TweetIds: parentReplyTweetIDs,
		ClientAccountID: params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get parent reply tweet infos",
				Err: err,
			},
		)
	}

	// Check if all tweet infos are found
	if len(originalTweetIDs) != len(originalTweetInfos) || len(parentReplyTweetIDs) != len(parentReplyTweetInfos) {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "tweet ids",
			},
		)
	}

	// Convert to replied tweet infos
	ret, err := r.convertToRepliedTweetInfos(replyRelations, originalTweetInfos, parentReplyTweetInfos)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "convert to replied tweet infos",
				Err: err,
			},
		)
	}

	return ret, nil
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

func (r *Repository) convertToRepliedTweetInfos(replyRelations []sqlcgen.GetReplyRelationsRow, originalTweetInfos, parentReplyTweetInfos []sqlcgen.GetTweetInfosByIDsRow) ([]*model.RepliedTweetInfoInternal, error) {
	// Create map for original tweet infos
	originalTweetInfoMap := make(map[int64]sqlcgen.GetTweetInfosByIDsRow, len(originalTweetInfos))
	for _, info := range originalTweetInfos {
		originalTweetInfoMap[info.ID] = info
	}

	// Create map for parent reply tweet infos
	parentReplyTweetInfoMap := make(map[int64]sqlcgen.GetTweetInfosByIDsRow, len(parentReplyTweetInfos))
	for _, info := range parentReplyTweetInfos {
		parentReplyTweetInfoMap[info.ID] = info
	}

	// Create replied tweet infos
	ret := make([]*model.RepliedTweetInfoInternal, 0, len(replyRelations))
	for _, relation := range replyRelations {
		originalTweet, ok := originalTweetInfoMap[relation.OriginalTweetID]
		if !ok {
			return nil, &apperrors.ErrRecordNotFound{
				Condition: "original tweet info",
			}
		}

		var parentReplyTweet *sqlcgen.GetTweetInfosByIDsRow
		if relation.ParentReplyID.Valid {
			info := parentReplyTweetInfoMap[relation.ParentReplyID.Int64]
			parentReplyTweet = &info
			if !ok {
				return nil, &apperrors.ErrRecordNotFound{
					Condition: "parent reply tweet info",
				}
			}
		}

		originalTweetInfo, err := convertToTweetInfoInternalFromGetTweetInfosByIDsRow(originalTweet)
		if err != nil {
			return nil, &apperrors.ErrOperationFailed{
				Operation: "convert to tweet info internal",
			}
		}

		var parentReplyTweetInfo *model.TweetInfoInternal
		var omittedReplyExist *bool
		if parentReplyTweet != nil {
			info, err := convertToTweetInfoInternalFromGetTweetInfosByIDsRow(*parentReplyTweet)
			if err != nil {
				return nil, &apperrors.ErrOperationFailed{
					Operation: "convert to tweet info internal",
				}
			}
			parentReplyTweetInfo = &info
			exist, err := r.q.CheckParentReplyExist(context.Background(), relation.ParentReplyID.Int64)
			if err != nil {
				return nil, apperrors.WrapRepositoryError(
					&apperrors.ErrOperationFailed{
						Operation: "check parent reply exist",
						Err: err,
					},
				)
			}
			omittedReplyExist = &exist
		}

		ret = append(ret, &model.RepliedTweetInfoInternal{
			OriginalTweet:     originalTweetInfo,
			ParentReplyTweet:  parentReplyTweetInfo,
			ReplyingTweetID:   relation.ReplyID,
			OmittedReplyExist: omittedReplyExist,
		})
	}

	return ret, nil
}