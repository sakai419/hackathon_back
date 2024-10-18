-- name: CreateMessage :one
INSERT INTO messages (conversation_id, sender_account_id, content, is_read)
VALUES ($1, $2, $3, FALSE)
RETURNING id;

-- name: GetMessageList :many
SELECT * FROM messages
WHERE conversation_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: MarkMessageListAsRead :exec
UPDATE messages
SET is_read = TRUE
WHERE conversation_id = $1
  AND sender_account_id <> $2
  AND is_read = FALSE
  AND created_at <= NOW();