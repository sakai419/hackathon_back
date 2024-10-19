package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) CreateMessage(ctx context.Context, arg *model.CreateMessageParams) error {
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
	query := sqlcgen.CreateMessageParams{
		ConversationID:  arg.ConversationID,
		SenderAccountID: arg.SenderAccountID,
		Content:         arg.Content,
	}
	messageID, err := q.CreateMessage(ctx, query)
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
	updateQuery := sqlcgen.UpdateLastMessageParams{
		ID: arg.ConversationID,
		LastMessageID:  sql.NullInt64{Int64: messageID, Valid: true},
	}
	if err := q.UpdateLastMessage(ctx, updateQuery); err != nil {
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

func (r *Repository) GetMessageList(ctx context.Context, arg *model.GetMessageListParams) ([]*model.MessageResponse, error) {
	// Get messages
	query := sqlcgen.GetMessageListParams{
		ConversationID: arg.ConversationID,
		Limit:          arg.Limit,
		Offset:         arg.Offset,
	}
	messages, err := r.q.GetMessageList(ctx, query)
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

func (r *Repository) MarkMessageListAsRead(ctx context.Context, arg *model.MarkMessageListAsReadParams) error {
	// Mark messages as read
	query := sqlcgen.MarkMessageListAsReadParams{
		ConversationID:  arg.ConversationID,
		SenderAccountID: arg.SenderAccountID,
	}
	if err := r.q.MarkMessageListAsRead(ctx, query); err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "mark messages as read",
				Err:       err,
			},
		)
	}

	return nil
}