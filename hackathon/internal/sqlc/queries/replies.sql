-- name: CreateReply :exec
INSERT INTO replies (reply_id, original_tweet_id, parent_reply_id, replying_account_id)
SELECT
    @reply_id AS reply_id,
    COALESCE(
        (SELECT r.original_tweet_id FROM replies AS r WHERE r.reply_id = @original_tweet_id),
        @original_tweet_id
    ) AS original_tweet_id,
    COALESCE(
        (SELECT r.reply_id FROM replies AS r WHERE r.reply_id = @original_tweet_id),
        NULL
    ) AS parent_reply_id,
    @replying_account_id AS replying_account_id;

-- name: GetReplyRelations :many
SELECT reply_id, original_tweet_id, parent_reply_id
FROM replies
WHERE reply_id = ANY(@tweet_ids::BIGINT[]);

-- name: GetReplyIDs :many
SELECT reply_id
FROM replies
WHERE
original_tweet_id = $1 AND parent_reply_id IS NULL
OR parent_reply_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CheckParentReplyExist :one
SELECT EXISTS(SELECT 1 FROM replies WHERE reply_id = $1 AND parent_reply_id IS NOT NULL);