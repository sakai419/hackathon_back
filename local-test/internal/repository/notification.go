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
		RecipientAccountID: params.ClientAccountID,
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
		RecipientAccountID: params.ClientAccountID,
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
	// Get unread notification count
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
	res, err := r.q.MarkNotificationAsRead(ctx, sqlcgen.MarkNotificationAsReadParams{
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

	// Check if notification is not found
	num, err := res.RowsAffected()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "mark notification as read",
				Err:       err,
			},
		)
	}
	if num == 0 {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "notification",
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

		var senderAccountID *string
		if n.SenderAccountID.Valid {
			senderAccountID = &n.SenderAccountID.String
		} else {
			senderAccountID = nil
		}

		var tweetID *int64
		if n.TweetID.Valid {
			tweetID = &n.TweetID.Int64
		} else {
			tweetID = nil
		}

		var content *string
		if n.Content.Valid {
			content = &n.Content.String
		} else {
			content = nil
		}

		items = append(items, &model.Notification{
			ID:                 n.ID,
			SenderAccountID:    senderAccountID,
			Type:               string(n.Type),
			Content: 		    content,
			TweetID:            tweetID,
			IsRead:             n.IsRead,
			CreatedAt:          n.CreatedAt,
		})
	}

	return items
}