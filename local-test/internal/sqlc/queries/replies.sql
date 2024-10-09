-- name: CreateReply :exec
INSERT INTO replies (original_tweet_id, parent_reply_id, replying_account_id)
VALUES (?, ?, ?);

-- name: GetReplyById :one
SELECT * FROM replies WHERE id = ?;

-- name: DeleteReply :exec
DELETE FROM replies WHERE id = ?;

-- -- name: GetRepliesByOriginalTweetId :many
-- SELECT r.*, t.content AS reply_content, a.user_name AS replier_name
-- FROM replies r
-- JOIN tweets t ON r.reply_id = t.id
-- JOIN accounts a ON r.replying_account_id = a.id
-- WHERE r.original_tweet_id = ?
-- ORDER BY r.created_at ASC
-- LIMIT ? OFFSET ?;

-- -- name: GetRepliesByParentReplyId :many
-- SELECT r.*, t.content AS reply_content, a.user_name AS replier_name
-- FROM replies r
-- JOIN tweets t ON r.reply_id = t.id
-- JOIN accounts a ON r.replying_account_id = a.id
-- WHERE r.parent_reply_id = ?
-- ORDER BY r.created_at ASC
-- LIMIT ? OFFSET ?;

-- -- name: GetRepliesByAccountId :many
-- SELECT r.*, t.content AS reply_content, ot.content AS original_tweet_content
-- FROM replies r
-- JOIN tweets t ON r.reply_id = t.id
-- JOIN tweets ot ON r.original_tweet_id = ot.id
-- WHERE r.replying_account_id = ?
-- ORDER BY r.created_at DESC
-- LIMIT ? OFFSET ?;

-- name: GetReplyCount :one
SELECT COUNT(*) FROM replies WHERE original_tweet_id = ?;

-- name: GetReplyThread :many
WITH RECURSIVE reply_thread AS (
    SELECT * FROM replies r0 WHERE r0.id = ?
    UNION ALL
    SELECT r.* FROM replies r
    JOIN reply_thread rt ON r.parent_reply_id = rt.id
)
SELECT rt.*, t.*
FROM reply_thread rt
JOIN tweets t ON rt.reply_id = t.id
ORDER BY rt.created_at ASC;