-- name: CreateFollow :exec
INSERT INTO follows (follower_account_id, following_account_id)
VALUES (?, ?);

-- name: DeleteFollow :exec
DELETE FROM follows
WHERE follower_account_id = ? AND following_account_id = ?;

-- name: CheckFollowExists :one
SELECT EXISTS(
    SELECT 1 FROM follows
    WHERE follower_account_id = ? AND following_account_id = ?
) AS is_following;

-- name: GetFollowers :many
SELECT follower_account_id
FROM follows
WHERE following_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetFollowing :many
SELECT following_account_id
FROM follows
WHERE follower_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetFollowerCount :one
SELECT COUNT(*) FROM follows
WHERE following_account_id = ?;

-- name: GetFollowingCount :one
SELECT COUNT(*) FROM follows
WHERE follower_account_id = ?;

-- name: GetMutualFollows :many
SELECT f1.following_account_id
FROM follows f1
JOIN follows f2 ON f1.following_account_id = f2.follower_account_id
WHERE f1.follower_account_id = ? AND f2.following_account_id = ?
LIMIT ? OFFSET ?;

-- name: GetFollowSuggestions :many
SELECT DISTINCT f2.following_account_id
FROM follows f1
JOIN follows f2 ON f1.following_account_id = f2.follower_account_id
WHERE f1.follower_account_id = ?
    AND f2.following_account_id != f1.follower_account_id
    AND NOT EXISTS (
        SELECT 1 FROM follows f3
        WHERE f3.follower_account_id = f1.follower_account_id
            AND f3.following_account_id = f2.following_account_id
    )
LIMIT ?;