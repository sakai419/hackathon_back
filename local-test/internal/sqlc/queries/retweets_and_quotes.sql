-- name: CreateRetweetOrQuote :exec
INSERT INTO retweets_and_quotes (retweet_id, retweeting_account_id, original_tweet_id)
VALUES ($1, $2, $3);

-- name: GetRetweetOrQuoteByID :one
SELECT * FROM retweets_and_quotes
WHERE retweet_id = $1;

-- name: DeleteRetweetOrQuote :exec
DELETE FROM retweets_and_quotes
WHERE retweet_id = $1;

-- name: GetRetweetsAndQuotesByOriginalTweetID :many
SELECT r.*, t.content AS retweet_content
FROM retweets_and_quotes r
JOIN tweets t ON r.tweet_id = t.id
WHERE r.original_tweet_id = $1
ORDER BY r.created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetRetweetsAndQuotesByAccountID :many
SELECT r.*, t.content AS original_tweet_content
FROM retweets_and_quotes r
JOIN tweets t ON r.original_tweet_id = t.id
WHERE r.retweeting_account_id = $1
ORDER BY r.created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetRetweetAndQuoteCount :one
SELECT COUNT(*) FROM retweets_and_quotes
WHERE original_tweet_id = $1;