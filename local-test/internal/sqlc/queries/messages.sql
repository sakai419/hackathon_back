-- name: CreateMessage :exec
INSERT INTO messages (sender_account_id, recipient_account_id, content)
VALUES ($1, $2, $3);

-- name: GetMessageById :one
SELECT * FROM messages
WHERE id = $1;

-- name: GetMessagesBetweenUsers :many
SELECT * FROM messages
WHERE (sender_account_id = $1 AND recipient_account_id = $2)
    OR (sender_account_id = $3 AND recipient_account_id = $4)
ORDER BY created_at DESC
LIMIT $5 OFFSET $6;

-- name: GetUnreadMessagesForUser :many
SELECT * FROM messages
WHERE recipient_account_id = $1 AND is_read = FALSE
ORDER BY created_at DESC;

-- name: MarkMessageAsRead :exec
UPDATE messages
SET is_read = TRUE
WHERE id = $1 AND recipient_account_id = $2;

-- name: MarkAllMessagesAsRead :exec
UPDATE messages
SET is_read = TRUE
WHERE recipient_account_id = $1 AND is_read = FALSE;

-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id = $1 AND (sender_account_id = $2 OR recipient_account_id = $3);

-- name: GetMessageCountForUser :one
SELECT COUNT(*) FROM messages
WHERE sender_account_id = $1 OR recipient_account_id = $2;

-- name: GetUnreadMessageCountForUser :one
SELECT COUNT(*) FROM messages
WHERE recipient_account_id = $1 AND is_read = FALSE;

-- name: GerUnreadMessageCountBetweenUsers :one
SELECT COUNT(*) FROM messages
WHERE recipient_account_id = $1 AND sender_account_id = $2 AND is_read = FALSE;

-- name: GetLatestMessageForEachConversation :many
SELECT m.*
FROM messages m
INNER JOIN (
    SELECT
        CASE
            WHEN sender_account_id < recipient_account_id
            THEN sender_account_id
            ELSE recipient_account_id
        END AS user1,
        CASE
            WHEN sender_account_id < recipient_account_id
            THEN recipient_account_id
            ELSE sender_account_id
        END AS user2,
        MAX(created_at) AS max_created_at
    FROM messages m2
    WHERE m2.sender_account_id = $1 OR m2.recipient_account_id = $2
    GROUP BY user1, user2
) latest ON (
    (m.sender_account_id = latest.user1 AND m.recipient_account_id = latest.user2) OR
    (m.sender_account_id = latest.user2 AND m.recipient_account_id = latest.user1)
) AND m.created_at = latest.max_created_at
ORDER BY m.created_at DESC;

-- name: SearchMessages :many
SELECT * FROM messages
WHERE (sender_account_id = $1 OR recipient_account_id = $2)
    AND content LIKE $3
ORDER BY created_at DESC
LIMIT $4 OFFSET $5;

-- name: DeleteOldMessages :exec
DELETE FROM messages
WHERE created_at < $1 AND is_read = TRUE;