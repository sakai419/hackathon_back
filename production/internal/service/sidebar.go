package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) GetSidebarInfo(ctx context.Context, clientAccountID string) (*model.SidebarInfo, error) {
	// Get unread conversation count
	unreadConversationCount, err := s.repo.GetUnreadConversationCount(ctx, clientAccountID)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("unread message count", "get unread message count", err)
	}

	// Get unread notification count
	unReadNotificationCount, err := s.repo.GetUnreadNotificationCount(ctx, clientAccountID)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("unread notification count", "get unread notification count", err)
	}

	return convertToSidebarInfo(unreadConversationCount, unReadNotificationCount), nil

}

func convertToSidebarInfo(unreadConversationCount int64, unReadNotificationCount int64) *model.SidebarInfo {
	ret := &model.SidebarInfo{
		UnreadConversationCount: unreadConversationCount,
		UnreadNotificationCount: unReadNotificationCount,
	}

	return ret
}

