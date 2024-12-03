// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: quotes.sql

package sqlcgen

import (
	"context"

	"github.com/lib/pq"
)

const createQuote = `-- name: CreateQuote :exec
INSERT INTO quotes (quote_id, quoting_account_id, original_tweet_id)
VALUES ($1, $2, $3)
`

type CreateQuoteParams struct {
	QuoteID          int64
	QuotingAccountID string
	OriginalTweetID  int64
}

func (q *Queries) CreateQuote(ctx context.Context, arg CreateQuoteParams) error {
	_, err := q.db.ExecContext(ctx, createQuote, arg.QuoteID, arg.QuotingAccountID, arg.OriginalTweetID)
	return err
}

const getQuoteRelations = `-- name: GetQuoteRelations :many
SELECT quote_id, original_tweet_id
FROM quotes
WHERE quote_id = ANY($1::BIGINT[])
`

type GetQuoteRelationsRow struct {
	QuoteID         int64
	OriginalTweetID int64
}

func (q *Queries) GetQuoteRelations(ctx context.Context, tweetIds []int64) ([]GetQuoteRelationsRow, error) {
	rows, err := q.db.QueryContext(ctx, getQuoteRelations, pq.Array(tweetIds))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetQuoteRelationsRow
	for rows.Next() {
		var i GetQuoteRelationsRow
		if err := rows.Scan(&i.QuoteID, &i.OriginalTweetID); err != nil {
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

const getQuotingAccountIDs = `-- name: GetQuotingAccountIDs :many
SELECT quoting_account_id
FROM quotes
WHERE original_tweet_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetQuotingAccountIDsParams struct {
	OriginalTweetID int64
	Limit           int32
	Offset          int32
}

func (q *Queries) GetQuotingAccountIDs(ctx context.Context, arg GetQuotingAccountIDsParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getQuotingAccountIDs, arg.OriginalTweetID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var quoting_account_id string
		if err := rows.Scan(&quoting_account_id); err != nil {
			return nil, err
		}
		items = append(items, quoting_account_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}