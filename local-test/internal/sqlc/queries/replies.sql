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

-- name: GetReplyThread :many
WITH RECURSIVE reply_thread AS (
    SELECT * FROM replies r0 WHERE r0.reply_id = $1
    UNION ALL
    SELECT r.* FROM replies r
    JOIN reply_thread rt ON r.parent_reply_id = rt.reply_id
)
SELECT rt.*, t.*
FROM reply_thread rt
JOIN tweets t ON rt.tweet_id = t.id
ORDER BY rt.created_at ASC;