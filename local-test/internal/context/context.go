package contextKey

import (
	"context"
	"errors"
)

type ctxKey string

const (
	UserIDKey ctxKey = "user_id"
)

func GetUserID(ctx context.Context) (string, error) {
	id, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		return "", errors.New("user_id not found in context")
	}
	return id, nil
}