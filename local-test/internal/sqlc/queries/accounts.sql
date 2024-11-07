-- name: CreateAccount :exec
INSERT INTO accounts (id, user_id, user_name)
VALUES ($1, $2, $3);

-- name: GetAccountIDByUserID :one
SELECT id FROM accounts
WHERE user_id = $1;

-- name: GetUserInfo :one
SELECT a.id, a.user_id, a.user_name, a.is_admin, p.bio, p.profile_image_url, s.is_private
FROM accounts a
JOIN profiles p ON a.id = p.account_id
JOIN settings s ON a.id = s.account_id
WHERE a.id = $1;

-- name: GetUserInfos :many
SELECT a.id, a.user_id, a.user_name, p.bio, p.profile_image_url
FROM accounts a
JOIN profiles p ON a.id = p.account_id
WHERE a.id = ANY(@IDs::VARCHAR[]) and a.is_suspended = FALSE
ORDER BY a.created_at DESC;

-- name: UpdateAccountInfos :execresult
UPDATE accounts
SET user_id = COALESCE(NULLIF(@user_id, ''), user_id),
    user_name = COALESCE(NULLIF(@user_name, ''), user_name)
WHERE id = @id;

-- name: SuspendAccount :execresult
UPDATE accounts
SET is_suspended = TRUE
WHERE id = $1;

-- name: UnsuspendAccount :execresult
UPDATE accounts
SET is_suspended = FALSE
WHERE id = $1;

-- name: DeleteAccount :execresult
DELETE FROM accounts
WHERE id = $1;

-- name: SearchAccountsByUserID :many
SELECT * FROM accounts
WHERE user_id LIKE CONCAT('%', $1, '%')
ORDER BY user_id
LIMIT $2 OFFSET $3;

-- name: SearchAccountsByUserName :many
SELECT * FROM accounts
WHERE user_name LIKE CONCAT('%', $1, '%')
ORDER BY user_name
LIMIT $2 OFFSET $3;

-- name: IsAdmin :one
SELECT is_admin FROM accounts
WHERE id = $1;

-- name: IsSuspended :one
SELECT is_suspended FROM accounts
WHERE id = $1;

-- name: GetAccountInfo :one
SELECT a.is_suspended, a.is_admin, s.is_private
FROM accounts a
JOIN settings s ON a.id = s.account_id
WHERE a.id = $1;