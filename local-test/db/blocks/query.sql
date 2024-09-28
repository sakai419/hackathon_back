-- name: CreateBlock :exec
INSERT INTO blocks (blocker_account_id, blocked_account_id)
VALUES (?, ?);

-- name: DeleteBlock :exec
DELETE FROM blocks
WHERE blocker_account_id = ? AND blocked_account_id = ?;

-- name: CheckBlockExists :one
SELECT EXISTS(
    SELECT 1 FROM blocks
    WHERE blocker_account_id = ? AND blocked_account_id = ?
) AS is_blocked;

-- name: GetBlockedUsers :many
SELECT blocked_account_id
FROM blocks
WHERE blocker_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetBlockersOfUser :many
SELECT blocker_account_id
FROM blocks
WHERE blocked_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetBlockCount :one
SELECT COUNT(*) FROM blocks
WHERE blocker_account_id = ?;

-- name: GetBlockedByCount :one
SELECT COUNT(*) FROM blocks
WHERE blocked_account_id = ?;

-- name: DeleteAllBlocksForUser :exec
DELETE FROM blocks
WHERE blocker_account_id = ? OR blocked_account_id = ?;