// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: settings.sql

package sqlcgen

import (
	"context"
	"database/sql"
)

const createSettingsWithDefaultValues = `-- name: CreateSettingsWithDefaultValues :exec
INSERT INTO settings (account_id)
VALUES ($1)
`

func (q *Queries) CreateSettingsWithDefaultValues(ctx context.Context, accountID string) error {
	_, err := q.db.ExecContext(ctx, createSettingsWithDefaultValues, accountID)
	return err
}

const updateSettings = `-- name: UpdateSettings :execresult
UPDATE settings
SET is_private = COALESCE($1, is_private)
WHERE account_id = $2
`

type UpdateSettingsParams struct {
	IsPrivate sql.NullBool
	AccountID string
}

func (q *Queries) UpdateSettings(ctx context.Context, arg UpdateSettingsParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateSettings, arg.IsPrivate, arg.AccountID)
}
