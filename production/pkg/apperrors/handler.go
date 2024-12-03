package apperrors

import (
	"errors"
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

func NewDecodeError(err error) *AppError {
    var invalidInputError *ErrInvalidInput
    var emptyRequestError *ErrEmptyRequest
    if errors.As(err, &invalidInputError) {
        return &AppError{
            Status:  http.StatusBadRequest,
            Code:    "BAD_REQUEST",
            Message: "Invalid input",
            Err:     WrapHandlerError(
                &ErrOperationFailed{
                    Operation: "decode request",
                    Err: invalidInputError,
                },
            ),
        }
    } else if errors.As(err, &emptyRequestError) {
        return &AppError{
            Status:  http.StatusBadRequest,
            Code:    "BAD_REQUEST",
            Message: "Empty request body",
            Err:     WrapHandlerError(
                &ErrOperationFailed{
                    Operation: "decode request",
                    Err: emptyRequestError,
                },
            ),
        }
    } else {
        return &AppError{
            Status:  http.StatusInternalServerError,
            Code:    "INTERNAL_SERVER_ERROR",
            Message: "Failed to decode request",
            Err:     WrapHandlerError(
                &ErrOperationFailed{
                    Operation: "decode request",
                    Err: err,
                },
            ),
        }
    }
}

func NewHandlerError(operation string, err error) error {
    return WrapHandlerError(
        &ErrOperationFailed{
            Operation: operation,
            Err: err,
        },
    )
}

func NewUnexpectedError(err error) *AppError {
	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "An unexpected error occurred",
		Err:     WrapServiceError(err),
	}
}