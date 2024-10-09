package middleware

import (
	"context"
	"fmt"
	contextKey "local-test/internal/context"
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
                http.Error(w, "Missing auth token", http.StatusUnauthorized)
                return
            }

            token := strings.TrimPrefix(authHeader, "Bearer ")
            uid, err := authenticate(token, client)
            if err != nil {
                http.Error(w, "Invalid auth token", http.StatusUnauthorized)
                return
            }

            // Add the user ID to the request context
            ctx := context.WithValue(r.Context(), contextKey.UserIDKey, uid)
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