package middleware

import (
	"context"
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
				next.ServeHTTP(w, r)
				return
			}

			// Get account_id by user_id
			ctx := r.Context()
			targetAccountID, err := repo.GetAccountIDByUserID(ctx, userID)
            if err != nil {
				utils.RespondError(w, apperrors.NewNotFoundAppError("account_id", "get account_id by user_id", err))
                return
            }

            // set account_id in context
            ctx = context.WithValue(ctx, key.TargetAccountID, targetAccountID)

			// Get account info
			accountInfo, err := repo.GetAccountInfo(ctx, targetAccountID)
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
