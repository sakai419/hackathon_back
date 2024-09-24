-- name: GetAccountById :one
SELECT * FROM account WHERE id = ?;

-- name: GetAccountsByUserId :many
