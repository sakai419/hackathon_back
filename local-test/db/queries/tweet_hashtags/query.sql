-- name: CreateTweetHashtag :exec
INSERT INTO tweet_hashtags (tweet_id, hashtag_id)
VALUES (?, ?);

-- name: DeleteTweetHashtag :exec
DELETE FROM tweet_hashtags
WHERE tweet_id = ? AND hashtag_id = ?;

-- name: GetHashtagsByTweetId :many
SELECT h.*
FROM hashtags h
JOIN tweet_hashtags th ON h.id = th.hashtag_id
WHERE th.tweet_id = ?;

-- name: GetTweetsByHashtagId :many
SELECT t.*
FROM tweets t
JOIN tweet_hashtags th ON t.id = th.tweet_id
WHERE th.hashtag_id = ?
ORDER BY t.created_at DESC
LIMIT ? OFFSET ?;

-- name: CheckTweetHashtagExists :one
SELECT EXISTS(
    SELECT 1 FROM tweet_hashtags
    WHERE tweet_id = ? AND hashtag_id = ?
) AS hashtag_exists;

-- name: GetTweetCountByHashtagId :one
SELECT COUNT(DISTINCT tweet_id)
FROM tweet_hashtags
WHERE hashtag_id = ?;

-- name: GetMostUsedHashtags :many
SELECT h.*, COUNT(th.tweet_id) as usage_count
FROM hashtags h
JOIN tweet_hashtags th ON h.id = th.hashtag_id
GROUP BY h.id
ORDER BY usage_count DESC
LIMIT ?;

-- name: GetRecentTweetsWithHashtag :many
SELECT t.*, h.tag
FROM tweets t
JOIN tweet_hashtags th ON t.id = th.tweet_id
JOIN hashtags h ON th.hashtag_id = h.id
WHERE h.tag = ?
ORDER BY t.created_at DESC
LIMIT ?;

-- name: DeleteAllHashtagsForTweet :exec
DELETE FROM tweet_hashtags
WHERE tweet_id = ?;