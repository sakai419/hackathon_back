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
	IsAdmin ctxKey = "is_admin"
	IsSuspended ctxKey = "is_suspended"
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

func GetIsAdmin(ctx context.Context) (bool, error) {
	isAdmin, ok := ctx.Value(IsAdmin).(bool)
	if !ok {
		return false, errors.New("is_admin not found in context")
	}
	return isAdmin, nil
}

func GetIsSuspended(ctx context.Context) (bool, error) {
	isSuspended, ok := ctx.Value(IsSuspended).(bool)
	if !ok {
		return false, errors.New("is_suspended not found in context")
	}
	return isSuspended, nil
}