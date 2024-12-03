-- name: CreateBlock :exec
INSERT INTO blocks (blocker_account_id, blocked_account_id)
VALUES ($1, $2);

-- name: DeleteBlock :execresult
DELETE FROM blocks
WHERE blocker_account_id = $1 AND blocked_account_id = $2;

-- name: CheckBlockExists :one
SELECT EXISTS(
    SELECT 1 FROM blocks
    WHERE blocker_account_id = $1 AND blocked_account_id = $2
) AS is_blocked;

-- name: GetBlockedAccountIDs :many
SELECT blocked_account_id
FROM blocks
WHERE blocker_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetBlockerAccountIDs :many
SELECT blocker_account_id
FROM blocks
WHERE blocked_account_id = @client_account_id
AND blocker_account_id = ANY(@ids::VARCHAR[]);

-- name: GetBlockCount :one
SELECT COUNT(*) FROM blocks
WHERE blocker_account_id = $1;