package key

import (
	"context"
	"errors"
)

type ctxKey string

const (
	AccountIDKey ctxKey = "account_id"
)

func GetAccountID(ctx context.Context) (string, error) {
	id, ok := ctx.Value(AccountIDKey).(string)
	if !ok {
		return "", errors.New("account_id not found in context")
	}
	return id, nil
}