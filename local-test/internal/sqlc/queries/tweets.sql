-- name: CreateTweet :one
INSERT INTO tweets (
    account_id, content, code, media
) VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: CreateTweetAsReply :one
INSERT INTO tweets (
    account_id, is_reply, content, code, media
) VALUES ($1, TRUE, $2, $3, $4)
RETURNING id;

-- name: CreateTweetAsQuote :one
INSERT INTO tweets (
    account_id, is_quote, content, code, media
) VALUES ($1, TRUE, $2, $3, $4)
RETURNING id;

-- name: GetAccountIDByTweetID :one
SELECT account_id FROM tweets WHERE id = $1;

-- name: GetRecentTweetMetadatas :many
SELECT
    t.id,
    t.account_id,
    t.likes_count,
    t.retweets_count,
    t.replies_count,
    l.label1,
    l.label2,
    l.label3
FROM tweets AS t
INNER JOIN labels AS l ON t.id = l.tweet_id
INNER JOIN settings AS s ON t.account_id = s.account_id
INNER JOIN accounts AS a ON t.account_id = a.id
WHERE s.is_private = FALSE AND a.is_suspended = FALSE AND a.id != @client_account_id
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetTweetInfosByAccountID :many
SELECT
    t.*,
    COALESCE(l.has_liked, FALSE) AS has_liked,
    COALESCE(r.has_retweeted, FALSE) AS has_retweeted
FROM tweets AS t
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_liked
    FROM likes
    WHERE liking_account_id = @client_account_id
) AS l ON t.id = l.original_tweet_id
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_retweeted
    FROM retweets
    WHERE retweeting_account_id = @client_account_id
) AS r ON t.id = r.original_tweet_id
WHERE t.account_id = @target_account_id
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetTweetInfosByIDs :many
SELECT
    t.*,
    COALESCE(l.has_liked, FALSE) AS has_liked,
    COALESCE(r.has_retweeted, FALSE) AS has_retweeted
FROM tweets AS t
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_liked
    FROM likes
    WHERE liking_account_id = @client_account_id
) AS l ON t.id = l.original_tweet_id
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_retweeted
    FROM retweets
    WHERE retweeting_account_id = @client_account_id
) AS r ON t.id = r.original_tweet_id
WHERE t.id = ANY(@tweet_ids::BIGINT[]);

-- name: SearchTweetsOrderByCreatedAt :many
SELECT
    t.*,
    COALESCE(l.has_liked, FALSE) AS has_liked,
    COALESCE(r.has_retweeted, FALSE) AS has_retweeted
FROM tweets AS t
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_liked
    FROM likes
    WHERE liking_account_id = @client_account_id
) AS l ON t.id = l.original_tweet_id
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_retweeted
    FROM retweets
    WHERE retweeting_account_id = @client_account_id
) AS r ON t.id = r.original_tweet_id
WHERE
    t.content ILIKE CONCAT('%', @keyword::VARCHAR, '%')
    OR t.code->>'Content' ILIKE CONCAT('%', @keyword::VARCHAR, '%')
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2;

-- name: SearchTweetsOrderByEngagementScore :many
SELECT
    t.*,
    COALESCE(l.has_liked, FALSE) AS has_liked,
    COALESCE(r.has_retweeted, FALSE) AS has_retweeted
FROM tweets AS t
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_liked
    FROM likes
    WHERE liking_account_id = @client_account_id
) AS l ON t.id = l.original_tweet_id
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_retweeted
    FROM retweets
    WHERE retweeting_account_id = @client_account_id
) AS r ON t.id = r.original_tweet_id
WHERE
    t.content ILIKE CONCAT('%', @keyword::VARCHAR, '%')
    OR t.code->>'Content' ILIKE CONCAT('%', @keyword::VARCHAR, '%')
ORDER BY
    (t.likes_count * 30 + t.retweets_count * 20 + t.replies_count * 1) DESC,
    t.created_at DESC
LIMIT $1 OFFSET $2;


-- name: GetTweetCountByAccountID :one
SELECT COUNT(*) FROM tweets
WHERE account_id = $1;

-- name: DeleteTweet :execresult
DELETE FROM tweets WHERE id = $1;

-- name: GetPinnedTweetForAccount :one
SELECT * FROM tweets
WHERE account_id = $1 AND is_pinned = TRUE
LIMIT 1;

-- name: SetTweetAsPinned :execresult
UPDATE tweets
SET is_pinned = TRUE
WHERE id = $1 AND account_id = $2;

-- name: UnsetTweetAsPinned :execresult
UPDATE tweets
SET is_pinned = FALSE
WHERE id = $1 AND account_id = $2;

-- name: GetPinnedTweetID :one
SELECT id FROM tweets
WHERE account_id = $1 AND is_pinned = TRUE
LIMIT 1;

-- name: SearchTweets :many
SELECT * FROM tweets
WHERE content LIKE $1 OR code LIKE $2
ORDER BY created_at DESC
LIMIT $3 OFFSET $4;