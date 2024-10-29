-- name: CreateQuote :exec
INSERT INTO quotes (quote_id, quoting_account_id, original_tweet_id)
VALUES ($1, $2, $3);