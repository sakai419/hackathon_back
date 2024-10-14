-- name: CreateFollowRequest :exec
INSERT INTO follow_requests (requester_account_id, requested_account_id)
VALUES (?, ?);

-- name: DeleteFollowRequest :exec
DELETE FROM follow_requests
WHERE requester_account_id = ? AND requested_account_id = ?;

-- name: GetPendingFollowRequests :many
SELECT requester_account_id
FROM follow_requests
WHERE requested_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetSentFollowRequests :many
SELECT requested_account_id
FROM follow_requests
WHERE requester_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetPendingFollowRequestCount :one
SELECT COUNT(*) FROM follow_requests
WHERE requested_account_id = ?;

-- name: GetSentFollowRequestCount :one
SELECT COUNT(*) FROM follow_requests
WHERE requester_account_id = ?;

-- name: DeleteOldFollowRequests :exec
DELETE FROM follow_requests
WHERE created_at < ? AND requested_account_id = ?;

-- name: AcceptFollowRequest :exec
INSERT INTO follows (follower_account_id, following_account_id)
SELECT requester_account_id, requested_account_id
FROM follow_requests
WHERE requester_account_id = ? AND requested_account_id = ?;

-- name: RejectFollowRequest :exec
DELETE FROM follow_requests
WHERE requester_account_id = ? AND requested_account_id = ?;