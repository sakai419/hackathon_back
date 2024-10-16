package key

import (
	"context"
	"errors"
)

type ctxKey string

const (
	ClientAccountIDKey ctxKey = "client_account_id"
	PathUserID ctxKey = "user_id"
	PathAccountID ctxKey = "account_id"
)

func GetClientAccountID(ctx context.Context) (string, error) {
	id, ok := ctx.Value(ClientAccountIDKey).(string)
	if !ok {
		return "", errors.New("account_id not found in context")
	}
	return id, nil
}

func GetUserIDFromPath(ctx context.Context) (string, error) {
	id, ok := ctx.Value(PathUserID).(string)
	if !ok {
		return "", errors.New("user_id not found in context")
	}
	return id, nil
}

func GetAccountIDFromPath(ctx context.Context) (string, error) {
	id, ok := ctx.Value(PathAccountID).(string)
	if !ok {
		return "", errors.New("account_id not found in context")
	}
	return id, nil
}