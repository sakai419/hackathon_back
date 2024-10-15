// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: messages.sql

package sqlcgen

import (
	"context"
	"database/sql"
	"time"
)

const createMessage = `-- name: CreateMessage :exec
INSERT INTO messages (sender_account_id, recipient_account_id, content)
VALUES ($1, $2, $3)
`

type CreateMessageParams struct {
	SenderAccountID    string
	RecipientAccountID string
	Content            sql.NullString
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) error {
	_, err := q.db.ExecContext(ctx, createMessage, arg.SenderAccountID, arg.RecipientAccountID, arg.Content)
	return err
}

const deleteMessage = `-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id = $1 AND (sender_account_id = $2 OR recipient_account_id = $3)
`

type DeleteMessageParams struct {
	ID                 int64
	SenderAccountID    string
	RecipientAccountID string
}

func (q *Queries) DeleteMessage(ctx context.Context, arg DeleteMessageParams) error {
	_, err := q.db.ExecContext(ctx, deleteMessage, arg.ID, arg.SenderAccountID, arg.RecipientAccountID)
	return err
}

const deleteOldMessages = `-- name: DeleteOldMessages :exec
DELETE FROM messages
WHERE created_at < $1 AND is_read = TRUE
`

func (q *Queries) DeleteOldMessages(ctx context.Context, createdAt time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteOldMessages, createdAt)
	return err
}

const gerUnreadMessageCountBetweenUsers = `-- name: GerUnreadMessageCountBetweenUsers :one
SELECT COUNT(*) FROM messages
WHERE recipient_account_id = $1 AND sender_account_id = $2 AND is_read = FALSE
`

type GerUnreadMessageCountBetweenUsersParams struct {
	RecipientAccountID string
	SenderAccountID    string
}

func (q *Queries) GerUnreadMessageCountBetweenUsers(ctx context.Context, arg GerUnreadMessageCountBetweenUsersParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, gerUnreadMessageCountBetweenUsers, arg.RecipientAccountID, arg.SenderAccountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getLatestMessageForEachConversation = `-- name: GetLatestMessageForEachConversation :many
SELECT m.id, m.sender_account_id, m.recipient_account_id, m.content, m.is_read, m.created_at
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
ORDER BY m.created_at DESC
`

type GetLatestMessageForEachConversationParams struct {
	SenderAccountID    string
	RecipientAccountID string
}

func (q *Queries) GetLatestMessageForEachConversation(ctx context.Context, arg GetLatestMessageForEachConversationParams) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, getLatestMessageForEachConversation, arg.SenderAccountID, arg.RecipientAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.SenderAccountID,
			&i.RecipientAccountID,
			&i.Content,
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

const getMessageByID = `-- name: GetMessageByID :one
SELECT id, sender_account_id, recipient_account_id, content, is_read, created_at FROM messages
WHERE id = $1
`

func (q *Queries) GetMessageByID(ctx context.Context, id int64) (Message, error) {
	row := q.db.QueryRowContext(ctx, getMessageByID, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.SenderAccountID,
		&i.RecipientAccountID,
		&i.Content,
		&i.IsRead,
		&i.CreatedAt,
	)
	return i, err
}

const getMessageCountForUser = `-- name: GetMessageCountForUser :one
SELECT COUNT(*) FROM messages
WHERE sender_account_id = $1 OR recipient_account_id = $2
`

type GetMessageCountForUserParams struct {
	SenderAccountID    string
	RecipientAccountID string
}

func (q *Queries) GetMessageCountForUser(ctx context.Context, arg GetMessageCountForUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getMessageCountForUser, arg.SenderAccountID, arg.RecipientAccountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getMessagesBetweenUsers = `-- name: GetMessagesBetweenUsers :many
SELECT id, sender_account_id, recipient_account_id, content, is_read, created_at FROM messages
WHERE (sender_account_id = $1 AND recipient_account_id = $2)
    OR (sender_account_id = $3 AND recipient_account_id = $4)
ORDER BY created_at DESC
LIMIT $5 OFFSET $6
`

type GetMessagesBetweenUsersParams struct {
	SenderAccountID      string
	RecipientAccountID   string
	SenderAccountID_2    string
	RecipientAccountID_2 string
	Limit                int32
	Offset               int32
}

func (q *Queries) GetMessagesBetweenUsers(ctx context.Context, arg GetMessagesBetweenUsersParams) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, getMessagesBetweenUsers,
		arg.SenderAccountID,
		arg.RecipientAccountID,
		arg.SenderAccountID_2,
		arg.RecipientAccountID_2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.SenderAccountID,
			&i.RecipientAccountID,
			&i.Content,
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

const getUnreadMessageCountForUser = `-- name: GetUnreadMessageCountForUser :one
SELECT COUNT(*) FROM messages
WHERE recipient_account_id = $1 AND is_read = FALSE
`

func (q *Queries) GetUnreadMessageCountForUser(ctx context.Context, recipientAccountID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getUnreadMessageCountForUser, recipientAccountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUnreadMessagesForUser = `-- name: GetUnreadMessagesForUser :many
SELECT id, sender_account_id, recipient_account_id, content, is_read, created_at FROM messages
WHERE recipient_account_id = $1 AND is_read = FALSE
ORDER BY created_at DESC
`

func (q *Queries) GetUnreadMessagesForUser(ctx context.Context, recipientAccountID string) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, getUnreadMessagesForUser, recipientAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.SenderAccountID,
			&i.RecipientAccountID,
			&i.Content,
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

const markAllMessagesAsRead = `-- name: MarkAllMessagesAsRead :exec
UPDATE messages
SET is_read = TRUE
WHERE recipient_account_id = $1 AND is_read = FALSE
`

func (q *Queries) MarkAllMessagesAsRead(ctx context.Context, recipientAccountID string) error {
	_, err := q.db.ExecContext(ctx, markAllMessagesAsRead, recipientAccountID)
	return err
}

const markMessageAsRead = `-- name: MarkMessageAsRead :exec
UPDATE messages
SET is_read = TRUE
WHERE id = $1 AND recipient_account_id = $2
`

type MarkMessageAsReadParams struct {
	ID                 int64
	RecipientAccountID string
}

func (q *Queries) MarkMessageAsRead(ctx context.Context, arg MarkMessageAsReadParams) error {
	_, err := q.db.ExecContext(ctx, markMessageAsRead, arg.ID, arg.RecipientAccountID)
	return err
}

const searchMessages = `-- name: SearchMessages :many
SELECT id, sender_account_id, recipient_account_id, content, is_read, created_at FROM messages
WHERE (sender_account_id = $1 OR recipient_account_id = $2)
    AND content LIKE $3
ORDER BY created_at DESC
LIMIT $4 OFFSET $5
`

type SearchMessagesParams struct {
	SenderAccountID    string
	RecipientAccountID string
	Content            sql.NullString
	Limit              int32
	Offset             int32
}

func (q *Queries) SearchMessages(ctx context.Context, arg SearchMessagesParams) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, searchMessages,
		arg.SenderAccountID,
		arg.RecipientAccountID,
		arg.Content,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.SenderAccountID,
			&i.RecipientAccountID,
			&i.Content,
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
