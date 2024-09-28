-- name: CreateNotification :exec
INSERT INTO notifications (sender_account_id, recipient_account_id, type, content)
VALUES (?, ?, ?, ?);

-- name: GetNotificationById :one
SELECT * FROM notifications
WHERE id = ?;

-- name: GetNotificationsByRecipientId :many
SELECT * FROM notifications
WHERE recipient_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetUnreadNotificationsByRecipientId :many
SELECT * FROM notifications
WHERE recipient_account_id = ? AND is_read = FALSE
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: MarkNotificationAsRead :exec
UPDATE notifications
SET is_read = TRUE
WHERE id = ? AND recipient_account_id = ?;

-- name: MarkAllNotificationsAsRead :exec
UPDATE notifications
SET is_read = TRUE
WHERE recipient_account_id = ? AND is_read = FALSE;

-- name: DeleteNotification :exec
DELETE FROM notifications
WHERE id = ? AND recipient_account_id = ?;

-- name: DeleteAllNotificationsForRecipient :exec
DELETE FROM notifications
WHERE recipient_account_id = ?;

-- name: GetNotificationCountByRecipientId :one
SELECT COUNT(*) FROM notifications
WHERE recipient_account_id = ?;

-- name: GetUnreadNotificationCountByRecipientId :one
SELECT COUNT(*) FROM notifications
WHERE recipient_account_id = ? AND is_read = FALSE;

-- name: GetNotificationsByType :many
SELECT * FROM notifications
WHERE recipient_account_id = ? AND type = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: DeleteOldReadNotifications :exec
DELETE FROM notifications
WHERE recipient_account_id = ? AND is_read = TRUE AND created_at < ?;