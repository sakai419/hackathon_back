package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"net/http"
)

func (s *Service) SendMessage(ctx context.Context, params *model.SendMessageParams) error {
	// Validate input
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

	// Get conversation id
	conversationID, err := s.repo.GetConversationID(ctx, &model.GetConversationIDParams{
		Account1ID:    params.ClientAccountID,
		Account2ID:    params.TargetAccountID,
	})
	if err != nil {
		return apperrors.NewNotFoundAppError("conversation", "get conversation id", err)
	}

	// Send message
	if err := s.repo.CreateMessage(ctx, &model.CreateMessageParams{
		ConversationID:  conversationID,
		Content:         params.Content,
		SenderAccountID: params.ClientAccountID,
	}); err != nil {
		return apperrors.NewInternalAppError("create message", err)
	}

	return nil
}

func (s *Service) GetMessages(ctx context.Context, params *model.GetMessagesParams) ([]*model.MessageResponse, error) {
	// Validate input
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get conversation id
	conversationID, err := s.repo.GetConversationID(ctx, &model.GetConversationIDParams{
		Account1ID:    params.ClientAccountID,
		Account2ID:    params.TargetAccountID,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("conversation", "get conversation id", err)
	}

	// Get messages
	messages, err := s.repo.GetMessageList(ctx, &model.GetMessageListParams{
		ConversationID: conversationID,
		Limit:          params.Limit,
		Offset:         params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get messages", err)
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
		return apperrors.NewNotFoundAppError("conversation", "get conversation id", err)
	}

	// Mark messages as read
	if err := s.repo.MarkMessageListAsRead(ctx, &model.MarkMessageListAsReadParams{
		ConversationID:  conversationID,
		SenderAccountID: params.TargetAccountID,
	}); err != nil {
		return apperrors.NewInternalAppError("mark messages as read", err)
	}

	return nil
}