// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: conversations.sql

package sqlcgen

import (
	"context"
	"database/sql"
	"time"
)

const getConversationID = `-- name: GetConversationID :one
WITH existing_conversation AS (
    SELECT id FROM conversations c
    WHERE (c.account1_id = $1 AND c.account2_id = $2)
       OR (c.account1_id = $2 AND c.account2_id = $1)
),
inserted_conversation AS (
    INSERT INTO conversations (account1_id, account2_id)
    SELECT $1, $2
    WHERE NOT EXISTS (SELECT 1 FROM existing_conversation)
    RETURNING id
)
SELECT id FROM inserted_conversation
UNION ALL
SELECT id FROM existing_conversation
LIMIT 1
`

type GetConversationIDParams struct {
	Account1ID string
	Account2ID string
}

func (q *Queries) GetConversationID(ctx context.Context, arg GetConversationIDParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getConversationID, arg.Account1ID, arg.Account2ID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getConversationList = `-- name: GetConversationList :many
SELECT
    c.id,
    c.account1_id,
    c.account2_id,
    c.last_message_time,
    m.content,
    m.sender_account_id,
    m.is_read
FROM
    conversations c
LEFT JOIN
    messages m ON c.last_message_id = m.id
WHERE
    c.account1_id = $1 OR c.account2_id = $1
ORDER BY
    c.last_message_time DESC
LIMIT $2 OFFSET $3
`

type GetConversationListParams struct {
	Account1ID string
	Limit      int32
	Offset     int32
}

type GetConversationListRow struct {
	ID              int64
	Account1ID      string
	Account2ID      string
	LastMessageTime time.Time
	Content         sql.NullString
	SenderAccountID sql.NullString
	IsRead          sql.NullBool
}

func (q *Queries) GetConversationList(ctx context.Context, arg GetConversationListParams) ([]GetConversationListRow, error) {
	rows, err := q.db.QueryContext(ctx, getConversationList, arg.Account1ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetConversationListRow
	for rows.Next() {
		var i GetConversationListRow
		if err := rows.Scan(
			&i.ID,
			&i.Account1ID,
			&i.Account2ID,
			&i.LastMessageTime,
			&i.Content,
			&i.SenderAccountID,
			&i.IsRead,
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

const getUnreadConversationCount = `-- name: GetUnreadConversationCount :one
SELECT COUNT(DISTINCT
    CASE
        WHEN c.account1_id = $1 THEN c.account2_id
        ELSE c.account1_id
    END) AS unread_account_count
FROM
    conversations c
JOIN
    messages m ON c.last_message_id = m.id
WHERE
    (c.account1_id = $1 OR c.account2_id = $1)
    AND m.is_read = FALSE
    AND m.sender_account_id <> $1
`

func (q *Queries) GetUnreadConversationCount(ctx context.Context, account1ID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getUnreadConversationCount, account1ID)
	var unread_account_count int64
	err := row.Scan(&unread_account_count)
	return unread_account_count, err
}

const updateLastMessage = `-- name: UpdateLastMessage :exec
UPDATE conversations
SET last_message_id = $1, last_message_time = NOW()
WHERE id = $2
`

type UpdateLastMessageParams struct {
	LastMessageID sql.NullInt64
	ID            int64
}

func (q *Queries) UpdateLastMessage(ctx context.Context, arg UpdateLastMessageParams) error {
	_, err := q.db.ExecContext(ctx, updateLastMessage, arg.LastMessageID, arg.ID)
	return err
}