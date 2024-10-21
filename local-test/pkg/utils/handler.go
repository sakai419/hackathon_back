package utils

import (
	"errors"
	"fmt"
	"local-test/internal/key"
	"local-test/pkg/apperrors"
	"net/http"
	"reflect"
)

func GetClientAccountID(w http.ResponseWriter, r *http.Request) (string, bool) {
	clientID, err := key.GetClientAccountID(r.Context())
	if err != nil {
		RespondError(w,
			&apperrors.AppError{
				Status:  http.StatusInternalServerError,
				Code:    "INTERNAL_SERVER_ERROR",
				Message: "Account ID not found in context",
				Err:     apperrors.WrapHandlerError(
					&apperrors.ErrOperationFailed{
						Operation: "get account ID",
						Err: err,
					},
				),
			},
		)
		return "", false
	}
	return clientID, true
}

func GetTargetAccountID(w http.ResponseWriter, r *http.Request) (string, bool) {
	accountID, err := key.GetTargetAccountID(r.Context())
	if err != nil {
		RespondError(w,
			&apperrors.AppError{
				Status:  http.StatusBadRequest,
				Code:    "BAD_REQUEST",
				Message: "Account ID not found in path",
				Err:     apperrors.WrapHandlerError(
					&apperrors.ErrOperationFailed{
						Operation: "get account ID",
						Err: err,
					},
				),
			},
		)
		return "", false
	}
	return accountID, true
}

func IsClientSuspended(w http.ResponseWriter, r *http.Request) bool {
	isClientSuspended, err := key.GetIsClientSuspended(r.Context())
	if err != nil {
		RespondError(w, &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Failed to get is_suspended",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get is_suspended",
					Err: err,
				},
			),
		})
		return true
	}

	if isClientSuspended {
		RespondError(w, &apperrors.AppError{
			Status:  http.StatusForbidden,
			Code:    "FORBIDDEN",
			Message: "User is suspended",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrForbidden{
					Message: "User is suspended",
				},
			),
		})
		return true
	}

	return false
}

func IsTargetSuspended(w http.ResponseWriter, r *http.Request) bool {
	isTargetSuspended, err := key.GetIsTargetSuspended(r.Context())
	if err != nil {
		RespondError(w, &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Failed to get is_suspended",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get is_suspended",
					Err: err,
				},
			),
		})
		return true
	}

	if isTargetSuspended {
		RespondError(w, &apperrors.AppError{
			Status:  http.StatusForbidden,
			Code:    "FORBIDDEN",
			Message: "User is suspended",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrForbidden{
					Message: "User is suspended",
				},
			),
		})
		return true
	}

	return false
}

func IsTargetPrivate(w http.ResponseWriter, r *http.Request) bool {
	isTargetPrivate, err := key.GetIsTargetPrivate(r.Context())
	if err != nil {
		RespondError(w, &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Failed to get is_private",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get is_private",
					Err: err,
				},
			),
		})
		return true
	}

	if isTargetPrivate {
		RespondError(w, &apperrors.AppError{
			Status:  http.StatusForbidden,
			Code:    "FORBIDDEN",
			Message: "User is private",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrForbidden{
					Message: "User is private",
				},
			),
		})
		return true
	}

	return false
}

func ValidateRequiredFields(req interface{}) error {
    v := reflect.ValueOf(req)
    if v.Kind() != reflect.Struct {
        return &apperrors.ErrInvalidRequest{
            Entity: "All fields",
            Err:    errors.New("provided value is not a struct"),
        }
    }

    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        fieldType := v.Type().Field(i)

        // ポインタ型ではない変数かつゼロ値の場合にエラーを返す
        if fieldType.Type.Kind() != reflect.Ptr && field.IsZero() {
            return &apperrors.ErrInvalidRequest{
                Entity: fieldType.Name,
                Err:    fmt.Errorf("%s is a zero value", fieldType.Name),
            }
        }
    }
    return nil
}