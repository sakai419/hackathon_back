-- name: CreateRetweetOrQuote :exec
INSERT INTO retweets_and_quotes (retweeting_account_id, original_tweet_id)
VALUES (?, ?);

-- name: GetRetweetOrQuoteById :one
SELECT * FROM retweets_and_quotes
WHERE id = ?;

-- name: DeleteRetweetOrQuote :exec
DELETE FROM retweets_and_quotes
WHERE id = ?;

-- name: GetRetweetsAndQuotesByOriginalTweetId :many
SELECT r.*, t.content AS retweet_content
FROM retweets_and_quotes r
JOIN tweets t ON r.id = t.id
WHERE r.original_tweet_id = ?
ORDER BY r.created_at DESC
LIMIT ? OFFSET ?;

-- name: GetRetweetsAndQuotesByAccountId :many
SELECT r.*, t.content AS original_tweet_content
FROM retweets_and_quotes r
JOIN tweets t ON r.original_tweet_id = t.id
WHERE r.retweeting_account_id = ?
ORDER BY r.created_at DESC
LIMIT ? OFFSET ?;

-- name: GetRetweetAndQuoteCount :one
SELECT COUNT(*) FROM retweets_and_quotes
WHERE original_tweet_id = ?;