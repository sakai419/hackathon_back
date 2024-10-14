-- name: CreateFollowRequest :exec
INSERT INTO follow_requests (requester_account_id, requested_account_id)
VALUES ($1, $2);

-- name: DeleteFollowRequest :exec
DELETE FROM follow_requests
WHERE requester_account_id = $1 AND requested_account_id = $2;

-- name: GetPendingFollowRequests :many
SELECT requester_account_id
FROM follow_requests
WHERE requested_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetSentFollowRequests :many
SELECT requested_account_id
FROM follow_requests
WHERE requester_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetPendingFollowRequestCount :one
SELECT COUNT(*) FROM follow_requests
WHERE requested_account_id = $1;

-- name: GetSentFollowRequestCount :one
SELECT COUNT(*) FROM follow_requests
WHERE requester_account_id = $1;

-- name: DeleteOldFollowRequests :exec
DELETE FROM follow_requests
WHERE created_at < $1 AND requested_account_id = $2;

-- name: AcceptFollowRequest :exec
INSERT INTO follows (follower_account_id, following_account_id)
SELECT requester_account_id, requested_account_id
FROM follow_requests
WHERE requester_account_id = $1 AND requested_account_id = $2;

-- name: RejectFollowRequest :exec
DELETE FROM follow_requests
WHERE requester_account_id = $1 AND requested_account_id = $2;