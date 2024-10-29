-- name: CreateRetweet :exec
INSERT INTO retweets (retweeting_account_id, original_tweet_id)
VALUES ($1, $2);

-- name: DeleteRetweet :execresult
DELETE FROM retweets
WHERE retweeting_account_id = $1 AND original_tweet_id = $2;

-- name: GetRetweetingAccountIDs :many
SELECT retweeting_account_id
FROM retweets
WHERE original_tweet_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;