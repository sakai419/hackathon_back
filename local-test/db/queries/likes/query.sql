-- name: CreateLike :exec
INSERT INTO likes (liking_account_id, original_tweet_id)
VALUES (?, ?);

-- name: DeleteLike :exec
DELETE FROM likes
WHERE liking_account_id = ? AND original_tweet_id = ?;

-- name: GetLikesByTweetId :many
SELECT * FROM likes
WHERE original_tweet_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetLikesByAccountId :many
SELECT * FROM likes
WHERE liking_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetLikeCount :one
SELECT COUNT(*) FROM likes
WHERE original_tweet_id = ?;