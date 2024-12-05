package apperrors

import (
	"errors"
	"fmt"
	"net/http"
)

func NewValidateAppError(err error) *AppError {
	return &AppError{
		Status:  http.StatusBadRequest,
		Code:    "BAD_REQUEST",
		Message: "Invalid request",
		Err:     WrapServiceError(
			&ErrOperationFailed{
				Operation: "validate request",
				Err: err,
			},
		),
	}
}

func NewInternalAppError(operation string, err error) *AppError {
	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "Internal server error",
		Err:     WrapServiceError(
			&ErrOperationFailed{
				Operation: operation,
				Err: err,
			},
		),
	}
}

func NewForbiddenAppError(operation string, err error) *AppError {
	return &AppError{
		Status:  http.StatusForbidden,
		Code:    "FORBIDDEN",
		Message: fmt.Sprintf("%s is forbidden", operation),
		Err:     WrapServiceError(
			&ErrOperationFailed{
				Operation: operation,
				Err: err,
			},
		),
	}
}

func NewBlockedAppError(operation string, err error) *AppError {
	return &AppError{
		Status:  http.StatusForbidden,
		Code:    "BLOCKED",
		Message: "You are blocked by target user",
		Err:     WrapServiceError(
			&ErrOperationFailed{
				Operation: operation,
				Err: err,
			},
		),
	}
}

func NewBlockingAppError(operation string, err error) *AppError {
	return &AppError{
		Status:  http.StatusForbidden,
		Code:    "BLOCKING",
		Message: "You are blocking target user",
		Err:     WrapServiceError(
			&ErrOperationFailed{
				Operation: operation,
				Err: err,
			},
		),
	}
}

func NewPrivateAccountAccessError(operation string, err error) *AppError {
	return &AppError{
		Status:  http.StatusForbidden,
		Code:    "PRIVATE_ACCOUNT",
		Message: "You are not allowed to access this private account",
		Err:     WrapServiceError(
			&ErrOperationFailed{
				Operation: operation,
				Err:       err,
			},
		),
	}
}

func NewDuplicateEntryAppError(entity, operation string, err error) *AppError {
	var duplicateErr *ErrDuplicateEntry
	if errors.As(err, &duplicateErr) {
		return &AppError{
			Status:  http.StatusConflict,
			Code:    "DUPLICATE_ENTRY",
			Message: fmt.Sprintf("%s already exists", entity),
			Err:     WrapServiceError(
				&ErrOperationFailed{
					Operation: operation,
					Err: duplicateErr,
				},
			),
		}
	}

	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    "DATABASE_ERROR",
		Message: fmt.Sprintf("Failed to %s", operation),
		Err:     WrapServiceError(
			&ErrOperationFailed{
				Operation: operation,
				Err: err,
			},
		),
	}
}

func NewNotFoundAppError(entity, operation string, err error) *AppError {
	var notFoundErr *ErrRecordNotFound
	if errors.As(err, &notFoundErr) {
		return &AppError{
			Status:  http.StatusNotFound,
			Code:    "NOT_FOUND",
			Message: fmt.Sprintf("%s not found", entity),
			Err:     WrapServiceError(
				&ErrOperationFailed{
					Operation: operation,
					Err: notFoundErr,
				},
			),
		}
	}

	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    "DATABASE_ERROR",
		Message: fmt.Sprintf("Failed to %s", operation),
		Err:     WrapServiceError(
			&ErrOperationFailed{
				Operation: operation,
				Err: err,
			},
		),
	}
}