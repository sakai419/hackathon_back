-- name: CreateMessage :one
INSERT INTO messages (conversation_id, sender_account_id, content, is_read)
VALUES ($1, $2, $3, FALSE)
RETURNING id;

-- name: GetMessageList :many
SELECT m.id, a.user_id, m.content, m.is_read, m.created_at FROM messages AS m
INNER JOIN accounts AS a ON m.sender_account_id = a.id
WHERE m.conversation_id = $1
ORDER BY m.created_at DESC
LIMIT $2 OFFSET $3;

-- name: MarkMessageListAsRead :exec
UPDATE messages
SET is_read = TRUE
WHERE conversation_id = $1
  AND sender_account_id <> $2
  AND is_read = FALSE
  AND created_at <= NOW();