package repository

import (
	"context"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) GetNotifications(ctx context.Context, params *model.GetNotificationsParams) ([]*model.Notification, error) {
	// Get notifications
	notifications, err := r.q.GetNotifications(ctx, sqlcgen.GetNotificationsParams{
		RecipientAccountID: params.RecipientAccountID,
		Limit:              params.Limit,
		Offset:             params.Offset,
	})
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

func (r *Repository) GetUnreadNotifications(ctx context.Context, params *model.GetUnreadNotificationsParams) ([]*model.Notification, error) {
	// Get unread notifications
	notifications, err := r.q.GetUnreadNotifications(ctx, sqlcgen.GetUnreadNotificationsParams{
		RecipientAccountID: params.RecipientAccountID,
		Limit:              params.Limit,
		Offset:             params.Offset,
	})
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

func (r *Repository) GetUnreadNotificationCount(ctx context.Context, accountID string) (int64, error) {
	count, err := r.q.GetUnreadNotificationCount(ctx, accountID)
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

func (r *Repository) MarkNotificationAsRead(ctx context.Context, params *model.MarkNotificationAsReadParams) error {
	// Mark notification as read
	err := r.q.MarkNotificationAsRead(ctx, sqlcgen.MarkNotificationAsReadParams{
		ID:                 params.ID,
		RecipientAccountID: params.RecipientAccountID,
	})
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

func (r *Repository) MarkAllNotificationsAsRead(ctx context.Context, accountID string) error {
	// Mark all notifications as read
	err := r.q.MarkAllNotificationsAsRead(ctx, accountID)
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

func convertToNotificationResponse(notifications []sqlcgen.Notification) []*model.Notification {
	var items []*model.Notification
	for _, n := range notifications {
		items = append(items, &model.Notification{
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