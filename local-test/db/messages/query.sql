-- name: CreateMessage :exec
INSERT INTO messages (sender_account_id, recipient_account_id, content)
VALUES (?, ?, ?);

-- name: GetMessageById :one
SELECT * FROM messages
WHERE id = ?;

-- name: GetMessagesBetweenUsers :many
SELECT * FROM messages
WHERE (sender_account_id = ? AND recipient_account_id = ?)
    OR (sender_account_id = ? AND recipient_account_id = ?)
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetUnreadMessagesForUser :many
SELECT * FROM messages
WHERE recipient_account_id = ? AND is_read = FALSE
ORDER BY created_at DESC;

-- name: MarkMessageAsRead :exec
UPDATE messages
SET is_read = TRUE
WHERE id = ? AND recipient_account_id = ?;

-- name: MarkAllMessagesAsRead :exec
UPDATE messages
SET is_read = TRUE
WHERE recipient_account_id = ? AND is_read = FALSE;

-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id = ? AND (sender_account_id = ? OR recipient_account_id = ?);

-- name: GetMessageCountForUser :one
SELECT COUNT(*) FROM messages
WHERE sender_account_id = ? OR recipient_account_id = ?;

-- name: GetUnreadMessageCountForUser :one
SELECT COUNT(*) FROM messages
WHERE recipient_account_id = ? AND is_read = FALSE;

-- name: GerUnreadMessageCountBetweenUsers :one
SELECT COUNT(*) FROM messages
WHERE recipient_account_id = ? AND sender_account_id = ? AND is_read = FALSE;

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
    FROM messages
    WHERE sender_account_id = ? OR recipient_account_id = ?
    GROUP BY user1, user2
) latest ON (
    (m.sender_account_id = latest.user1 AND m.recipient_account_id = latest.user2) OR
    (m.sender_account_id = latest.user2 AND m.recipient_account_id = latest.user1)
) AND m.created_at = latest.max_created_at
ORDER BY m.created_at DESC;

-- name: SearchMessages :many
SELECT * FROM messages
WHERE (sender_account_id = ? OR recipient_account_id = ?)
    AND content LIKE ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: DeleteOldMessages :exec
DELETE FROM messages
WHERE created_at < ? AND is_read = TRUE;