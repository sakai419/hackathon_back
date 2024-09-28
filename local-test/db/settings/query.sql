-- name: CreateSettings :exec
INSERT INTO settings (account_id, is_private)
VALUES (?, ?);

-- name: GetSettingsByAccountId :one
SELECT * FROM settings
WHERE account_id = ?;

-- name: UpdateSettingsPrivacy :exec
UPDATE settings
SET is_private = ?
WHERE account_id = ?;

-- name: DeleteSettings :exec
DELETE FROM settings
WHERE account_id = ?;

-- name: CheckSettingsExist :one
SELECT EXISTS(SELECT 1 FROM settings WHERE account_id = ?);