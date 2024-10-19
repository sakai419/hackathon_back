// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: retweets_and_quotes.sql

package sqlcgen

import (
	"context"
	"database/sql"
	"time"
)

const createRetweetOrQuote = `-- name: CreateRetweetOrQuote :exec
INSERT INTO retweets_and_quotes (retweeting_account_id, original_tweet_id)
VALUES ($1, $2)
`

type CreateRetweetOrQuoteParams struct {
	RetweetingAccountID string
	OriginalTweetID     int64
}

func (q *Queries) CreateRetweetOrQuote(ctx context.Context, arg CreateRetweetOrQuoteParams) error {
	_, err := q.db.ExecContext(ctx, createRetweetOrQuote, arg.RetweetingAccountID, arg.OriginalTweetID)
	return err
}

const deleteRetweetOrQuote = `-- name: DeleteRetweetOrQuote :exec
DELETE FROM retweets_and_quotes
WHERE tweet_id = $1
`

func (q *Queries) DeleteRetweetOrQuote(ctx context.Context, tweetID int64) error {
	_, err := q.db.ExecContext(ctx, deleteRetweetOrQuote, tweetID)
	return err
}

const getRetweetAndQuoteCount = `-- name: GetRetweetAndQuoteCount :one
SELECT COUNT(*) FROM retweets_and_quotes
WHERE original_tweet_id = $1
`

func (q *Queries) GetRetweetAndQuoteCount(ctx context.Context, originalTweetID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getRetweetAndQuoteCount, originalTweetID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getRetweetOrQuoteByID = `-- name: GetRetweetOrQuoteByID :one
SELECT tweet_id, retweeting_account_id, original_tweet_id, created_at FROM retweets_and_quotes
WHERE tweet_id = $1
`

func (q *Queries) GetRetweetOrQuoteByID(ctx context.Context, tweetID int64) (RetweetsAndQuote, error) {
	row := q.db.QueryRowContext(ctx, getRetweetOrQuoteByID, tweetID)
	var i RetweetsAndQuote
	err := row.Scan(
		&i.TweetID,
		&i.RetweetingAccountID,
		&i.OriginalTweetID,
		&i.CreatedAt,
	)
	return i, err
}

const getRetweetsAndQuotesByAccountID = `-- name: GetRetweetsAndQuotesByAccountID :many
SELECT r.tweet_id, r.retweeting_account_id, r.original_tweet_id, r.created_at, t.content AS original_tweet_content
FROM retweets_and_quotes r
JOIN tweets t ON r.original_tweet_id = t.id
WHERE r.retweeting_account_id = $1
ORDER BY r.created_at DESC
LIMIT $2 OFFSET $3
`

type GetRetweetsAndQuotesByAccountIDParams struct {
	RetweetingAccountID string
	Limit               int32
	Offset              int32
}

type GetRetweetsAndQuotesByAccountIDRow struct {
	TweetID              int64
	RetweetingAccountID  string
	OriginalTweetID      int64
	CreatedAt            time.Time
	OriginalTweetContent sql.NullString
}

func (q *Queries) GetRetweetsAndQuotesByAccountID(ctx context.Context, arg GetRetweetsAndQuotesByAccountIDParams) ([]GetRetweetsAndQuotesByAccountIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getRetweetsAndQuotesByAccountID, arg.RetweetingAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRetweetsAndQuotesByAccountIDRow
	for rows.Next() {
		var i GetRetweetsAndQuotesByAccountIDRow
		if err := rows.Scan(
			&i.TweetID,
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

const getRetweetsAndQuotesByOriginalTweetID = `-- name: GetRetweetsAndQuotesByOriginalTweetID :many
SELECT r.tweet_id, r.retweeting_account_id, r.original_tweet_id, r.created_at, t.content AS retweet_content
FROM retweets_and_quotes r
JOIN tweets t ON r.tweet_id = t.id
WHERE r.original_tweet_id = $1
ORDER BY r.created_at DESC
LIMIT $2 OFFSET $3
`

type GetRetweetsAndQuotesByOriginalTweetIDParams struct {
	OriginalTweetID int64
	Limit           int32
	Offset          int32
}

type GetRetweetsAndQuotesByOriginalTweetIDRow struct {
	TweetID             int64
	RetweetingAccountID string
	OriginalTweetID     int64
	CreatedAt           time.Time
	RetweetContent      sql.NullString
}

func (q *Queries) GetRetweetsAndQuotesByOriginalTweetID(ctx context.Context, arg GetRetweetsAndQuotesByOriginalTweetIDParams) ([]GetRetweetsAndQuotesByOriginalTweetIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getRetweetsAndQuotesByOriginalTweetID, arg.OriginalTweetID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRetweetsAndQuotesByOriginalTweetIDRow
	for rows.Next() {
		var i GetRetweetsAndQuotesByOriginalTweetIDRow
		if err := rows.Scan(
			&i.TweetID,
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