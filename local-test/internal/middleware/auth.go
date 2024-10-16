package middleware

import (
	"context"
	"fmt"
	"local-test/internal/key"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
)

// AuthMiddleware is a middleware that authenticates the user
func AuthMiddleware(client *auth.Client) func(http.Handler) http.Handler {
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
            ctx := context.WithValue(r.Context(), key.ClientAccountIDKey, uid)
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