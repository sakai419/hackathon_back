// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: retweets_and_quotes.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createRetweetOrQuote = `-- name: CreateRetweetOrQuote :exec
INSERT INTO retweets_and_quotes (retweet_id, retweeting_account_id, original_tweet_id)
VALUES (?, ?, ?)
`

type CreateRetweetOrQuoteParams struct {
	RetweetID           uint64
	RetweetingAccountID string
	OriginalTweetID     uint64
}

func (q *Queries) CreateRetweetOrQuote(ctx context.Context, arg CreateRetweetOrQuoteParams) error {
	_, err := q.db.ExecContext(ctx, createRetweetOrQuote, arg.RetweetID, arg.RetweetingAccountID, arg.OriginalTweetID)
	return err
}

const deleteRetweetOrQuote = `-- name: DeleteRetweetOrQuote :exec
DELETE FROM retweets_and_quotes
WHERE retweet_id = ?
`

func (q *Queries) DeleteRetweetOrQuote(ctx context.Context, retweetID uint64) error {
	_, err := q.db.ExecContext(ctx, deleteRetweetOrQuote, retweetID)
	return err
}

const getRetweetAndQuoteCount = `-- name: GetRetweetAndQuoteCount :one
SELECT COUNT(*) FROM retweets_and_quotes
WHERE original_tweet_id = ?
`

func (q *Queries) GetRetweetAndQuoteCount(ctx context.Context, originalTweetID uint64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getRetweetAndQuoteCount, originalTweetID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getRetweetOrQuoteById = `-- name: GetRetweetOrQuoteById :one
SELECT retweet_id, retweeting_account_id, original_tweet_id, created_at FROM retweets_and_quotes
WHERE retweet_id = ?
`

func (q *Queries) GetRetweetOrQuoteById(ctx context.Context, retweetID uint64) (RetweetsAndQuote, error) {
	row := q.db.QueryRowContext(ctx, getRetweetOrQuoteById, retweetID)
	var i RetweetsAndQuote
	err := row.Scan(
		&i.RetweetID,
		&i.RetweetingAccountID,
		&i.OriginalTweetID,
		&i.CreatedAt,
	)
	return i, err
}

const getRetweetsAndQuotesByAccountId = `-- name: GetRetweetsAndQuotesByAccountId :many
SELECT r.retweet_id, r.retweeting_account_id, r.original_tweet_id, r.created_at, t.content AS original_tweet_content
FROM retweets_and_quotes r
JOIN tweets t ON r.original_tweet_id = t.id
WHERE r.retweeting_account_id = ?
ORDER BY r.created_at DESC
LIMIT ? OFFSET ?
`

type GetRetweetsAndQuotesByAccountIdParams struct {
	RetweetingAccountID string
	Limit               int32
	Offset              int32
}

type GetRetweetsAndQuotesByAccountIdRow struct {
	RetweetID            uint64
	RetweetingAccountID  string
	OriginalTweetID      uint64
	CreatedAt            time.Time
	OriginalTweetContent sql.NullString
}

func (q *Queries) GetRetweetsAndQuotesByAccountId(ctx context.Context, arg GetRetweetsAndQuotesByAccountIdParams) ([]GetRetweetsAndQuotesByAccountIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getRetweetsAndQuotesByAccountId, arg.RetweetingAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRetweetsAndQuotesByAccountIdRow
	for rows.Next() {
		var i GetRetweetsAndQuotesByAccountIdRow
		if err := rows.Scan(
			&i.RetweetID,
			&i.RetweetingAccountID,
			&i.OriginalTweetID,
			&i.CreatedAt,
			&i.OriginalTweetContent,
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

const getRetweetsAndQuotesByOriginalTweetId = `-- name: GetRetweetsAndQuotesByOriginalTweetId :many
SELECT r.retweet_id, r.retweeting_account_id, r.original_tweet_id, r.created_at, t.content AS retweet_content
FROM retweets_and_quotes r
JOIN tweets t ON r.retweet_id = t.id
WHERE r.original_tweet_id = ?
ORDER BY r.created_at DESC
LIMIT ? OFFSET ?
`

type GetRetweetsAndQuotesByOriginalTweetIdParams struct {
	OriginalTweetID uint64
	Limit           int32
	Offset          int32
}

type GetRetweetsAndQuotesByOriginalTweetIdRow struct {
	RetweetID           uint64
	RetweetingAccountID string
	OriginalTweetID     uint64
	CreatedAt           time.Time
	RetweetContent      sql.NullString
}

func (q *Queries) GetRetweetsAndQuotesByOriginalTweetId(ctx context.Context, arg GetRetweetsAndQuotesByOriginalTweetIdParams) ([]GetRetweetsAndQuotesByOriginalTweetIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getRetweetsAndQuotesByOriginalTweetId, arg.OriginalTweetID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRetweetsAndQuotesByOriginalTweetIdRow
	for rows.Next() {
		var i GetRetweetsAndQuotesByOriginalTweetIdRow
		if err := rows.Scan(
			&i.RetweetID,
			&i.RetweetingAccountID,
			&i.OriginalTweetID,
			&i.CreatedAt,
			&i.RetweetContent,
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
