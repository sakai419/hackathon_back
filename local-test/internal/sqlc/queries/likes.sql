-- name: CreateLike :exec
INSERT INTO likes (liking_account_id, original_tweet_id)
VALUES ($1, $2);

-- name: DeleteLike :exec
DELETE FROM likes
WHERE liking_account_id = $1 AND original_tweet_id = $2;

-- name: GetLikesByTweetId :many
SELECT * FROM likes
WHERE original_tweet_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetLikesByAccountId :many
SELECT * FROM likes
WHERE liking_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetLikeCount :one
SELECT COUNT(*) FROM likes
WHERE original_tweet_id = $1;