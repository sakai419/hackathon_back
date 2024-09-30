// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: notifications.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createNotification = `-- name: CreateNotification :exec
INSERT INTO notifications (sender_account_id, recipient_account_id, type, content)
VALUES (?, ?, ?, ?)
`

type CreateNotificationParams struct {
	SenderAccountID    sql.NullString
	RecipientAccountID string
	Type               string
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
WHERE recipient_account_id = ?
`

func (q *Queries) DeleteAllNotificationsForRecipient(ctx context.Context, recipientAccountID string) error {
	_, err := q.db.ExecContext(ctx, deleteAllNotificationsForRecipient, recipientAccountID)
	return err
}

const deleteNotification = `-- name: DeleteNotification :exec
DELETE FROM notifications
WHERE id = ? AND recipient_account_id = ?
`

type DeleteNotificationParams struct {
	ID                 uint32
	RecipientAccountID string
}

func (q *Queries) DeleteNotification(ctx context.Context, arg DeleteNotificationParams) error {
	_, err := q.db.ExecContext(ctx, deleteNotification, arg.ID, arg.RecipientAccountID)
	return err
}

const deleteOldReadNotifications = `-- name: DeleteOldReadNotifications :exec
DELETE FROM notifications
WHERE recipient_account_id = ? AND is_read = TRUE AND created_at < ?
`

type DeleteOldReadNotificationsParams struct {
	RecipientAccountID string
	CreatedAt          time.Time
}

func (q *Queries) DeleteOldReadNotifications(ctx context.Context, arg DeleteOldReadNotificationsParams) error {
	_, err := q.db.ExecContext(ctx, deleteOldReadNotifications, arg.RecipientAccountID, arg.CreatedAt)
	return err
}

const getNotificationById = `-- name: GetNotificationById :one
SELECT id, sender_account_id, recipient_account_id, type, content, is_read, created_at, updated_at FROM notifications
WHERE id = ?
`

func (q *Queries) GetNotificationById(ctx context.Context, id uint32) (Notification, error) {
	row := q.db.QueryRowContext(ctx, getNotificationById, id)
	var i Notification
	err := row.Scan(
		&i.ID,
		&i.SenderAccountID,
		&i.RecipientAccountID,
		&i.Type,
		&i.Content,
		&i.IsRead,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getNotificationCountByRecipientId = `-- name: GetNotificationCountByRecipientId :one
SELECT COUNT(*) FROM notifications
WHERE recipient_account_id = ?
`

func (q *Queries) GetNotificationCountByRecipientId(ctx context.Context, recipientAccountID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getNotificationCountByRecipientId, recipientAccountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getNotificationsByRecipientId = `-- name: GetNotificationsByRecipientId :many
SELECT id, sender_account_id, recipient_account_id, type, content, is_read, created_at, updated_at FROM notifications
WHERE recipient_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?
`

type GetNotificationsByRecipientIdParams struct {
	RecipientAccountID string
	Limit              int32
	Offset             int32
}

func (q *Queries) GetNotificationsByRecipientId(ctx context.Context, arg GetNotificationsByRecipientIdParams) ([]Notification, error) {
	rows, err := q.db.QueryContext(ctx, getNotificationsByRecipientId, arg.RecipientAccountID, arg.Limit, arg.Offset)
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
			&i.IsRead,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getNotificationsByType = `-- name: GetNotificationsByType :many
SELECT id, sender_account_id, recipient_account_id, type, content, is_read, created_at, updated_at FROM notifications
WHERE recipient_account_id = ? AND type = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?
`

type GetNotificationsByTypeParams struct {
	RecipientAccountID string
	Type               string
	Limit              int32
	Offset             int32
}

func (q *Queries) GetNotificationsByType(ctx context.Context, arg GetNotificationsByTypeParams) ([]Notification, error) {
	rows, err := q.db.QueryContext(ctx, getNotificationsByType,
		arg.RecipientAccountID,
		arg.Type,
		arg.Limit,
		arg.Offset,
	)
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
			&i.IsRead,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getUnreadNotificationCountByRecipientId = `-- name: GetUnreadNotificationCountByRecipientId :one
SELECT COUNT(*) FROM notifications
WHERE recipient_account_id = ? AND is_read = FALSE
`

func (q *Queries) GetUnreadNotificationCountByRecipientId(ctx context.Context, recipientAccountID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getUnreadNotificationCountByRecipientId, recipientAccountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUnreadNotificationsByRecipientId = `-- name: GetUnreadNotificationsByRecipientId :many
SELECT id, sender_account_id, recipient_account_id, type, content, is_read, created_at, updated_at FROM notifications
WHERE recipient_account_id = ? AND is_read = FALSE
ORDER BY created_at DESC
LIMIT ? OFFSET ?
`

type GetUnreadNotificationsByRecipientIdParams struct {
	RecipientAccountID string
	Limit              int32
	Offset             int32
}

func (q *Queries) GetUnreadNotificationsByRecipientId(ctx context.Context, arg GetUnreadNotificationsByRecipientIdParams) ([]Notification, error) {
	rows, err := q.db.QueryContext(ctx, getUnreadNotificationsByRecipientId, arg.RecipientAccountID, arg.Limit, arg.Offset)
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
			&i.IsRead,
			&i.CreatedAt,
			&i.UpdatedAt,
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
WHERE recipient_account_id = ? AND is_read = FALSE
`

func (q *Queries) MarkAllNotificationsAsRead(ctx context.Context, recipientAccountID string) error {
	_, err := q.db.ExecContext(ctx, markAllNotificationsAsRead, recipientAccountID)
	return err
}

const markNotificationAsRead = `-- name: MarkNotificationAsRead :exec
UPDATE notifications
SET is_read = TRUE
WHERE id = ? AND recipient_account_id = ?
`

type MarkNotificationAsReadParams struct {
	ID                 uint32
	RecipientAccountID string
}

func (q *Queries) MarkNotificationAsRead(ctx context.Context, arg MarkNotificationAsReadParams) error {
	_, err := q.db.ExecContext(ctx, markNotificationAsRead, arg.ID, arg.RecipientAccountID)
	return err
}
