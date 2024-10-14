-- name: CreateFollow :exec
INSERT INTO follows (follower_account_id, following_account_id, status)
VALUES ($1, $2, 'accepted');

-- name: CreateFollowRequest :exec
INSERT INTO follows (follower_account_id, following_account_id, status)
VALUES ($1, $2, 'pending');

-- name: AcceptFollowRequest :exec
UPDATE follows
SET status = 'accepted'
WHERE follower_account_id = $1 AND following_account_id = $2 AND status = 'pending';

-- name: DeleteFollow :execresult
DELETE FROM follows
WHERE follower_account_id = $1 AND following_account_id = $2 AND status = 'accepted';

-- name: DeleteFollowRequest :exec
DELETE FROM follows
WHERE follower_account_id = $1 AND following_account_id = $2 AND status = 'pending';

-- name: GetFollowStatus :one
SELECT status
FROM follows
WHERE follower_account_id = $1 AND following_account_id = $2;

-- name: ListPendingRequests :many
SELECT follower_account_id, following_account_id, created_at
FROM follows
WHERE following_account_id = $1 AND status = 'pending';

-- name: GetFollowerAccountIDs :many
SELECT follower_account_id
FROM follows
WHERE following_account_id = $1 AND status = 'accepted'
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetFollowingAccountIDs :many
SELECT following_account_id
FROM follows
WHERE follower_account_id = $1 AND status = 'accepted'
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;