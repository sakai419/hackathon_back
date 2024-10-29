// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: likes.sql

package sqlcgen

import (
	"context"
	"database/sql"
)

const createLike = `-- name: CreateLike :exec
INSERT INTO likes (liking_account_id, original_tweet_id)
VALUES ($1, $2)
`

type CreateLikeParams struct {
	LikingAccountID string
	OriginalTweetID int64
}

func (q *Queries) CreateLike(ctx context.Context, arg CreateLikeParams) error {
	_, err := q.db.ExecContext(ctx, createLike, arg.LikingAccountID, arg.OriginalTweetID)
	return err
}

const deleteLike = `-- name: DeleteLike :execresult
DELETE FROM likes
WHERE liking_account_id = $1 AND original_tweet_id = $2
`

type DeleteLikeParams struct {
	LikingAccountID string
	OriginalTweetID int64
}

func (q *Queries) DeleteLike(ctx context.Context, arg DeleteLikeParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteLike, arg.LikingAccountID, arg.OriginalTweetID)
}

const getLikeCount = `-- name: GetLikeCount :one
SELECT COUNT(*) FROM likes
WHERE original_tweet_id = $1
`

func (q *Queries) GetLikeCount(ctx context.Context, originalTweetID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getLikeCount, originalTweetID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getLikesByAccountID = `-- name: GetLikesByAccountID :many
SELECT liking_account_id, original_tweet_id, created_at FROM likes
WHERE liking_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetLikesByAccountIDParams struct {
	LikingAccountID string
	Limit           int32
	Offset          int32
}

func (q *Queries) GetLikesByAccountID(ctx context.Context, arg GetLikesByAccountIDParams) ([]Like, error) {
	rows, err := q.db.QueryContext(ctx, getLikesByAccountID, arg.LikingAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Like
	for rows.Next() {
		var i Like
		if err := rows.Scan(&i.LikingAccountID, &i.OriginalTweetID, &i.CreatedAt); err != nil {
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

const getLikingAccountIDs = `-- name: GetLikingAccountIDs :many
SELECT liking_account_id FROM likes
WHERE original_tweet_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetLikingAccountIDsParams struct {
	OriginalTweetID int64
	Limit           int32
	Offset          int32
}

func (q *Queries) GetLikingAccountIDs(ctx context.Context, arg GetLikingAccountIDsParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getLikingAccountIDs, arg.OriginalTweetID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var liking_account_id string
		if err := rows.Scan(&liking_account_id); err != nil {
			return nil, err
		}
		items = append(items, liking_account_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
