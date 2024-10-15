package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"net/http"
)

func (s *Service) CreateAccount(ctx context.Context, arg *model.CreateAccountParams) error {
	// Validate params
	if err := arg.Validate(); err != nil {
		return &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		}
	}

	// Create account
    if err := s.repo.CreateAccount(ctx, arg); err != nil {
		// Check if the error is a duplicate entry error
		var duplicateErr *apperrors.ErrDuplicateEntry
		if errors.As(err, &duplicateErr) {
			return &apperrors.AppError{
				Status:  http.StatusConflict,
				Code:    "DUPLICATE_ENTRY",
				Message: "Account already exists",
				Err:     apperrors.WrapServiceError(
					&apperrors.ErrOperationFailed{
						Operation: "create account",
						Err: duplicateErr,
					},
				),
			}
		}

        return &apperrors.AppError{
            Status:  http.StatusInternalServerError,
            Code:    "DATABASE_ERROR",
            Message: "Failed to create account",
            Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "create account",
					Err: err,
				},
			),
        }
    }

    return nil
}

func (s *Service) DeleteMyAccount(ctx context.Context, id string) error {
	if err := s.repo.DeleteMyAccount(ctx, id); err != nil {
		// Check if the error is a record not found error
		var notFoundErr *apperrors.ErrRecordNotFound
		if errors.As(err, &notFoundErr) {
			return &apperrors.AppError{
				Status:  http.StatusNotFound,
				Code:    "ACCOUNT_NOT_FOUND",
				Message: "Account not found",
				Err:     apperrors.WrapServiceError(
					&apperrors.ErrOperationFailed{
						Operation: "delete account",
						Err: notFoundErr,
					},
				),
			}
		}

		return &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to delete account",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "delete account",
					Err: err,
				},
			),
		}
	}

	return nil
}