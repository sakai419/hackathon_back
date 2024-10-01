// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: likes.sql

package sqlc

import (
	"context"
)

const createLike = `-- name: CreateLike :exec
INSERT INTO likes (liking_account_id, original_tweet_id)
VALUES (?, ?)
`

type CreateLikeParams struct {
	LikingAccountID string
	OriginalTweetID uint64
}

func (q *Queries) CreateLike(ctx context.Context, arg CreateLikeParams) error {
	_, err := q.db.ExecContext(ctx, createLike, arg.LikingAccountID, arg.OriginalTweetID)
	return err
}

const deleteLike = `-- name: DeleteLike :exec
DELETE FROM likes
WHERE liking_account_id = ? AND original_tweet_id = ?
`

type DeleteLikeParams struct {
	LikingAccountID string
	OriginalTweetID uint64
}

func (q *Queries) DeleteLike(ctx context.Context, arg DeleteLikeParams) error {
	_, err := q.db.ExecContext(ctx, deleteLike, arg.LikingAccountID, arg.OriginalTweetID)
	return err
}

const getLikeCount = `-- name: GetLikeCount :one
SELECT COUNT(*) FROM likes
WHERE original_tweet_id = ?
`

func (q *Queries) GetLikeCount(ctx context.Context, originalTweetID uint64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getLikeCount, originalTweetID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getLikesByAccountId = `-- name: GetLikesByAccountId :many
SELECT liking_account_id, original_tweet_id, created_at FROM likes
WHERE liking_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?
`

type GetLikesByAccountIdParams struct {
	LikingAccountID string
	Limit           int32
	Offset          int32
}

func (q *Queries) GetLikesByAccountId(ctx context.Context, arg GetLikesByAccountIdParams) ([]Like, error) {
	rows, err := q.db.QueryContext(ctx, getLikesByAccountId, arg.LikingAccountID, arg.Limit, arg.Offset)
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

const getLikesByTweetId = `-- name: GetLikesByTweetId :many
SELECT liking_account_id, original_tweet_id, created_at FROM likes
WHERE original_tweet_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?
`

type GetLikesByTweetIdParams struct {
	OriginalTweetID uint64
	Limit           int32
	Offset          int32
}

func (q *Queries) GetLikesByTweetId(ctx context.Context, arg GetLikesByTweetIdParams) ([]Like, error) {
	rows, err := q.db.QueryContext(ctx, getLikesByTweetId, arg.OriginalTweetID, arg.Limit, arg.Offset)
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
