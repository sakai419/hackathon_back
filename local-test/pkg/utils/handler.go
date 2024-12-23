package utils

import (
	"local-test/internal/key"
	"local-test/pkg/apperrors"
	"net/http"
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

func IsNotTargetPrivate(w http.ResponseWriter, r *http.Request) bool {
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

	if !isTargetPrivate {
		RespondError(w, &apperrors.AppError{
			Status:  http.StatusForbidden,
			Code:    "FORBIDDEN",
			Message: "User is not private",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrForbidden{
					Message: "User is not private",
				},
			),
		})
		return true
	}

	return false
}

func IsClientAdmin(w http.ResponseWriter, r *http.Request) bool {
	isClientAdmin, err := key.GetIsClientAdmin(r.Context())
	if err != nil {
		RespondError(w, &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Failed to get is_admin",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get is_admin",
					Err: err,
				},
			),
		})
		return false
	}

	if !isClientAdmin {
		RespondError(w, &apperrors.AppError{
			Status:  http.StatusForbidden,
			Code:    "FORBIDDEN",
			Message: "User is not admin",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrForbidden{
					Message: "User is not admin",
				},
			),
		})
		return false
	}

	return true
}
