-- name: CreateQuote :exec
INSERT INTO quotes (quote_id, quoting_account_id, original_tweet_id)
VALUES ($1, $2, $3);

-- name: GetQuotingAccountIDs :many
SELECT quoting_account_id
FROM quotes
WHERE original_tweet_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;