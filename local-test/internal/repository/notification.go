package repository

import (
	"context"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) GetNotifications(ctx context.Context, arg *model.GetNotificationsParams) ([]*model.NotificationResponse, error) {
	// Get notifications
	query := sqlcgen.GetNotificationsParams{
		RecipientAccountID: arg.RecipientAccountID,
		Limit:              arg.Limit,
		Offset:             arg.Offset,
	}
	notifications, err := r.q.GetNotifications(ctx, query)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get notifications",
				Err:       err,
			},
		)
	}

	// Convert to model
	items := convertToNotificationResponse(notifications)

	return items, nil
}

func (r *Repository) GetUnreadNotifications(ctx context.Context, arg *model.GetUnreadNotificationsParams) ([]*model.NotificationResponse, error) {
	// Get unread notifications
	query := sqlcgen.GetUnreadNotificationsParams{
		RecipientAccountID: arg.RecipientAccountID,
		Limit:              arg.Limit,
		Offset:             arg.Offset,
	}
	notifications, err := r.q.GetUnreadNotifications(ctx, query)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get unread notifications",
				Err:       err,
			},
		)
	}

	// Convert to model
	items := convertToNotificationResponse(notifications)

	return items, nil
}

func (r *Repository) GetUnreadNotificationCount(ctx context.Context, recipientAccountID string) (int64, error) {
	count, err := r.q.GetUnreadNotificationCount(ctx, recipientAccountID)
	if err != nil {
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get unread notification count",
				Err:       err,
			},
		)
	}

	return count, nil
}

func (r *Repository) MarkNotificationAsRead(ctx context.Context, arg *model.MarkNotificationAsReadParams) error {
	// Mark notification as read
	query := sqlcgen.MarkNotificationAsReadParams{
		ID:                 arg.ID,
		RecipientAccountID: arg.RecipientAccountID,
	}
	err := r.q.MarkNotificationAsRead(ctx, query)
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "mark notification as read",
				Err:       err,
			},
		)
	}

	return nil
}

func (r *Repository) MarkAllNotificationsAsRead(ctx context.Context, recipientAccountID string) error {
	// Mark all notifications as read
	err := r.q.MarkAllNotificationsAsRead(ctx, recipientAccountID)
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "mark all notifications as read",
				Err:       err,
			},
		)
	}

	return nil
}

func convertToNotificationResponse(notifications []sqlcgen.Notification) []*model.NotificationResponse {
	var items []*model.NotificationResponse
	for _, n := range notifications {
		items = append(items, &model.NotificationResponse{
			ID:                 n.ID,
			SenderAccountID:    &n.SenderAccountID.String,
			Type:               string(n.Type),
			Content: 		    &n.Content.String,
			TweetID:            &n.TweetID.Int64,
			IsRead:             n.IsRead,
			CreatedAt:          n.CreatedAt,
		})
	}

	return items
}