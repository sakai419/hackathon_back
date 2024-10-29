package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) GetConversations(ctx context.Context, params *model.GetConversationsParams) ([]*model.ConversationResponse, error) {
	// Validate input
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get conversations
	conversations, err := s.repo.GetConversationList(ctx, &model.GetConversationListParams{
		AccountID: params.ClientAccountID,
		Limit:     params.Limit,
		Offset:    params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get conversations", err)
	}

	// Get opponent info
	ids := make([]string, 0)
	for _, conversation := range conversations {
		ids = append(ids, conversation.OpponentID)
	}
	opponentInfos, err := s.repo.GetUserInfos(ctx, ids)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("opponent info", "get opponent infos", err)
	}

	// Convert to response
	conversationResponses := convertToConversationResponse(conversations, opponentInfos)

	return conversationResponses, nil
}

func (s *Service) GetUnreadConversationCount(ctx context.Context, AccountID string) (int64, error) {
	count, err := s.repo.GetUnreadConversationCount(ctx, AccountID)
	if err != nil {
		return 0, apperrors.NewInternalAppError("get unread conversation count", err)
	}

	return count, nil
}

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

	// Mark messages as read
	if err := s.repo.MarkMessageListAsRead(ctx, &model.MarkMessageListAsReadParams{
		ConversationID:  conversationID,
		SenderAccountID: params.TargetAccountID,
	}); err != nil {
		return apperrors.NewInternalAppError("mark messages as read", err)
	}

	return nil
}

func convertToConversationResponse(conversations []*model.Conversation, opponentInfos []*model.UserInfoInternal) []*model.ConversationResponse {
	var conversationResponses []*model.ConversationResponse

	// convert opponentInfos to map
	opponentInfoMap := make(map[string]*model.UserInfoInternal)
	for _, opponentInfo := range opponentInfos {
		opponentInfoMap[opponentInfo.ID] = opponentInfo
	}

	// convert conversations to response
	for _, conversation := range conversations {
		if opponentInfo, exists := opponentInfoMap[conversation.OpponentID]; exists {
			info := model.UserInfoWithoutBio{
				UserID:          opponentInfo.UserID,
				UserName:        opponentInfo.UserName,
				ProfileImageURL: opponentInfo.ProfileImageURL,
			}

			conversationResponses = append(conversationResponses, &model.ConversationResponse{
				ID:              conversation.ID,
				OpponentInfo:    info,
				LastMessageTime: conversation.LastMessageTime,
				Content:         conversation.Content,
				SenderUserID:    conversation.SenderUserID,
				IsRead:          conversation.IsRead,
			})
		}
	}

	return conversationResponses
}
