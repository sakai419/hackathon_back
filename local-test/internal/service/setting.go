package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"net/http"
)

func (s *Service) UpdateSettings(ctx context.Context, arg *model.UpdateSettingsParams) error {
	// Update settings
	if err := s.repo.UpdateSettings(ctx, arg); err != nil {
		var notFoundErr *apperrors.ErrRecordNotFound
		if errors.As(err, &notFoundErr) {
			return &apperrors.AppError{
				Status:  http.StatusNotFound,
				Code:    "NOT_FOUND",
				Message: "Setting not found",
				Err:     apperrors.WrapServiceError(
					&apperrors.ErrOperationFailed{
						Operation: "update settings",
						Err: notFoundErr,
					},
				),
			}
		}

		return &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to update settings",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "update settings",
					Err: err,
				},
			),
		}
	}

	return nil
}