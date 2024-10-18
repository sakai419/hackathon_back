-- name: CreateAccount :exec
INSERT INTO accounts (id, user_id, user_name)
VALUES ($1, $2, $3);

-- name: GetAccountIDByUserID :one
SELECT id FROM accounts
WHERE user_id = $1;

-- name: GetUserInfos :many
SELECT a.user_id, a.user_name, p.bio, p.profile_image_url
FROM accounts a
JOIN profiles p ON a.id = p.account_id
WHERE a.id = ANY(@IDs::VARCHAR[]) and a.is_suspended = FALSE
ORDER BY a.created_at DESC;

-- name: UpdateAccountUserName :exec
UPDATE accounts
SET user_name = $1
WHERE id = $2;

-- name: UpdateAccountUserID :exec
UPDATE accounts
SET user_id = $1
WHERE id = $2;

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