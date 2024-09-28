-- name: CreateTweet :exec
INSERT INTO tweets (
    account_id, content, code, media
) VALUES (?, ?, ?, ?);

-- name: CreateRetweet :exec
INSERT INTO tweets (
    account_id, is_retweet
) VALUES (?, TRUE);

-- name: CreateReply :exec

-- name: GetTweetById :one
SELECT * FROM tweets WHERE id = ?;

-- name: GetTweetsByAccountId :many
SELECT * FROM tweets
WHERE account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: UpdateTweetContent :exec
UPDATE tweets
SET content = ?
WHERE id = ? AND account_id = ?;

-- name: UpdateTweetCode :exec
UPDATE tweets
SET code = ?
WHERE id = ? AND account_id = ?;

-- name: DeleteTweet :exec
DELETE FROM tweets WHERE id = ? AND account_id = ?;

-- name: IncrementLikesCount :exec
UPDATE tweets SET likes_count = likes_count + 1 WHERE id = ?;

-- name: IncrementRepliesCount :exec
UPDATE tweets SET replies_count = replies_count + 1 WHERE id = ?;

-- name: IncrementRetweetsCount :exec
UPDATE tweets SET retweets_count = retweets_count + 1 WHERE id = ?;

-- name: UpdateEngagementScore :exec
UPDATE tweets
SET engagement_score = likes_count + replies_count + retweets_count
WHERE id = ?;

-- name: GetTrendingTweets :many
SELECT * FROM tweets
ORDER BY engagement_score DESC
LIMIT ?;

-- name: GetPinnedTweetForAccount :one
SELECT * FROM tweets
WHERE account_id = ? AND is_pinned = TRUE
LIMIT 1;

-- name: SetTweetAsPinned :exec
UPDATE tweets
SET is_pinned = TRUE
WHERE id = ? AND account_id = ?;

-- name: UnpinTweet :exec
UPDATE tweets
SET is_pinned = FALSE
WHERE id = ? AND account_id = ?;

-- name: SearchTweets :many
SELECT * FROM tweets
WHERE content LIKE ? OR code LIKE ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetTweetCountByAccountId :one
SELECT COUNT(*) FROM tweets WHERE account_id = ?;