package middleware

import (
	"context"
	"fmt"
	"local-test/internal/key"
	"local-test/internal/repository"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
)

func AuthClientMiddleware(client *auth.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.RespondError(w,
					&apperrors.AppError{
						Status: http.StatusUnauthorized,
						Code:   "UNAUTHORIZED",
						Message: "Authorization header is required",
						Err:    &apperrors.ErrOperationFailed{
							Operation: "get authorization header",
							Err: fmt.Errorf("authorization header is required"),
						},
					},
				)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			uid, err := authenticate(token, client)
			if err != nil {
				utils.RespondError(w,
					&apperrors.AppError{
						Status: http.StatusUnauthorized,
						Code:   "UNAUTHORIZED",
						Message: "Failed to authenticate token",
						Err:    &apperrors.ErrOperationFailed{
							Operation: "authenticate token",
							Err: err,
						},
					},
				)
				return
			}

			// Add the user ID to the request context
			ctx := context.WithValue(r.Context(), key.ClientAccountID, uid)

			// Call the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthClientAndGetInfoMiddleware(repo *repository.Repository, client *auth.Client) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
				utils.RespondError(w,
					&apperrors.AppError{
						Status: http.StatusUnauthorized,
						Code:   "UNAUTHORIZED",
						Message: "Authorization header is required",
						Err:    &apperrors.ErrOperationFailed{
							Operation: "get authorization header",
							Err: fmt.Errorf("authorization header is required"),
						},
					},
				)
                return
            }

            token := strings.TrimPrefix(authHeader, "Bearer ")
            uid, err := authenticate(token, client)
            if err != nil {
				utils.RespondError(w,
					&apperrors.AppError{
						Status: http.StatusUnauthorized,
						Code:   "UNAUTHORIZED",
						Message: "Failed to authenticate token",
						Err:    &apperrors.ErrOperationFailed{
							Operation: "authenticate token",
							Err: err,
						},
					},
				)
                return
            }

            // Add the user ID to the request context
            ctx := context.WithValue(r.Context(), key.ClientAccountID, uid)

			// Get account info
			accountInfo, err := repo.GetAccountInfo(ctx, uid)
			if err != nil {
				utils.RespondError(w, apperrors.NewNotFoundAppError("account info", "get account info", err))
				return
			}

			// Add the is_suspended flag to the request context
			ctx = context.WithValue(ctx, key.IsClientSuspended, accountInfo.IsSuspended)

			// Add the is_admin flag to the request context
			ctx = context.WithValue(ctx, key.IsClientAdmin, accountInfo.IsAdmin)

			// Call the next handler
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

// authenticate verifies the token and returns the user ID
func authenticate(token string, client *auth.Client) (string, error) {
    tokenDecoded, err := client.VerifyIDToken(context.Background(), token)
    if err != nil {
        return "", fmt.Errorf("error verifying ID token: %v", err)
    }
    return tokenDecoded.UID, nil
}