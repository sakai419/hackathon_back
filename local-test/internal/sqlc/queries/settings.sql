-- name: CreateSettingsWithDefaultValues :exec
INSERT INTO settings (account_id)
VALUES ($1);

-- name: GetSettingsByAccountID :one
SELECT * FROM settings
WHERE account_id = $1;

-- name: UpdateSettingsPrivacy :exec
UPDATE settings
SET is_private = $1
WHERE account_id = $2;

-- name: DeleteSettings :exec
DELETE FROM settings
WHERE account_id = $1;

-- name: CheckSettingsExist :one
SELECT EXISTS(SELECT 1 FROM settings WHERE account_id = $1);