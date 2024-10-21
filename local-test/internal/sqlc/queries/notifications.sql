-- name: CreateNotification :exec
INSERT INTO notifications (sender_account_id, recipient_account_id, type, content)
VALUES ($1, $2, $3, $4);

-- name: GetNotifications :many
SELECT * FROM notifications
WHERE recipient_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetUnreadNotifications :many
SELECT * FROM notifications
WHERE recipient_account_id = $1 AND is_read = FALSE
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: MarkNotificationAsRead :execresult
UPDATE notifications
SET is_read = TRUE
WHERE id = $1 AND recipient_account_id = $2;

-- name: MarkAllNotificationsAsRead :exec
UPDATE notifications
SET is_read = TRUE
WHERE recipient_account_id = $1 AND is_read = FALSE;

-- name: DeleteNotification :execresult
DELETE FROM notifications
WHERE id = $1 AND recipient_account_id = $2;

-- name: DeleteAllNotificationsForRecipient :exec
DELETE FROM notifications
WHERE recipient_account_id = $1;

-- name: DeleteAllNotificationsFromSender :exec
DELETE FROM notifications
WHERE sender_account_id = $1 AND recipient_account_id = $2;

-- name: GetNotificationCount :one
SELECT COUNT(*) FROM notifications
WHERE recipient_account_id = $1;

-- name: GetUnreadNotificationCount :one
SELECT COUNT(*) FROM notifications
WHERE recipient_account_id = $1 AND is_read = FALSE;

-- name: DeleteOldReadNotifications :exec
DELETE FROM notifications
WHERE recipient_account_id = $1 AND is_read = TRUE AND created_at < $2;