-- name: CreateBlock :exec
INSERT INTO blocks (blocker_account_id, blocked_account_id)
VALUES ($1, $2);

-- name: DeleteBlock :exec
DELETE FROM blocks
WHERE blocker_account_id = $1 AND blocked_account_id = $2;

-- name: CheckBlockExists :one
SELECT EXISTS(
    SELECT 1 FROM blocks
    WHERE blocker_account_id = $1 AND blocked_account_id = $2
) AS is_blocked;

-- name: GetBlockedUsers :many
SELECT blocked_account_id
FROM blocks
WHERE blocker_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetBlockersOfUser :many
SELECT blocker_account_id
FROM blocks
WHERE blocked_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetBlockCount :one
SELECT COUNT(*) FROM blocks
WHERE blocker_account_id = $1;

-- name: GetBlockedByCount :one
SELECT COUNT(*) FROM blocks
WHERE blocked_account_id = $1;

-- name: DeleteAllBlocksForUser :exec
DELETE FROM blocks
WHERE blocker_account_id = $1 OR blocked_account_id = $1;