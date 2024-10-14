// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tweets.sql

package sqlcgen

import (
	"context"
	"database/sql"

	"github.com/sqlc-dev/pqtype"
)

const createTweet = `-- name: CreateTweet :exec
INSERT INTO tweets (
    account_id, content, code, media
) VALUES ($1, $2, $3, $4)
`

type CreateTweetParams struct {
	AccountID string
	Content   sql.NullString
	Code      sql.NullString
	Media     pqtype.NullRawMessage
}

func (q *Queries) CreateTweet(ctx context.Context, arg CreateTweetParams) error {
	_, err := q.db.ExecContext(ctx, createTweet,
		arg.AccountID,
		arg.Content,
		arg.Code,
		arg.Media,
	)
	return err
}

const createTweetAsQuote = `-- name: CreateTweetAsQuote :exec
INSERT INTO tweets (
    account_id, is_quote, content, code, media
) VALUES ($1, TRUE, $2, $3, $4)
`

type CreateTweetAsQuoteParams struct {
	AccountID string
	Content   sql.NullString
	Code      sql.NullString
	Media     pqtype.NullRawMessage
}

func (q *Queries) CreateTweetAsQuote(ctx context.Context, arg CreateTweetAsQuoteParams) error {
	_, err := q.db.ExecContext(ctx, createTweetAsQuote,
		arg.AccountID,
		arg.Content,
		arg.Code,
		arg.Media,
	)
	return err
}

const createTweetAsReply = `-- name: CreateTweetAsReply :exec
INSERT INTO tweets (
    account_id, is_reply, content, code, media
) VALUES ($1, TRUE, $2, $3, $4)
`

type CreateTweetAsReplyParams struct {
	AccountID string
	Content   sql.NullString
	Code      sql.NullString
	Media     pqtype.NullRawMessage
}

func (q *Queries) CreateTweetAsReply(ctx context.Context, arg CreateTweetAsReplyParams) error {
	_, err := q.db.ExecContext(ctx, createTweetAsReply,
		arg.AccountID,
		arg.Content,
		arg.Code,
		arg.Media,
	)
	return err
}

const createTweetAsRetweet = `-- name: CreateTweetAsRetweet :exec
INSERT INTO tweets (
    account_id, is_retweet
) VALUES ($1, TRUE)
`

func (q *Queries) CreateTweetAsRetweet(ctx context.Context, accountID string) error {
	_, err := q.db.ExecContext(ctx, createTweetAsRetweet, accountID)
	return err
}

const deleteTweet = `-- name: DeleteTweet :exec
DELETE FROM tweets WHERE id = $1 AND account_id = $2
`

type DeleteTweetParams struct {
	ID        int64
	AccountID string
}

func (q *Queries) DeleteTweet(ctx context.Context, arg DeleteTweetParams) error {
	_, err := q.db.ExecContext(ctx, deleteTweet, arg.ID, arg.AccountID)
	return err
}

const getPinnedTweetForAccount = `-- name: GetPinnedTweetForAccount :one
SELECT id, account_id, is_pinned, content, code, likes_count, replies_count, retweets_count, is_retweet, is_reply, is_quote, engagement_score, media, created_at, updated_at FROM tweets
WHERE account_id = $1 AND is_pinned = TRUE
LIMIT 1
`

func (q *Queries) GetPinnedTweetForAccount(ctx context.Context, accountID string) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, getPinnedTweetForAccount, accountID)
	var i Tweet
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.IsPinned,
		&i.Content,
		&i.Code,
		&i.LikesCount,
		&i.RepliesCount,
		&i.RetweetsCount,
		&i.IsRetweet,
		&i.IsReply,
		&i.IsQuote,
		&i.EngagementScore,
		&i.Media,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTrendingTweets = `-- name: GetTrendingTweets :many
SELECT id, account_id, is_pinned, content, code, likes_count, replies_count, retweets_count, is_retweet, is_reply, is_quote, engagement_score, media, created_at, updated_at FROM tweets
ORDER BY engagement_score DESC
LIMIT $1 OFFSET $2
`

type GetTrendingTweetsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetTrendingTweets(ctx context.Context, arg GetTrendingTweetsParams) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, getTrendingTweets, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tweet
	for rows.Next() {
		var i Tweet
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.IsPinned,
			&i.Content,
			&i.Code,
			&i.LikesCount,
			&i.RepliesCount,
			&i.RetweetsCount,
			&i.IsRetweet,
			&i.IsReply,
			&i.IsQuote,
			&i.EngagementScore,
			&i.Media,
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

const getTweetById = `-- name: GetTweetById :one
SELECT id, account_id, is_pinned, content, code, likes_count, replies_count, retweets_count, is_retweet, is_reply, is_quote, engagement_score, media, created_at, updated_at FROM tweets WHERE id = $1
`

func (q *Queries) GetTweetById(ctx context.Context, id int64) (Tweet, error) {
	row := q.db.QueryRowContext(ctx, getTweetById, id)
	var i Tweet
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.IsPinned,
		&i.Content,
		&i.Code,
		&i.LikesCount,
		&i.RepliesCount,
		&i.RetweetsCount,
		&i.IsRetweet,
		&i.IsReply,
		&i.IsQuote,
		&i.EngagementScore,
		&i.Media,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTweetCountByAccountId = `-- name: GetTweetCountByAccountId :one
SELECT COUNT(*) FROM tweets WHERE account_id = $1
`

func (q *Queries) GetTweetCountByAccountId(ctx context.Context, accountID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTweetCountByAccountId, accountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getTweetsByAccountId = `-- name: GetTweetsByAccountId :many
SELECT id, account_id, is_pinned, content, code, likes_count, replies_count, retweets_count, is_retweet, is_reply, is_quote, engagement_score, media, created_at, updated_at FROM tweets
WHERE account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetTweetsByAccountIdParams struct {
	AccountID string
	Limit     int32
	Offset    int32
}

func (q *Queries) GetTweetsByAccountId(ctx context.Context, arg GetTweetsByAccountIdParams) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, getTweetsByAccountId, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tweet
	for rows.Next() {
		var i Tweet
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.IsPinned,
			&i.Content,
			&i.Code,
			&i.LikesCount,
			&i.RepliesCount,
			&i.RetweetsCount,
			&i.IsRetweet,
			&i.IsReply,
			&i.IsQuote,
			&i.EngagementScore,
			&i.Media,
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

const incrementLikesCount = `-- name: IncrementLikesCount :exec
UPDATE tweets SET likes_count = likes_count + 1 WHERE id = $1
`

func (q *Queries) IncrementLikesCount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, incrementLikesCount, id)
	return err
}

const incrementRepliesCount = `-- name: IncrementRepliesCount :exec
UPDATE tweets SET replies_count = replies_count + 1 WHERE id = $1
`

func (q *Queries) IncrementRepliesCount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, incrementRepliesCount, id)
	return err
}

const incrementRetweetsCount = `-- name: IncrementRetweetsCount :exec
UPDATE tweets SET retweets_count = retweets_count + 1 WHERE id = $1
`

func (q *Queries) IncrementRetweetsCount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, incrementRetweetsCount, id)
	return err
}

const searchTweets = `-- name: SearchTweets :many
SELECT id, account_id, is_pinned, content, code, likes_count, replies_count, retweets_count, is_retweet, is_reply, is_quote, engagement_score, media, created_at, updated_at FROM tweets
WHERE content LIKE $1 OR code LIKE $2
ORDER BY created_at DESC
LIMIT $3 OFFSET $4
`

type SearchTweetsParams struct {
	Content sql.NullString
	Code    sql.NullString
	Limit   int32
	Offset  int32
}

func (q *Queries) SearchTweets(ctx context.Context, arg SearchTweetsParams) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, searchTweets,
		arg.Content,
		arg.Code,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tweet
	for rows.Next() {
		var i Tweet
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.IsPinned,
			&i.Content,
			&i.Code,
			&i.LikesCount,
			&i.RepliesCount,
			&i.RetweetsCount,
			&i.IsRetweet,
			&i.IsReply,
			&i.IsQuote,
			&i.EngagementScore,
			&i.Media,
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

const setTweetAsPinned = `-- name: SetTweetAsPinned :exec
UPDATE tweets
SET is_pinned = TRUE
WHERE id = $1 AND account_id = $2
`

type SetTweetAsPinnedParams struct {
	ID        int64
	AccountID string
}

func (q *Queries) SetTweetAsPinned(ctx context.Context, arg SetTweetAsPinnedParams) error {
	_, err := q.db.ExecContext(ctx, setTweetAsPinned, arg.ID, arg.AccountID)
	return err
}

const unpinTweet = `-- name: UnpinTweet :exec
UPDATE tweets
SET is_pinned = FALSE
WHERE id = $1 AND account_id = $2
`

type UnpinTweetParams struct {
	ID        int64
	AccountID string
}

func (q *Queries) UnpinTweet(ctx context.Context, arg UnpinTweetParams) error {
	_, err := q.db.ExecContext(ctx, unpinTweet, arg.ID, arg.AccountID)
	return err
}

const updateEngagementScore = `-- name: UpdateEngagementScore :exec
UPDATE tweets
SET engagement_score = likes_count + replies_count + retweets_count
WHERE id = $1
`

func (q *Queries) UpdateEngagementScore(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, updateEngagementScore, id)
	return err
}

const updateTweetCode = `-- name: UpdateTweetCode :exec
UPDATE tweets
SET code = $1
WHERE id = $2 AND account_id = $3
`

type UpdateTweetCodeParams struct {
	Code      sql.NullString
	ID        int64
	AccountID string
}

func (q *Queries) UpdateTweetCode(ctx context.Context, arg UpdateTweetCodeParams) error {
	_, err := q.db.ExecContext(ctx, updateTweetCode, arg.Code, arg.ID, arg.AccountID)
	return err
}

const updateTweetContent = `-- name: UpdateTweetContent :exec
UPDATE tweets
SET content = $1
WHERE id = $2 AND account_id = $3
`

type UpdateTweetContentParams struct {
	Content   sql.NullString
	ID        int64
	AccountID string
}

func (q *Queries) UpdateTweetContent(ctx context.Context, arg UpdateTweetContentParams) error {
	_, err := q.db.ExecContext(ctx, updateTweetContent, arg.Content, arg.ID, arg.AccountID)
	return err
}
