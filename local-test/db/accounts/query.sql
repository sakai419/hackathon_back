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

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = ?;

-- name: SearchAccountsByUserName :many
SELECT * FROM accounts
WHERE user_name LIKE ?
ORDER BY user_name
LIMIT ? OFFSET ?;

-- name: GetAccountCreationDate :one
SELECT created_at FROM accounts
WHERE id = ?;

-- name: CountAccounts :one
SELECT COUNT(*) FROM accounts;

SELECT EXISTS(SELECT 1 FROM accounts WHERE user_name = ?);

-- name: CheckUserIdExists :one
SELECT EXISTS(SELECT 1 FROM accounts WHERE user_id = ?);