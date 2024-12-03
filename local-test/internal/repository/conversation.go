package repository

import (
	"context"
	"database/sql"
	"errors"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) GetConversationID(ctx context.Context, params *model.GetConversationIDParams) (int64, error) {
	// Get conversation id
	conversationID, err := r.q.GetConversationID(ctx, sqlcgen.GetConversationIDParams{
		Account1ID: params.Account1ID,
		Account2ID: params.Account2ID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, apperrors.WrapRepositoryError(
				&apperrors.ErrRecordNotFound{
					Condition: "conversation id",
				},
			)
		}

		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get conversation id",
				Err: err,
			},
		)
	}

	return conversationID, nil
}

func (r *Repository) GetConversationList(ctx context.Context, params *model.GetConversationListParams) ([]*model.Conversation, error) {
	// Get conversations
	conversations, err := r.q.GetConversationList(ctx, sqlcgen.GetConversationListParams{
		Account1ID: params.AccountID,
		Limit:      params.Limit,
		Offset:     params.Offset,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get conversations",
				Err: err,
			},
		)
	}

	// Convert to model
	conversationModels := convertToConversationModel(conversations, params.AccountID)

	return conversationModels, nil
}

func (r *Repository) GetUnreadConversationCount(ctx context.Context, AccountID string) (int64, error) {
	count, err := r.q.GetUnreadConversationCount(ctx, AccountID)
	if err != nil {
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get unread conversation count",
				Err: err,
			},
		)
	}

	return count, nil
}

func convertToConversationModel(conversations []sqlcgen.GetConversationListRow, AccountID string) []*model.Conversation {
	var conversationModels []*model.Conversation
	for _, c := range conversations {
		conversationModels = append(conversationModels, &model.Conversation{
			ID:              c.ID,
			OpponentID: func() string {
				if c.Account1ID == AccountID {
					return c.Account2ID
				}
				return c.Account1ID
			}(),
			LastMessageTime: c.LastMessageTime,
			Content:         c.Content.String,
			SenderUserID:    c.SenderUserID.String,
			IsRead:          c.IsRead.Bool,
		})
	}
	return conversationModels
}
