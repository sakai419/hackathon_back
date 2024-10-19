package notification

import (
	"errors"
	"local-test/internal/key"
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
	// Get client ID
	clientID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get notifications
	arg := &model.GetNotificationsParams{
		RecipientAccountID: clientID,
		Limit:              params.Limit,
		Offset:             params.Offset,
	}
	notifications, err := h.svc.GetNotifications(r.Context(), arg)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, notifications)
}

// Get unread notifications
// (GET /notifications/unread)
func (h *NotificationHandler) GetUnreadNotifications(w http.ResponseWriter, r *http.Request, params GetUnreadNotificationsParams) {
	// Get client ID
	clientID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get unread notifications
	arg := &model.GetUnreadNotificationsParams{
		RecipientAccountID: clientID,
		Limit:              params.Limit,
		Offset:             params.Offset,
	}
	notifications, err := h.svc.GetUnreadNotifications(r.Context(), arg)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Convert to response
	resp := convertToNotificationResponse(notifications)

	utils.Respond(w, resp)
}

// Get unread notifications count
// (GET /notifications/unread/count)
func (h *NotificationHandler) GetUnreadNotificationsCount(w http.ResponseWriter, r *http.Request) {
	// Get client ID
	clientID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get unread notification count
	count, err := h.svc.GetUnreadNotificationCount(r.Context(), clientID)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, count)
}

// Mark notification as read
// (PUT /notifications/{notification_id}/read)
func (h *NotificationHandler) MarkNotificationAsRead(w http.ResponseWriter, r *http.Request, notificationID int64) {
	// Get client ID
	clientID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Mark notification as read
	arg := &model.MarkNotificationAsReadParams{
		RecipientAccountID: clientID,
		ID:                 notificationID,
	}
	if err := h.svc.MarkNotificationAsRead(r.Context(), arg); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// Mark all notifications as read
// (PUT /notifications/read/all)
func (h *NotificationHandler) MarkAllNotificationsAsRead(w http.ResponseWriter, r *http.Request) {
	// Get client ID
	clientID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Mark all notifications as read
	if err := h.svc.MarkAllNotificationsAsRead(r.Context(), clientID); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// ErrHandleFunc handles errors
func ErrHandleFunc(w http.ResponseWriter, r *http.Request, err error) {
	var invalidParamFormatError *InvalidParamFormatError
	if errors.As(err, &invalidParamFormatError) {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid parameter format",
			Err:     err,
		})
		return
	} else {
		utils.RespondError(w, err)
	}
}

// Get client account ID from context
func getClientAccountID(w http.ResponseWriter, r *http.Request) (string, bool) {
	clientID, err := key.GetClientAccountID(r.Context())
	if err != nil {
		utils.RespondError(w,
			&apperrors.AppError{
				Status:  http.StatusInternalServerError,
				Code:    "INTERNAL_SERVER_ERROR",
				Message: "Account ID not found in context",
				Err:     apperrors.WrapHandlerError(
					&apperrors.ErrOperationFailed{
						Operation: "get account ID",
						Err: err,
					},
				),
			},
		)
		return "", false
	}
	return clientID, true
}

func convertToNotificationResponse(notifications []*model.NotificationResponse) []Notification {
	res := make([]Notification, len(notifications))
	for i, notification := range notifications {
		res[i] = Notification{
			Id:        notification.ID,
			SenderInfo: &UserInfo{
				UserId:          notification.SenderInfo.UserID,
				UserName:        notification.SenderInfo.UserName,
				ProfileImageUrl: notification.SenderInfo.ProfileImageURL,
			},
			Type:      notification.Type,
			Content:   notification.Content,
			TweetId:   notification.TweetID,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
		}
	}
	return res
}