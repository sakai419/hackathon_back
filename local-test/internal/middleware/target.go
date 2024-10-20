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

func GetTargetInfoMiddleware(repo *repository.Repository) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := mux.Vars(r)["user_id"]
			if userID == "" {
				utils.RespondError(w, apperrors.NewRequiredParamError("user_id", errors.New("user_id is required")))
				return
			}

			// Get account_id by user_id
			ctx := r.Context()
			pathAccountID, err := repo.GetAccountIDByUserID(ctx, userID)
            if err != nil {
				utils.RespondError(w, apperrors.NewNotFoundAppError("account_id", "get account_id by user_id", err))
                return
            }

            // set account_id in context
            ctx = context.WithValue(ctx, key.TargetAccountID, pathAccountID)

			// Get account info
			accountInfo, err := repo.GetAccountInfo(ctx, pathAccountID)
			if err != nil {
				utils.RespondError(w, apperrors.NewNotFoundAppError("account info", "get account info", err))
				return
			}

			// set is_suspended in context
			ctx = context.WithValue(ctx, key.IsTargetSuspended, accountInfo.IsSuspended)

			// set is_private in context
			ctx = context.WithValue(ctx, key.IsTargetPrivate, accountInfo.IsPrivate)

			// Call the next handler, which can be another middleware in the chain, or the final handler.
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
