package key

import (
	"context"
	"errors"
)

type ctxKey string

const (
	ClientAccountID ctxKey = "client_account_id"
	TargetAccountID ctxKey = "account_id"
	IsClientAdmin ctxKey = "is_client_admin"
	IsTargetPrivate ctxKey = "is_target_private"
	IsTargetSuspended ctxKey = "is_target_suspended"
	IsClientSuspended ctxKey = "is_client_suspended"
)

func GetClientAccountID(ctx context.Context) (string, error) {
	id, ok := ctx.Value(ClientAccountID).(string)
	if !ok {
		return "", errors.New("account_id not found in context")
	}
	return id, nil
}

func GetTargetAccountID(ctx context.Context) (string, error) {
	id, ok := ctx.Value(TargetAccountID).(string)
	if !ok {
		return "", errors.New("account_id not found in context")
	}
	return id, nil
}

func GetIsClientAdmin(ctx context.Context) (bool, error) {
	isAdmin, ok := ctx.Value(IsClientAdmin).(bool)
	if !ok {
		return false, errors.New("is_admin not found in context")
	}
	return isAdmin, nil
}

func GetIsTargetPrivate(ctx context.Context) (bool, error) {
	isPrivate, ok := ctx.Value(IsTargetPrivate).(bool)
	if !ok {
		return false, errors.New("is_private not found in context")
	}
	return isPrivate, nil
}

func GetIsTargetSuspended(ctx context.Context) (bool, error) {
	isSuspended, ok := ctx.Value(IsTargetSuspended).(bool)
	if !ok {
		return false, errors.New("is_suspended not found in context")
	}
	return isSuspended, nil
}

func GetIsClientSuspended(ctx context.Context) (bool, error) {
	isSuspended, ok := ctx.Value(IsClientSuspended).(bool)
	if !ok {
		return false, errors.New("is_suspended not found in context")
	}
	return isSuspended, nil
}