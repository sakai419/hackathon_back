package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"net/http"
)

func (s *Service) SendMessage(ctx context.Context, params *model.SendMessageParams) error {
	// Validate input
	if err := params.Validate(); err != nil {
		return &apperrors.AppError{
			Status:   http.StatusBadRequest,
			Code:    "INVALID_INPUT",
			Message: "Invalid input",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate input",
					Err:       err,
				},
			),
		}
	}

	// Get conversation id
	conversationID, err := s.repo.GetConversationID(ctx, &model.GetConversationIDParams{
		Account1ID:    params.ClientAccountID,
		Account2ID:    params.TargetAccountID,
	})
	if err != nil {
		return conversationErrHandler(err)
	}

	// Send message
	if err := s.repo.CreateMessage(ctx, &model.CreateMessageParams{
		ConversationID:  conversationID,
		Content:         params.Content,
		SenderAccountID: params.ClientAccountID,
	}); err != nil {
		return &apperrors.AppError{
			Status:   http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to send message",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "create message",
					Err:       err,
				},
			),
		}
	}

	return nil
}

func (s *Service) GetMessages(ctx context.Context, params *model.GetMessagesParams) ([]*model.MessageResponse, error) {
	// Validate input
	if err := params.Validate(); err != nil {
		return nil, &apperrors.AppError{
			Status:   http.StatusBadRequest,
			Code:    "INVALID_INPUT",
			Message: "Invalid input",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate input",
					Err:       err,
				},
			),
		}
	}

	// Get conversation id
	conversationID, err := s.repo.GetConversationID(ctx, &model.GetConversationIDParams{
		Account1ID:    params.ClientAccountID,
		Account2ID:    params.TargetAccountID,
	})
	if err != nil {
		return nil, conversationErrHandler(err)
	}

	// Get messages
	messages, err := s.repo.GetMessageList(ctx, &model.GetMessageListParams{
		ConversationID: conversationID,
		Limit:          params.Limit,
		Offset:         params.Offset,
	})
	if err != nil {
		return nil, &apperrors.AppError{
			Status:   http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to get messages",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "get messages",
					Err:       err,
				},
			),
		}
	}

	return messages, nil
}

func (s *Service) MarkMessagesAsRead(ctx context.Context, params *model.MarkMessagesAsReadParams) error {
	// Validate input
	if err := params.Validate(); err != nil {
		return &apperrors.AppError{
			Status:   http.StatusBadRequest,
			Code:    "INVALID_INPUT",
			Message: "Invalid input",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate input",
					Err:       err,
				},
			),
		}
	}

	// Get conversation id
	conversationID, err := s.repo.GetConversationID(ctx, &model.GetConversationIDParams{
		Account1ID:    params.ClientAccountID,
		Account2ID:    params.TargetAccountID,
	})
	if err != nil {
		return conversationErrHandler(err)
	}

	// Mark messages as read
	if err := s.repo.MarkMessageListAsRead(ctx, &model.MarkMessageListAsReadParams{
		ConversationID:  conversationID,
		SenderAccountID: params.TargetAccountID,
	}); err != nil {
		return &apperrors.AppError{
			Status:   http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to mark messages as read",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "mark messages as read",
					Err:       err,
				},
			),
		}
	}

	return nil
}

func conversationErrHandler(err error) error {
	var notFoundErr *apperrors.ErrRecordNotFound
	if errors.As(err, &notFoundErr) {
		return &apperrors.AppError{
			Status:  http.StatusNotFound,
			Code:    "CONVERSATION_NOT_FOUND",
			Message: "Conversation not found",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "get conversation id",
					Err:       notFoundErr,
				},
			),
		}
	}

	return &apperrors.AppError{
		Status:   http.StatusInternalServerError,
		Code:    "DATABASE_ERROR",
		Message: "Failed to get conversation id",
		Err:     apperrors.WrapServiceError(
			&apperrors.ErrOperationFailed{
				Operation: "get conversation id",
				Err:       err,
			},
		),
	}
}