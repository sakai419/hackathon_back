-- name: CreateAccount :exec
INSERT INTO accounts (id, user_id, user_name)
VALUES (?, ?, ?);

-- name: GetAccountById :one
SELECT * FROM accounts
WHERE id = ?;

-- name: GetAccountByUserId :one
SELECT * FROM accounts
WHERE user_id = ?;

-- name: GetAccountByUserName :one
SELECT * FROM accounts
WHERE user_name = ?;

-- name: UpdateAccountUserName :exec
UPDATE accounts
SET user_name = ?
WHERE id = ?;

-- name: UpdateAccountUserId :exec
UPDATE accounts
SET user_id = ?
WHERE id = ?;

-- name: SuspendAccount :execresult
UPDATE accounts
SET is_suspended = TRUE
WHERE id = ?;

-- name: UnsuspendAccount :execresult
UPDATE accounts
SET is_suspended = FALSE
WHERE id = ?;

-- name: DeleteAccount :execresult
DELETE FROM accounts
WHERE id = ?;

-- name: SearchAccountsByUserId :many
SELECT * FROM accounts
WHERE user_id LIKE CONCAT('%', ?, '%')
ORDER BY user_id
LIMIT ? OFFSET ?;

-- name: SearchAccountsByUserName :many
SELECT * FROM accounts
WHERE user_name LIKE CONCAT('%', ?, '%')
ORDER BY user_name
LIMIT ? OFFSET ?;

-- name: GetAccountCreationDate :one
SELECT created_at FROM accounts
WHERE id = ?;

-- name: CountAccounts :one
SELECT COUNT(*) FROM accounts;

-- name: CheckUserNameExists :one
SELECT EXISTS(SELECT 1 FROM accounts WHERE user_name = ?);

-- name: CheckUserIdExists :one
SELECT EXISTS(SELECT 1 FROM accounts WHERE user_id = ?);