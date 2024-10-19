package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) CreateMessage(ctx context.Context, params *model.CreateMessageParams) error {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "begin transaction",
				Err:       err,
			},
		)
	}

	// Create query object with transaction
	q := r.q.WithTx(tx)

	// Create message
	messageID, err := q.CreateMessage(ctx, sqlcgen.CreateMessageParams{
		ConversationID:  params.ConversationID,
		SenderAccountID: params.SenderAccountID,
		Content:         params.Content,
	})
	if err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create message",
				Err:       err,
			},
		)
	}

	// Update last message id in conversation
	if err := q.UpdateLastMessage(ctx, sqlcgen.UpdateLastMessageParams{
		ID:            params.ConversationID,
		LastMessageID: sql.NullInt64{Int64: messageID, Valid: true},
	}); err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "update last message",
				Err:       err,
			},
		)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "commit transaction",
				Err:       err,
			},
		)
	}

	return nil
}

func (r *Repository) GetMessageList(ctx context.Context, params *model.GetMessageListParams) ([]*model.MessageResponse, error) {
	// Get messages
	messages, err := r.q.GetMessageList(ctx, sqlcgen.GetMessageListParams{
		ConversationID: params.ConversationID,
		Limit:          params.Limit,
		Offset:         params.Offset,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get messages",
				Err:       err,
			},
		)
	}

	// Convert to model
	var items []*model.MessageResponse
	for _, m := range messages {
		items = append(items, &model.MessageResponse{
			SenderAccountID: m.SenderAccountID,
			Content:         m.Content,
			IsRead:          m.IsRead,
			CreatedAt:       m.CreatedAt,
		})
	}

	return items, nil
}

func (r *Repository) MarkMessageListAsRead(ctx context.Context, params *model.MarkMessageListAsReadParams) error {
	// Mark messages as read
	if err := r.q.MarkMessageListAsRead(ctx, sqlcgen.MarkMessageListAsReadParams{
		ConversationID:  params.ConversationID,
		SenderAccountID: params.SenderAccountID,
	}); err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "mark messages as read",
				Err:       err,
			},
		)
	}

	return nil
}