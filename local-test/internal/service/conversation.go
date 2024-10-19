package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"net/http"
)

func (s *Service) GetConversations(ctx context.Context, params *model.GetConversationsParams) ([]*model.ConversationResponse, error) {
	// Validate input
	if err := params.Validate(); err != nil {
		return nil, &apperrors.AppError{
			Status: http.StatusBadRequest,
			Code:   "INVALID_INPUT",
			Message: "Invalid input",
			Err:    apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate input",
					Err:       err,
				},
			),
		}
	}

	// Get conversations
	conversations, err := s.repo.GetConversationList(ctx, &model.GetConversationListParams{
		AccountID: params.ClientAccountID,
		Limit:     params.Limit,
		Offset:    params.Offset,
	})
	if err != nil {
		return nil, err
	}

	// Get opponent info
	ids := make([]string, 0)
	for _, conversation := range conversations {
		ids = append(ids, conversation.OpponentID)
	}
	opponentInfos, err := s.repo.GetUserInfos(ctx, ids)
	if err != nil {
		return nil, &apperrors.AppError{
			Status: http.StatusInternalServerError,
			Code:   "INTERNAL_SERVER_ERROR",
			Message: "Failed to get opponent info",
			Err:    apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "get opponent info",
					Err:       err,
				},
			),
		}
	}

	// Convert to response
	conversationResponses := convertToConversationResponse(conversations, opponentInfos)

	return conversationResponses, nil
}

func (s *Service) GetUnreadConversationCount(ctx context.Context, AccountID string) (int64, error) {
	count, err := s.repo.GetUnreadConversationCount(ctx, AccountID)
	if err != nil {
		return 0, &apperrors.AppError{
			Status: http.StatusInternalServerError,
			Code:   "INTERNAL_SERVER_ERROR",
			Message: "Failed to get unread conversation count",
			Err:    apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "get unread conversation count",
					Err:       err,
				},
			),
		}
	}

	return count, nil
}

func convertToConversationResponse(conversations []*model.Conversation, opponentInfos []*model.UserInfoInternal) []*model.ConversationResponse {
	var conversationResponses []*model.ConversationResponse
	for i := range len(conversations) {
		info := model.UserInfoWithoutBio{
			UserID:          opponentInfos[i].UserID,
			UserName:        opponentInfos[i].UserName,
			ProfileImageURL: opponentInfos[i].ProfileImageURL,
		}
		conversationResponses = append(conversationResponses, &model.ConversationResponse{
			ID:              conversations[i].ID,
			OpponentInfo:    info,
			LastMessageTime: conversations[i].LastMessageTime,
			Content:         conversations[i].Content,
			SenderID:        conversations[i].SenderAccountID,
			IsRead:          conversations[i].IsRead,
		})
	}

	return conversationResponses
}