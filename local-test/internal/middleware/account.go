package middleware

import (
	"context"
	"errors"
	"local-test/internal/key"
	"local-test/internal/repository"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func AccountInfoMiddleware(repo *repository.Repository) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := mux.Vars(r)["user_id"]
			if userID == "" {
				utils.RespondError(w,
					&apperrors.AppError{
						Status: http.StatusBadRequest,
						Code:   "BAD_REQUEST",
						Message: "User ID is required",
						Err:    &apperrors.ErrInvalidInput{
							Message: "User ID is required",
						},
					},
				)
				return
			}

			// Get account_id by user_id
			ctx := r.Context()
			pathAccountID, err := repo.GetAccountIDByUserID(ctx, userID)
            if err != nil {
				if errors.Is(err, &apperrors.ErrRecordNotFound{}) {
					utils.RespondError(w,
						&apperrors.AppError{
							Status: http.StatusNotFound,
							Code:   "ACCOUNT_NOT_FOUND",
							Message: "Account not found",
							Err:    &apperrors.ErrOperationFailed{
								Operation: "get account_id by user_id",
								Err: err,
							},
						},
					)
				} else {
					utils.RespondError(w,
						&apperrors.AppError{
							Status: http.StatusInternalServerError,
							Code:   "INTERNAL_SERVER_ERROR",
							Message: "Failed to get account_id by user_id",
							Err:    &apperrors.ErrOperationFailed{
								Operation: "get account_id by user_id",
								Err: err,
							},
						},
					)
				}
                return
            }

            // set account_id in context
            ctx = context.WithValue(ctx, key.PathAccountID, pathAccountID)

			// Check if account is admin
			isAdmin, err := repo.IsAdmin(ctx, pathAccountID)
			if err != nil {
				utils.RespondError(w,
					&apperrors.AppError{
						Status: http.StatusInternalServerError,
						Code:   "INTERNAL_SERVER_ERROR",
						Message: "Failed to check if account is admin",
						Err:    &apperrors.ErrOperationFailed{
							Operation: "check if account is admin",
							Err: err,
						},
					},
				)
				return
			}

			// set is_admin in context
			ctx = context.WithValue(ctx, key.IsAdmin, isAdmin)

			// Check if account is suspended
			isSuspended, err := repo.IsSuspended(ctx, pathAccountID)
			if err != nil {
				utils.RespondError(w,
					&apperrors.AppError{
						Status: http.StatusInternalServerError,
						Code:   "INTERNAL_SERVER_ERROR",
						Message: "Failed to check if account is suspended",
						Err:    &apperrors.ErrOperationFailed{
							Operation: "check if account is suspended",
							Err: err,
						},
					},
				)
				return
			}

			// set is_suspended in context
			ctx = context.WithValue(ctx, key.IsSuspended, isSuspended)

			// Call the next handler, which can be another middleware in the chain, or the final handler.
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
