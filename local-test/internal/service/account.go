package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/utils"
	"net/http"
)

func (s *Service) CreateAccount(ctx context.Context, arg *model.CreateAccountParams) error {
	// Validate params
	if err := validateCreateAccountParams(arg); err != nil {
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

	// Create account
    if err := s.repo.CreateAccount(ctx, arg); err != nil {
		// Check if the error is a duplicate entry error
		var duplicateErr *utils.ErrDuplicateEntry
		if errors.As(err, &duplicateErr) {
			return &utils.AppError{
				Status:  http.StatusConflict,
				Code:    "DUPLICATE_ENTRY",
				Message: "Account already exists",
				Err:     utils.WrapServiceError(
					&utils.ErrOperationFailed{
						Operation: "create account",
						Err: duplicateErr,
					},
				),
			}
		}

        return &utils.AppError{
            Status:  http.StatusInternalServerError,
            Code:    "DATABASE_ERROR",
            Message: "Failed to create account",
            Err:     utils.WrapServiceError(
				&utils.ErrOperationFailed{
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
		var notFoundErr *utils.ErrRecordNotFound
		if errors.As(err, &notFoundErr) {
			return &utils.AppError{
				Status:  http.StatusNotFound,
				Code:    "ACCOUNT_NOT_FOUND",
				Message: "Account not found",
				Err:     utils.WrapServiceError(
					&utils.ErrOperationFailed{
						Operation: "delete account",
						Err: notFoundErr,
					},
				),
			}
		}

		return &utils.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to delete account",
			Err:     utils.WrapServiceError(
				&utils.ErrOperationFailed{
					Operation: "delete account",
					Err: err,
				},
			),
		}
	}

	return nil
}

// Validate params
func validateCreateAccountParams(arg *model.CreateAccountParams) error {
	if len(arg.ID) != 28 {
		return &utils.ErrInvalidInput{Message: "invalid firebase uid"}
	}
	if len(arg.UserID) > 30 {
		return &utils.ErrInvalidInput{Message: "user id is too long"}
	}
	if len(arg.UserName) > 30 {
		return &utils.ErrInvalidInput{Message: "user name is too long"}
	}
	return nil
}