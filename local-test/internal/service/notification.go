package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) GetNotifications(ctx context.Context, params *model.GetNotificationsParams) ([]*model.NotificationResponse, error) {
	// Validate input
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get notifications
	notifications, err := s.repo.GetNotifications(ctx, params)
	if err != nil {
		return nil, apperrors.NewInternalAppError("get notifications", err)
	}

	// Get sender info
	senderAccountIDs := convertToSenderAccountIDs(notifications)
	senderInfos, err := s.repo.GetUserInfos(ctx, senderAccountIDs)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("sender info", "get sender infos", err)
	}

	// Get Tweet infos
	tweetIDs := extractTweetIDs(notifications)
	tweetInfos, err := s.repo.GetTweetInfosByIDs(ctx, &model.GetTweetInfosByIDsParams{
		TweetIDs: tweetIDs,
		ClientAccountID: params.RecipientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("tweet info", "get tweet infos", err)
	}

	// Get client user info
	clientUserInfo, err := s.repo.GetUserInfo(ctx, params.RecipientAccountID)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("client user info", "get client user info", err)
	}

	// Convert to response
	notificationsResponse := convertToNotificationResponse(notifications, senderInfos, tweetInfos, clientUserInfo)

	return notificationsResponse, nil
}

func (s *Service) GetUnreadNotifications(ctx context.Context, params *model.GetUnreadNotificationsParams) ([]*model.NotificationResponse, error) {
	// Validate input
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get unread notifications
	notifications, err := s.repo.GetUnreadNotifications(ctx, params)
	if err != nil {
		return nil, err
	}

	// Get sender info
    senderAccountIDs := convertToSenderAccountIDs(notifications)
	senderInfos, err := s.repo.GetUserInfos(ctx, senderAccountIDs)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("sender info", "get sender infos", err)
	}

	// Get Tweet infos
	tweetIDs := extractTweetIDs(notifications)
	tweetInfos, err := s.repo.GetTweetInfosByIDs(ctx, &model.GetTweetInfosByIDsParams{
		TweetIDs: tweetIDs,
		ClientAccountID: params.RecipientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("tweet info", "get tweet infos", err)
	}

	// Get client user info
	clientUserInfo, err := s.repo.GetUserInfo(ctx, params.RecipientAccountID)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("client user info", "get client user info", err)
	}

	// Convert to response
	notificationsResponse := convertToNotificationResponse(notifications, senderInfos, tweetInfos, clientUserInfo)

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

func (s *Service) MarkNotificationAsRead(ctx context.Context, params *model.MarkNotificationAsReadParams) error {
	// Mark notification as read
	if err := s.repo.MarkNotificationAsRead(ctx, params); err != nil {
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

func convertToSenderAccountIDs(notifications []*model.Notification) []string {
	accountIDMap := make(map[string]bool)
	senderAccountIDs := make([]string, 0)
	for _, notification := range notifications {
		if notification.SenderAccountID != nil {
			if !accountIDMap[*notification.SenderAccountID] {
				accountIDMap[*notification.SenderAccountID] = true
				senderAccountIDs = append(senderAccountIDs, *notification.SenderAccountID)
			}
		}
	}

	return senderAccountIDs
}

func extractTweetIDs(notifications []*model.Notification) []int64 {
	tweetIDMap := make(map[int64]bool)
	tweetIDs := make([]int64, 0)
	for _, notification := range notifications {
		if notification.TweetID != nil {
			if !tweetIDMap[*notification.TweetID] {
				tweetIDMap[*notification.TweetID] = true
				tweetIDs = append(tweetIDs, *notification.TweetID)
			}
		}
	}

	return tweetIDs
}

func convertToNotificationResponse(notifications []*model.Notification, senderInfos []*model.UserInfoInternal, tweetInfos []*model.TweetInfoInternal, clientUserInfo *model.UserInfoInternal) []*model.NotificationResponse {
	// Create user info map
	userInfoMap := make(map[string]*model.UserInfoInternal)
	for _, userInfo := range senderInfos {
		userInfoMap[userInfo.ID] = userInfo
	}

	// Create tweet info map
	tweetInfoMap := make(map[int64]*model.TweetInfoInternal)
	for _, tweetInfo := range tweetInfos {
		tweetInfoMap[tweetInfo.TweetID] = tweetInfo
	}

	// Convert to response
	items := []*model.NotificationResponse{}
	for _, notification := range notifications {
		item := &model.NotificationResponse{
			ID:        notification.ID,
			Type:      notification.Type,
			Content:   notification.Content,
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

		if notification.TweetID != nil {
			tweetInfo, ok := tweetInfoMap[*notification.TweetID]
			if ok {
				item.RelatedTweet = &model.TweetInfo{
					TweetID:       tweetInfo.TweetID,
					UserInfo:      model.UserInfoWithoutBio{
						UserID:          clientUserInfo.UserID,
						UserName: 	     clientUserInfo.UserName,
						ProfileImageURL: clientUserInfo.ProfileImageURL,
					},
					Content:       tweetInfo.Content,
					Code:          tweetInfo.Code,
					Media:         tweetInfo.Media,
					LikesCount:    tweetInfo.LikesCount,
					RetweetsCount: tweetInfo.RetweetsCount,
					RepliesCount:  tweetInfo.RepliesCount,
					IsQuote:       tweetInfo.IsQuote,
					IsReply:       tweetInfo.IsReply,
					IsPinned:      tweetInfo.IsPinned,
					HasLiked:      tweetInfo.HasLiked,
					HasRetweeted:  tweetInfo.HasRetweeted,
					CreatedAt:     tweetInfo.CreatedAt,
				}
			}
		}


		items = append(items, item)
	}

	return items
}