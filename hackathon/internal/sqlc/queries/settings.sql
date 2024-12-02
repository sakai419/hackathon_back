-- name: CreateSettingsWithDefaultValues :exec
INSERT INTO settings (account_id)
VALUES ($1);

-- name: UpdateSettings :execresult
UPDATE settings
SET is_private = COALESCE($1, is_private)
WHERE account_id = $2;