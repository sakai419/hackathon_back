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

-- name: CheckIsFollowed :one
SELECT EXISTS(
    SELECT 1
    FROM follows
    WHERE follower_account_id = $1 AND following_account_id = $2 AND status = 'accepted'
);

-- name: CheckMultipleFollowStatus :many
SELECT
    account_id,
    EXISTS(
        SELECT 1
        FROM follows as f1
        WHERE f1.follower_account_id = @client_account_id
          AND f1.following_account_id = account_id
          AND f1.status = 'accepted'
    ) AS is_following,
    EXISTS(
        SELECT 1
        FROM follows as f2
        WHERE f2.follower_account_id = account_id
          AND f2.following_account_id = @client_account_id
          AND f2.status = 'accepted'
    ) AS is_followed
FROM UNNEST(@account_ids::VARCHAR[]) AS account_id;



-- name: IsPrivateAndNotFollowing :one
SELECT
    CASE
        WHEN s.is_private = TRUE AND f.follower_account_id IS NULL THEN TRUE
        ELSE FALSE
    END AS is_private_and_not_following
FROM
    settings AS s
LEFT JOIN
    follows AS f
ON
    s.account_id = f.following_account_id
    AND f.follower_account_id = @client_account_id
WHERE
    s.account_id = @target_account_id;
