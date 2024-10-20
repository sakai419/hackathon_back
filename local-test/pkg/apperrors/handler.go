package apperrors

import (
	"fmt"
	"net/http"
)

func NewInvalidParamFormatError(param string, err error) *AppError {
	return &AppError{
		Status: http.StatusBadRequest,
		Code:   "BAD_REQUEST",
		Message: fmt.Sprintf("Invalid parameter format: %s", param),
		Err:    WrapHandlerError(
			&ErrOperationFailed{
				Operation: "check parameter format",
				Err:      err,
			},
		),
	}
}

func NewRequiredParamError(param string, err error) *AppError {
	return &AppError{
		Status: http.StatusBadRequest,
		Code:   "BAD_REQUEST",
		Message: fmt.Sprintf("Required parameter is missing: %s", param),
		Err:    WrapHandlerError(
			&ErrOperationFailed{
				Operation: "check required parameter",
				Err:      err,
			},
		),
	}
}

func NewGetAuthHeaderError() *AppError {
	return &AppError{
		Status: http.StatusUnauthorized,
		Code:   "UNAUTHORIZED",
		Message: "Authorization header is required",
		Err:    WrapHandlerError(
			&ErrOperationFailed{
				Operation: "get authorization header",
				Err:      fmt.Errorf("authorization header is required"),
			},
		),
	}
}

func NewAuthenticateTokenError(err error) *AppError {
	return &AppError{
		Status: http.StatusUnauthorized,
		Code:   "UNAUTHORIZED",
		Message: "Failed to authenticate token",
		Err:    WrapHandlerError(
			&ErrOperationFailed{
				Operation: "authenticate token",
				Err:      err,
			},
		),
	}
}

func NewUnexpectedError(err error) *AppError {
	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "An unexpected error occurred",
		Err:     WrapServiceError(err),
	}
}