package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"net/http"
)

func (s *Service) GetNotifications(ctx context.Context, arg *model.GetNotificationsParams) ([]*model.Notification, error) {
	// Validate input
	if err := arg.Validate(); err != nil {
		return nil, &apperrors.AppError{
			Status: http.StatusBadRequest,
			Code:   "BAD_REQUEST",
			Message: "Invalid request",
			Err: apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err:       err,
				},
			),
		}
	}

	// Get notifications
	notifications, err := s.repo.GetNotifications(ctx, arg)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (s *Service) GetUnreadNotifications(ctx context.Context, arg *model.GetUnreadNotificationParams) ([]*model.Notification, error) {
	// Validate input
	if err := arg.Validate(); err != nil {
		return nil, &apperrors.AppError{
			Status: http.StatusBadRequest,
			Code:   "BAD_REQUEST",
			Message: "Invalid request",
			Err: apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err:       err,
				},
			),
		}
	}

	// Get unread notifications
	notifications, err := s.repo.GetUnreadNotifications(ctx, arg)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (s *Service) GetUnreadNotificationCount(ctx context.Context, recipientAccountID string) (int64, error) {
	// Get notification count
	count, err := s.repo.GetUnreadNotificationCount(ctx, recipientAccountID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *Service) MarkNotificationAsRead(ctx context.Context, arg *model.MarkNotificationAsReadParams) error {
	// Mark notification as read
	if err := s.repo.MarkNotificationAsRead(ctx, arg); err != nil {
		return &apperrors.AppError{
			Status: http.StatusInternalServerError,
			Code:   "DATABASE_ERROR",
			Message: "Failed to mark notification as read",
			Err: apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "mark notification as read",
					Err:       err,
				},
			),
		}
	}

	return nil
}

func (s *Service) MarkAllNotificationsAsRead(ctx context.Context, recipientAccountID string) error {
	// Mark all notifications as read
	if err := s.repo.MarkAllNotificationsAsRead(ctx, recipientAccountID); err != nil {
		return &apperrors.AppError{
			Status: http.StatusInternalServerError,
			Code:   "DATABASE_ERROR",
			Message: "Failed to mark all notifications as read",
			Err: apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "mark all notifications as read",
					Err:       err,
				},
			),
		}
	}

	return nil
}
