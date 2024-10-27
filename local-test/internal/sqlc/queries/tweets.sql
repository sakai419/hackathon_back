-- name: CreateTweet :one
INSERT INTO tweets (
    account_id, content, code, media
) VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: CreateTweetAsRetweet :one
INSERT INTO tweets (
    account_id, original_tweet_id, is_retweet
) VALUES ($1, $2, TRUE)
RETURNING id;

-- name: CreateTweetAsReply :exec
INSERT INTO tweets (
    account_id, is_reply, content, code, media
) VALUES ($1, TRUE, $2, $3, $4);

-- name: CreateTweetAsQuote :exec
INSERT INTO tweets (
    account_id, is_quote, content, code, media
) VALUES ($1, TRUE, $2, $3, $4);

-- name: GetTweetByID :one
SELECT * FROM tweets WHERE id = $1;

-- name: GetTweetsByAccountID :many
SELECT * FROM tweets
WHERE account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateTweetContent :exec
UPDATE tweets
SET content = $1
WHERE id = $2 AND account_id = $3;

-- name: UpdateTweetCode :exec
UPDATE tweets
SET code = $1
WHERE id = $2 AND account_id = $3;

-- name: DeleteTweet :exec
DELETE FROM tweets WHERE id = $1 AND account_id = $2;

-- name: IncrementLikesCount :exec
UPDATE tweets SET likes_count = likes_count + 1 WHERE id = $1;

-- name: IncrementRepliesCount :exec
UPDATE tweets SET replies_count = replies_count + 1 WHERE id = $1;

-- name: IncrementRetweetsCount :exec
UPDATE tweets SET retweets_count = retweets_count + 1 WHERE id = $1;

-- name: UpdateEngagementScore :exec
UPDATE tweets
SET engagement_score = likes_count + replies_count + retweets_count
WHERE id = $1;

-- name: GetTrendingTweets :many
SELECT * FROM tweets
ORDER BY engagement_score DESC
LIMIT $1 OFFSET $2;

-- name: GetPinnedTweetForAccount :one
SELECT * FROM tweets
WHERE account_id = $1 AND is_pinned = TRUE
LIMIT 1;

-- name: SetTweetAsPinned :exec
UPDATE tweets
SET is_pinned = TRUE
WHERE id = $1 AND account_id = $2;

-- name: UnpinTweet :exec
UPDATE tweets
SET is_pinned = FALSE
WHERE id = $1 AND account_id = $2;

-- name: SearchTweets :many
SELECT * FROM tweets
WHERE content LIKE $1 OR code LIKE $2
ORDER BY created_at DESC
LIMIT $3 OFFSET $4;

-- name: GetTweetCountByAccountID :one
SELECT COUNT(*) FROM tweets WHERE account_id = $1;