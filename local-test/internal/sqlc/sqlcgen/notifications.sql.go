// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: notifications.sql

package sqlcgen

import (
	"context"
	"database/sql"
	"time"
)

const createNotification = `-- name: CreateNotification :exec
INSERT INTO notifications (sender_account_id, recipient_account_id, type, content)
VALUES ($1, $2, $3, $4)
`

type CreateNotificationParams struct {
	SenderAccountID    sql.NullString
	RecipientAccountID string
	Type               NotificationType
	Content            sql.NullString
}

func (q *Queries) CreateNotification(ctx context.Context, arg CreateNotificationParams) error {
	_, err := q.db.ExecContext(ctx, createNotification,
		arg.SenderAccountID,
		arg.RecipientAccountID,
		arg.Type,
		arg.Content,
	)
	return err
}

const deleteAllNotificationsForRecipient = `-- name: DeleteAllNotificationsForRecipient :exec
DELETE FROM notifications
WHERE recipient_account_id = $1
`

func (q *Queries) DeleteAllNotificationsForRecipient(ctx context.Context, recipientAccountID string) error {
	_, err := q.db.ExecContext(ctx, deleteAllNotificationsForRecipient, recipientAccountID)
	return err
}

const deleteNotification = `-- name: DeleteNotification :execresult
DELETE FROM notifications
WHERE id = $1 AND recipient_account_id = $2
`

type DeleteNotificationParams struct {
	ID                 int64
	RecipientAccountID string
}

func (q *Queries) DeleteNotification(ctx context.Context, arg DeleteNotificationParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteNotification, arg.ID, arg.RecipientAccountID)
}

const deleteOldReadNotifications = `-- name: DeleteOldReadNotifications :exec
DELETE FROM notifications
WHERE recipient_account_id = $1 AND is_read = TRUE AND created_at < $2
`

type DeleteOldReadNotificationsParams struct {
	RecipientAccountID string
	CreatedAt          time.Time
}

func (q *Queries) DeleteOldReadNotifications(ctx context.Context, arg DeleteOldReadNotificationsParams) error {
	_, err := q.db.ExecContext(ctx, deleteOldReadNotifications, arg.RecipientAccountID, arg.CreatedAt)
	return err
}

const getNotificationCount = `-- name: GetNotificationCount :one
SELECT COUNT(*) FROM notifications
WHERE recipient_account_id = $1
`

func (q *Queries) GetNotificationCount(ctx context.Context, recipientAccountID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getNotificationCount, recipientAccountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getNotifications = `-- name: GetNotifications :many
SELECT id, sender_account_id, recipient_account_id, type, content, tweet_id, is_read, created_at FROM notifications
WHERE recipient_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetNotificationsParams struct {
	RecipientAccountID string
	Limit              int32
	Offset             int32
}

func (q *Queries) GetNotifications(ctx context.Context, arg GetNotificationsParams) ([]Notification, error) {
	rows, err := q.db.QueryContext(ctx, getNotifications, arg.RecipientAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Notification
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.ID,
			&i.SenderAccountID,
			&i.RecipientAccountID,
			&i.Type,
			&i.Content,
			&i.TweetID,
			&i.IsRead,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnreadNotificationCount = `-- name: GetUnreadNotificationCount :one
SELECT COUNT(*) FROM notifications
WHERE recipient_account_id = $1 AND is_read = FALSE
`

func (q *Queries) GetUnreadNotificationCount(ctx context.Context, recipientAccountID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getUnreadNotificationCount, recipientAccountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUnreadNotifications = `-- name: GetUnreadNotifications :many
SELECT id, sender_account_id, recipient_account_id, type, content, tweet_id, is_read, created_at FROM notifications
WHERE recipient_account_id = $1 AND is_read = FALSE
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetUnreadNotificationsParams struct {
	RecipientAccountID string
	Limit              int32
	Offset             int32
}

func (q *Queries) GetUnreadNotifications(ctx context.Context, arg GetUnreadNotificationsParams) ([]Notification, error) {
	rows, err := q.db.QueryContext(ctx, getUnreadNotifications, arg.RecipientAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Notification
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.ID,
			&i.SenderAccountID,
			&i.RecipientAccountID,
			&i.Type,
			&i.Content,
			&i.TweetID,
			&i.IsRead,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const markAllNotificationsAsRead = `-- name: MarkAllNotificationsAsRead :exec
UPDATE notifications
SET is_read = TRUE
WHERE recipient_account_id = $1 AND is_read = FALSE
`

func (q *Queries) MarkAllNotificationsAsRead(ctx context.Context, recipientAccountID string) error {
	_, err := q.db.ExecContext(ctx, markAllNotificationsAsRead, recipientAccountID)
	return err
}

const markNotificationAsRead = `-- name: MarkNotificationAsRead :execresult
UPDATE notifications
SET is_read = TRUE
WHERE id = $1 AND recipient_account_id = $2
`

type MarkNotificationAsReadParams struct {
	ID                 int64
	RecipientAccountID string
}

func (q *Queries) MarkNotificationAsRead(ctx context.Context, arg MarkNotificationAsReadParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, markNotificationAsRead, arg.ID, arg.RecipientAccountID)
}
