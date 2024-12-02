// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: messages.sql

package sqlcgen

import (
	"context"
	"time"
)

const createMessage = `-- name: CreateMessage :one
INSERT INTO messages (conversation_id, sender_account_id, content, is_read)
VALUES ($1, $2, $3, FALSE)
RETURNING id
`

type CreateMessageParams struct {
	ConversationID  int64
	SenderAccountID string
	Content         string
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createMessage, arg.ConversationID, arg.SenderAccountID, arg.Content)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getMessageList = `-- name: GetMessageList :many
SELECT m.id, a.user_id, m.content, m.is_read, m.created_at FROM messages AS m
INNER JOIN accounts AS a ON m.sender_account_id = a.id
WHERE m.conversation_id = $1
ORDER BY m.created_at DESC
LIMIT $2 OFFSET $3
`

type GetMessageListParams struct {
	ConversationID int64
	Limit          int32
	Offset         int32
}

type GetMessageListRow struct {
	ID        int64
	UserID    string
	Content   string
	IsRead    bool
	CreatedAt time.Time
}

func (q *Queries) GetMessageList(ctx context.Context, arg GetMessageListParams) ([]GetMessageListRow, error) {
	rows, err := q.db.QueryContext(ctx, getMessageList, arg.ConversationID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMessageListRow
	for rows.Next() {
		var i GetMessageListRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
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

const markMessageListAsRead = `-- name: MarkMessageListAsRead :exec
UPDATE messages
SET is_read = TRUE
WHERE conversation_id = $1
  AND sender_account_id <> $2
  AND is_read = FALSE
  AND created_at <= NOW()
`

type MarkMessageListAsReadParams struct {
	ConversationID  int64
	SenderAccountID string
}

func (q *Queries) MarkMessageListAsRead(ctx context.Context, arg MarkMessageListAsReadParams) error {
	_, err := q.db.ExecContext(ctx, markMessageListAsRead, arg.ConversationID, arg.SenderAccountID)
	return err
}
