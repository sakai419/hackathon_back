package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) GetLeftSidebarInfo(ctx context.Context, clientAccountID string) (*model.LeftSidebarInfo, error) {
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

	// convert to model
	ret := &model.LeftSidebarInfo{
		UnreadConversationCount: unreadConversationCount,
		UnreadNotificationCount: unReadNotificationCount,
	}

	return ret, nil
}

func (s *Service) GetRightSidebarInfo(ctx context.Context, clientAccountID string) (*model.RightSidebarInfo, error) {
	// Get recent labels
	recentLabels, err := s.repo.GetRecentLabels(ctx, 100)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("recent labels", "get recent labels", err)
	}

	// Get follow suggestions
	followSuggestionIDs, err := s.repo.GetFollowSuggestions(ctx, clientAccountID)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("follow suggestions", "get follow suggestions", err)
	}

	// Filter accesible account ids
	accessibleAccountIDs, err := s.repo.FilterAccessibleAccountIDs(ctx, &model.FilterAccessibleAccountIDsParams{
		AccountIDs:       followSuggestionIDs,
		ClientAccountID:  clientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("filter accessible account ids", err)
	}

	// Get user infos
	userInfos, err := s.repo.GetUserInfos(ctx, &model.GetUserInfosParams{
		TargetAccountIDs: accessibleAccountIDs,
		ClientAccountID: clientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get user infos", err)
	}

	// convert to model
	followSuggestions := make([]*model.UserInfoWithoutBio, 0)
	for _, userInfo := range userInfos {
		followSuggestions = append(followSuggestions, convertToUserInfoWithoutBio(userInfo))
	}
	ret := &model.RightSidebarInfo{
		RecentLabels:       recentLabels,
		FollowSuggestions:  followSuggestions,
	}

	return ret, nil
}