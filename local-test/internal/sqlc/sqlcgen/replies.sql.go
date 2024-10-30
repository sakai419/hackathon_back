// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: replies.sql

package sqlcgen

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const checkParentReplyExist = `-- name: CheckParentReplyExist :one
SELECT EXISTS(SELECT 1 FROM replies WHERE reply_id = $1 AND parent_reply_id IS NOT NULL)
`

func (q *Queries) CheckParentReplyExist(ctx context.Context, replyID int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkParentReplyExist, replyID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createReply = `-- name: CreateReply :exec
INSERT INTO replies (reply_id, original_tweet_id, parent_reply_id, replying_account_id)
SELECT
    $1 AS reply_id,
    COALESCE(
        (SELECT r.original_tweet_id FROM replies AS r WHERE r.reply_id = $2),
        $2
    ) AS original_tweet_id,
    COALESCE(
        (SELECT r.reply_id FROM replies AS r WHERE r.reply_id = $2),
        NULL
    ) AS parent_reply_id,
    $3 AS replying_account_id
`

type CreateReplyParams struct {
	ReplyID           int64
	OriginalTweetID   int64
	ReplyingAccountID string
}

func (q *Queries) CreateReply(ctx context.Context, arg CreateReplyParams) error {
	_, err := q.db.ExecContext(ctx, createReply, arg.ReplyID, arg.OriginalTweetID, arg.ReplyingAccountID)
	return err
}

const getReplyRelations = `-- name: GetReplyRelations :many
SELECT reply_id, original_tweet_id, parent_reply_id
FROM replies
WHERE reply_id = ANY($1::BIGINT[])
`

type GetReplyRelationsRow struct {
	ReplyID         int64
	OriginalTweetID int64
	ParentReplyID   sql.NullInt64
}

func (q *Queries) GetReplyRelations(ctx context.Context, tweetIds []int64) ([]GetReplyRelationsRow, error) {
	rows, err := q.db.QueryContext(ctx, getReplyRelations, pq.Array(tweetIds))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetReplyRelationsRow
	for rows.Next() {
		var i GetReplyRelationsRow
		if err := rows.Scan(&i.ReplyID, &i.OriginalTweetID, &i.ParentReplyID); err != nil {
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
