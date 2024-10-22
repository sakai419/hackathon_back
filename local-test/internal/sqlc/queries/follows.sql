-- name: CreateFollow :exec
INSERT INTO follows (follower_account_id, following_account_id, status)
VALUES ($1, $2, 'accepted');

-- name: CreateFollowRequest :exec
INSERT INTO follows (follower_account_id, following_account_id, status)
VALUES ($1, $2, 'pending');

-- name: AcceptFollowRequest :execresult
UPDATE follows
SET status = 'accepted'
WHERE follower_account_id = $1 AND following_account_id = $2 AND status = 'pending';

-- name: DeleteFollow :execresult
DELETE FROM follows
WHERE follower_account_id = $1 AND following_account_id = $2 AND status = 'accepted';

-- name: DeleteFollowRequest :execresult
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

-- name: GetFollowCounts :one
SELECT
    COUNT(CASE WHEN following_account_id = $1 AND status = 'accepted' THEN 1 END) AS follower_count,
    COUNT(CASE WHEN follower_account_id = $1 AND status = 'accepted' THEN 1 END) AS following_count
FROM follows;


-- name: GetFollowRequestCount :one
SELECT COUNT(*)
FROM follows
WHERE following_account_id = $1 AND status = 'pending';