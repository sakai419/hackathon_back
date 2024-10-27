-- name: CreateRetweet :exec
INSERT INTO retweets (retweeting_account_id, original_tweet_id)
VALUES ($1, $2);