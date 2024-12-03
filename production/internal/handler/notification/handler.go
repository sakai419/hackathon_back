package notification

import (
	"errors"
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type NotificationHandler struct {
	svc *service.Service
}

func NewNotificationHandler(svc *service.Service) ServerInterface {
	return &NotificationHandler{
		svc: svc,
	}
}

// Get notifications
// (GET /notifications)
func (h *NotificationHandler) GetNotifications(w http.ResponseWriter, r *http.Request, params GetNotificationsParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get notifications
	notifications, err := h.svc.GetNotifications(r.Context(), &model.GetNotificationsParams{
		ClientAccountID: clientAccountID,
		Limit:              params.Limit,
		Offset:             params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get notifications", err))
		return
	}

	utils.Respond(w, convertToNotificationResponse(notifications))
}

// Get unread notifications
// (GET /notifications/unread)
func (h *NotificationHandler) GetUnreadNotifications(w http.ResponseWriter, r *http.Request, params GetUnreadNotificationsParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get unread notifications
	notifications, err := h.svc.GetUnreadNotifications(r.Context(), &model.GetUnreadNotificationsParams{
		ClientAccountID: clientAccountID,
		Limit:              params.Limit,
		Offset:             params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get unread notifications", err))
		return
	}

	utils.Respond(w, convertToNotificationResponse(notifications))
}

// Get unread notifications count
// (GET /notifications/unread/count)
func (h *NotificationHandler) GetUnreadNotificationsCount(w http.ResponseWriter, r *http.Request) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get unread notification count
	count, err := h.svc.GetUnreadNotificationCount(r.Context(), clientAccountID)
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get unread notification count", err))
		return
	}

	utils.Respond(w, Count{Count: count})
}

// Mark notification as read
// (PATCH /notifications/{notification_id}/read)
func (h *NotificationHandler) MarkNotificationAsRead(w http.ResponseWriter, r *http.Request, notificationID int64) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Mark notification as read
	if err := h.svc.MarkNotificationAsRead(r.Context(), &model.MarkNotificationAsReadParams{
		RecipientAccountID: clientAccountID,
		ID:                 notificationID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("mark notification as read", err))
		return
	}

	utils.Respond(w, nil)
}

// Mark all notifications as read
// (PATCH /notifications/read/all)
func (h *NotificationHandler) MarkAllNotificationsAsRead(w http.ResponseWriter, r *http.Request) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Mark all notifications as read
	if err := h.svc.MarkAllNotificationsAsRead(r.Context(), clientAccountID); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("mark all notifications as read", err))
		return
	}

	utils.Respond(w, nil)
}

// ErrorHandlerFunc is the error handler for the notification handler
func ErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	var invalidParamFormatError *InvalidParamFormatError
	var requiredParamError *RequiredParamError
	if errors.As(err, &invalidParamFormatError) {
		utils.RespondError(w, apperrors.NewInvalidParamFormatError(
			invalidParamFormatError.ParamName,
			invalidParamFormatError.Err,
		))
		return
	} else if errors.As(err, &requiredParamError) {
		utils.RespondError(w, apperrors.NewRequiredParamError(
			requiredParamError.ParamName,
			requiredParamError,
		))
		return
	} else {
		utils.RespondError(w, apperrors.NewUnexpectedError(err))
		return
	}
}

func convertToNotificationResponse(notifications []*model.NotificationResponse) []Notification {
	res := make([]Notification, len(notifications))
	for i, notification := range notifications {
		res[i] = Notification{
			Id:        notification.ID,
			Type:      notification.Type,
			Content:   notification.Content,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
		}

		if notification.SenderInfo != nil {
			res[i].SenderInfo = &UserInfo{
				UserId:          notification.SenderInfo.UserID,
				UserName:        notification.SenderInfo.UserName,
				ProfileImageUrl: notification.SenderInfo.ProfileImageURL,
				Bio: 		     notification.SenderInfo.Bio,
				IsPrivate: 	     notification.SenderInfo.IsPrivate,
				IsAdmin: 	     notification.SenderInfo.IsAdmin,
				IsFollowing:     notification.SenderInfo.IsFollowing,
				IsFollowed:      notification.SenderInfo.IsFollowed,
			}
		}

		if notification.RelatedTweet != nil {
			res[i].RelatedTweet = &TweetInfo{
				TweetId: notification.RelatedTweet.TweetID,
				UserInfo: UserInfoWithoutBio{
					UserId:          notification.RelatedTweet.UserInfo.UserID,
					UserName:        notification.RelatedTweet.UserInfo.UserName,
					ProfileImageUrl: notification.RelatedTweet.UserInfo.ProfileImageURL,
					IsPrivate: 	     notification.RelatedTweet.UserInfo.IsPrivate,
					IsAdmin: 	     notification.RelatedTweet.UserInfo.IsAdmin,
					IsFollowing:     notification.RelatedTweet.UserInfo.IsFollowing,
					IsFollowed:      notification.RelatedTweet.UserInfo.IsFollowed,
				},
				Content:       notification.RelatedTweet.Content,
				LikesCount:    notification.RelatedTweet.LikesCount,
				RetweetsCount: notification.RelatedTweet.RetweetsCount,
				RepliesCount:  notification.RelatedTweet.RepliesCount,
				IsQuote:       notification.RelatedTweet.IsQuote,
				IsReply:       notification.RelatedTweet.IsReply,
				IsPinned:      notification.RelatedTweet.IsPinned,
				HasLiked:      notification.RelatedTweet.HasLiked,
				HasRetweeted:  notification.RelatedTweet.HasRetweeted,
				CreatedAt:     notification.RelatedTweet.CreatedAt,
			}

			if notification.RelatedTweet.Media != nil {
				res[i].RelatedTweet.Media = &Media{
					Url:  notification.RelatedTweet.Media.URL,
					Type: notification.RelatedTweet.Media.Type,
				}
			}
		}
	}
	return res
}