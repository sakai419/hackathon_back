-- name: CreateFollow :exec
INSERT INTO follows (follower_account_id, following_account_id)
VALUES ($1, $2);

-- name: DeleteFollow :execresult
DELETE FROM follows
WHERE follower_account_id = $1 AND following_account_id = $2;

-- name: CheckFollowExists :one
SELECT EXISTS(
    SELECT 1 FROM follows
    WHERE follower_account_id = $1 AND following_account_id = $2
) AS is_following;

-- name: GetFollowerAccountIDs :many
SELECT follower_account_id
FROM follows
WHERE following_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetFollowing :many
SELECT following_account_id
FROM follows
WHERE follower_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetFollowerCount :one
SELECT COUNT(*) FROM follows
WHERE following_account_id = $1;

-- name: GetFollowingCount :one
SELECT COUNT(*) FROM follows
WHERE follower_account_id = $1;

-- name: GetMutualFollows :many
SELECT f1.following_account_id
FROM follows f1
JOIN follows f2 ON f1.following_account_id = f2.follower_account_id
WHERE f1.follower_account_id = $1 AND f2.following_account_id = $2
LIMIT $3 OFFSET $4;

-- name: GetFollowSuggestions :many
SELECT DISTINCT f2.following_account_id
FROM follows f1
JOIN follows f2 ON f1.following_account_id = f2.follower_account_id
WHERE f1.follower_account_id = $1
    AND f2.following_account_id != f1.follower_account_id
    AND NOT EXISTS (
        SELECT 1 FROM follows f3
        WHERE f3.follower_account_id = f1.follower_account_id
            AND f3.following_account_id = f2.following_account_id
    )
LIMIT $2 OFFSET $3;