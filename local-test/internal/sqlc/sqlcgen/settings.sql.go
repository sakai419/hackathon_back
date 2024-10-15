// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: settings.sql

package sqlcgen

import (
	"context"
)

const checkSettingsExist = `-- name: CheckSettingsExist :one
SELECT EXISTS(SELECT 1 FROM settings WHERE account_id = $1)
`

func (q *Queries) CheckSettingsExist(ctx context.Context, accountID string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkSettingsExist, accountID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createSettingsWithDefaultValues = `-- name: CreateSettingsWithDefaultValues :exec
INSERT INTO settings (account_id)
VALUES ($1)
`

func (q *Queries) CreateSettingsWithDefaultValues(ctx context.Context, accountID string) error {
	_, err := q.db.ExecContext(ctx, createSettingsWithDefaultValues, accountID)
	return err
}

const deleteSettings = `-- name: DeleteSettings :exec
DELETE FROM settings
WHERE account_id = $1
`

func (q *Queries) DeleteSettings(ctx context.Context, accountID string) error {
	_, err := q.db.ExecContext(ctx, deleteSettings, accountID)
	return err
}

const getSettingsByAccountID = `-- name: GetSettingsByAccountID :one
SELECT account_id, is_private, created_at, updated_at FROM settings
WHERE account_id = $1
`

func (q *Queries) GetSettingsByAccountID(ctx context.Context, accountID string) (Setting, error) {
	row := q.db.QueryRowContext(ctx, getSettingsByAccountID, accountID)
	var i Setting
	err := row.Scan(
		&i.AccountID,
		&i.IsPrivate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateSettingsPrivacy = `-- name: UpdateSettingsPrivacy :exec
UPDATE settings
SET is_private = $1
WHERE account_id = $2
`

type UpdateSettingsPrivacyParams struct {
	IsPrivate bool
	AccountID string
}

func (q *Queries) UpdateSettingsPrivacy(ctx context.Context, arg UpdateSettingsPrivacyParams) error {
	_, err := q.db.ExecContext(ctx, updateSettingsPrivacy, arg.IsPrivate, arg.AccountID)
	return err
}
