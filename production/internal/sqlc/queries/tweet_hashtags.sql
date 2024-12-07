-- name: AssociateTweetWithHashtags :exec
INSERT INTO tweet_hashtags (tweet_id, hashtag_id)
VALUES
    (@tweet_id, unnest(@hashtag_ids::bigint[]));

-- name: GetHashtagsByTweetID :many
SELECT h.*
FROM hashtags h
JOIN tweet_hashtags th ON h.id = th.hashtag_id
WHERE th.tweet_id = $1;

-- name: GetTweetsByHashtagID :many
SELECT t.*
FROM tweets t
JOIN tweet_hashtags th ON t.id = th.tweet_id
WHERE th.hashtag_id = $1
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2;

-- name: CheckTweetHashtagExists :one
SELECT EXISTS(
    SELECT 1 FROM tweet_hashtags
    WHERE tweet_id = $1 AND hashtag_id = $2
) AS hashtag_exists;

-- name: GetTweetCountByHashtagID :one
SELECT COUNT(DISTINCT tweet_id)
FROM tweet_hashtags
WHERE hashtag_id = $1;

-- name: GetMostUsedHashtags :many
SELECT h.*, COUNT(th.tweet_id) as usage_count
FROM hashtags h
JOIN tweet_hashtags th ON h.id = th.hashtag_id
GROUP BY h.id
ORDER BY usage_count DESC
LIMIT $1 OFFSET $2;

-- name: GetRecentTweetsWithHashtag :many
SELECT t.*, h.tag
FROM tweets t
JOIN tweet_hashtags th ON t.id = th.tweet_id
JOIN hashtags h ON th.hashtag_id = h.id
WHERE h.tag = $1
ORDER BY t.created_at DESC
LIMIT $2 OFFSET $3;

-- name: DeleteAllHashtagsForTweet :exec
DELETE FROM tweet_hashtags
WHERE tweet_id = $1;

-- name: SearchTweetsByHashtagOrderByCreatedAt :many
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
JOIN tweet_hashtags th ON t.id = th.tweet_id
JOIN hashtags h ON th.hashtag_id = h.id
WHERE h.tag = @hashtag
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2;

-- name: SearchTweetsByHashtagOrderByEngagementScore :many
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
JOIN tweet_hashtags th ON t.id = th.tweet_id
JOIN hashtags h ON th.hashtag_id = h.id
WHERE h.tag = @hashtag
ORDER BY
    (t.likes_count * 30 + t.retweets_count * 20 + t.replies_count * 1) DESC,
    t.created_at DESC
LIMIT $1 OFFSET $2;
