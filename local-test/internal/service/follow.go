package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/utils"
	"net/http"
)

func (s *Service) FollowAndNotify(ctx context.Context, arg *model.FollowAndNotifyParams) error {
	// Validate params
	if err := arg.Validate(); err != nil {
		return &utils.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     utils.WrapServiceError(
				&utils.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		}
	}

	// Create follow
	if err := s.repo.FollowAndNotify(ctx, arg); err != nil {
		// Check if the error is a duplicate entry error
		var duplicateErr *utils.ErrDuplicateEntry
		if errors.As(err, &duplicateErr) {
			return &utils.AppError{
				Status:  http.StatusConflict,
				Code:    "DUPLICATE_ENTRY",
				Message: "Follow already exists",
				Err:     utils.WrapServiceError(
					&utils.ErrOperationFailed{
						Operation: "create follow",
						Err: duplicateErr,
					},
				),
			}
		}

		return &utils.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to create follow",
			Err:     utils.WrapServiceError(
				&utils.ErrOperationFailed{
					Operation: "create follow",
					Err: err,
				},
			),
		}
	}

	return nil
}

func (s *Service) Unfollow(ctx context.Context, arg *model.UnfollowParams) error {
	// Validate params
	if err := arg.Validate(); err != nil {
		return &utils.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     utils.WrapServiceError(
				&utils.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		}
	}

	// Unfollow
	if err := s.repo.Unfollow(ctx, arg); err != nil {
		// Check if the error is a record not found error
		var notFoundErr *utils.ErrRecordNotFound
		if errors.As(err, &notFoundErr) {
			return &utils.AppError{
				Status:  http.StatusNotFound,
				Code:    "FOLLOW_NOT_FOUND",
				Message: "Follow not found",
				Err:     utils.WrapServiceError(
					&utils.ErrOperationFailed{
						Operation: "unfollow",
						Err: notFoundErr,
					},
				),
			}
		}

		return &utils.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to unfollow",
			Err:     utils.WrapServiceError(
				&utils.ErrOperationFailed{
					Operation: "unfollow",
					Err: err,
				},
			),
		}
	}

	return nil
}