package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) GetNotifications(ctx context.Context, arg *model.GetNotificationsParams) ([]*model.NotificationResponse, error) {
	// Validate input
	if err := arg.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get notifications
	notifications, err := s.repo.GetNotifications(ctx, arg)
	if err != nil {
		return nil, apperrors.NewInternalAppError("get notifications", err)
	}

	// Get sender info
	var senderAccountIDs []string
	for _, notification := range notifications {
		if notification.SenderAccountID != nil {
			senderAccountIDs = append(senderAccountIDs, *notification.SenderAccountID)
		}
	}
	senderInfos, err := s.repo.GetUserInfos(ctx, senderAccountIDs)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("sender info", "get sender infos", err)
	}

	// Convert to response
	notificationsResponse := convertToNotificationResponse(notifications, senderInfos)

	return notificationsResponse, nil
}

func (s *Service) GetUnreadNotifications(ctx context.Context, arg *model.GetUnreadNotificationsParams) ([]*model.NotificationResponse, error) {
	// Validate input
	if err := arg.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get unread notifications
	notifications, err := s.repo.GetUnreadNotifications(ctx, arg)
	if err != nil {
		return nil, err
	}

	// Get sender info
	var senderAccountIDs []string
	for _, notification := range notifications {
		if notification.SenderAccountID != nil {
			senderAccountIDs = append(senderAccountIDs, *notification.SenderAccountID)
		}
	}
	senderInfos, err := s.repo.GetUserInfos(ctx, senderAccountIDs)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("sender info", "get sender infos", err)
	}

	// Convert to response
	notificationsResponse := convertToNotificationResponse(notifications, senderInfos)

	return notificationsResponse, nil
}

func (s *Service) GetUnreadNotificationCount(ctx context.Context, recipientAccountID string) (int64, error) {
	// Get notification count
	count, err := s.repo.GetUnreadNotificationCount(ctx, recipientAccountID)
	if err != nil {
		return 0, apperrors.NewInternalAppError("get unread notification count", err)
	}

	return count, nil
}

func (s *Service) MarkNotificationAsRead(ctx context.Context, arg *model.MarkNotificationAsReadParams) error {
	// Mark notification as read
	if err := s.repo.MarkNotificationAsRead(ctx, arg); err != nil {
		return apperrors.NewInternalAppError("mark notification as read", err)
	}

	return nil
}

func (s *Service) MarkAllNotificationsAsRead(ctx context.Context, recipientAccountID string) error {
	// Mark all notifications as read
	if err := s.repo.MarkAllNotificationsAsRead(ctx, recipientAccountID); err != nil {
		return apperrors.NewInternalAppError("mark all notifications as read", err)
	}

	return nil
}

func convertToNotificationResponse(notifications []*model.Notification, senderInfos []*model.UserInfoInternal) []*model.NotificationResponse {
	// Create user info map
	userInfoMap := make(map[string]*model.UserInfoInternal)
	for _, userInfo := range senderInfos {
		userInfoMap[userInfo.ID] = userInfo
	}

	// Convert to response
	items := []*model.NotificationResponse{}
	for _, notification := range notifications {
		item := &model.NotificationResponse{
			ID:        notification.ID,
			Type:      notification.Type,
			Content:   notification.Content,
			TweetID:   notification.TweetID,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
		}

		if notification.SenderAccountID != nil {
			senderInfo, ok := userInfoMap[*notification.SenderAccountID]
			if ok {
				item.SenderInfo = &model.UserInfo{
					UserID:          senderInfo.UserID,
					UserName:        senderInfo.UserName,
					Bio:             senderInfo.Bio,
					ProfileImageURL: senderInfo.ProfileImageURL,
				}
			}
		}

		items = append(items, item)
	}

	return items
}