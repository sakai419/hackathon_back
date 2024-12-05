// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tweets.sql

package sqlcgen

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/sqlc-dev/pqtype"
)

const createTweet = `-- name: CreateTweet :one
INSERT INTO tweets (
    account_id, content, code, media
) VALUES ($1, $2, $3, $4)
RETURNING id
`

type CreateTweetParams struct {
	AccountID string
	Content   sql.NullString
	Code      pqtype.NullRawMessage
	Media     pqtype.NullRawMessage
}

func (q *Queries) CreateTweet(ctx context.Context, arg CreateTweetParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createTweet,
		arg.AccountID,
		arg.Content,
		arg.Code,
		arg.Media,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createTweetAsQuote = `-- name: CreateTweetAsQuote :one
INSERT INTO tweets (
    account_id, is_quote, content, code, media
) VALUES ($1, TRUE, $2, $3, $4)
RETURNING id
`

type CreateTweetAsQuoteParams struct {
	AccountID string
	Content   sql.NullString
	Code      pqtype.NullRawMessage
	Media     pqtype.NullRawMessage
}

func (q *Queries) CreateTweetAsQuote(ctx context.Context, arg CreateTweetAsQuoteParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createTweetAsQuote,
		arg.AccountID,
		arg.Content,
		arg.Code,
		arg.Media,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createTweetAsReply = `-- name: CreateTweetAsReply :one
INSERT INTO tweets (
    account_id, is_reply, content, code, media
) VALUES ($1, TRUE, $2, $3, $4)
RETURNING id
`

type CreateTweetAsReplyParams struct {
	AccountID string
	Content   sql.NullString
	Code      pqtype.NullRawMessage
	Media     pqtype.NullRawMessage
}

func (q *Queries) CreateTweetAsReply(ctx context.Context, arg CreateTweetAsReplyParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createTweetAsReply,
		arg.AccountID,
		arg.Content,
		arg.Code,
		arg.Media,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteTweet = `-- name: DeleteTweet :execresult
DELETE FROM tweets WHERE id = $1
`

func (q *Queries) DeleteTweet(ctx context.Context, id int64) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteTweet, id)
}

const getAccountIDByTweetID = `-- name: GetAccountIDByTweetID :one
SELECT account_id FROM tweets WHERE id = $1
`

func (q *Queries) GetAccountIDByTweetID(ctx context.Context, id int64) (string, error) {
	row := q.db.QueryRowContext(ctx, getAccountIDByTweetID, id)
	var account_id string
	err := row.Scan(&account_id)
	return account_id, err
}

const getPinnedTweetForAccount = `-- name: GetPinnedTweetForAccount :one
SELECT id, account_id, is_pinned, content, code, likes_count, replies_count, retweets_count, is_reply, is_quote, media, created_at FROM tweets
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
		&i.IsReply,
		&i.IsQuote,
		&i.Media,
		&i.CreatedAt,
	)
	return i, err
}

const getPinnedTweetID = `-- name: GetPinnedTweetID :one
SELECT id FROM tweets
WHERE account_id = $1 AND is_pinned = TRUE
LIMIT 1
`

func (q *Queries) GetPinnedTweetID(ctx context.Context, accountID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getPinnedTweetID, accountID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getRecentTweetInfos = `-- name: GetRecentTweetInfos :many
SELECT
    t.id, t.account_id, t.is_pinned, t.content, t.code, t.likes_count, t.replies_count, t.retweets_count, t.is_reply, t.is_quote, t.media, t.created_at,
    COALESCE(l.has_liked, FALSE) AS has_liked,
    COALESCE(r.has_retweeted, FALSE) AS has_retweeted
FROM tweets AS t
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_liked
    FROM likes
    WHERE liking_account_id = $3
) AS l ON t.id = l.original_tweet_id
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_retweeted
    FROM retweets
    WHERE retweeting_account_id = $3
) AS r ON t.id = r.original_tweet_id
WHERE t.account_id != $3
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2
`

type GetRecentTweetInfosParams struct {
	Limit           int32
	Offset          int32
	ClientAccountID string
}

type GetRecentTweetInfosRow struct {
	ID            int64
	AccountID     string
	IsPinned      bool
	Content       sql.NullString
	Code          pqtype.NullRawMessage
	LikesCount    int32
	RepliesCount  int32
	RetweetsCount int32
	IsReply       bool
	IsQuote       bool
	Media         pqtype.NullRawMessage
	CreatedAt     time.Time
	HasLiked      bool
	HasRetweeted  bool
}

func (q *Queries) GetRecentTweetInfos(ctx context.Context, arg GetRecentTweetInfosParams) ([]GetRecentTweetInfosRow, error) {
	rows, err := q.db.QueryContext(ctx, getRecentTweetInfos, arg.Limit, arg.Offset, arg.ClientAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRecentTweetInfosRow
	for rows.Next() {
		var i GetRecentTweetInfosRow
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.IsPinned,
			&i.Content,
			&i.Code,
			&i.LikesCount,
			&i.RepliesCount,
			&i.RetweetsCount,
			&i.IsReply,
			&i.IsQuote,
			&i.Media,
			&i.CreatedAt,
			&i.HasLiked,
			&i.HasRetweeted,
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

const getRecentTweetMetadatas = `-- name: GetRecentTweetMetadatas :many
SELECT
    t.id,
    t.account_id,
    t.likes_count,
    t.retweets_count,
    t.replies_count,
    l.label1,
    l.label2,
    l.label3
FROM tweets AS t
INNER JOIN labels AS l ON t.id = l.tweet_id
INNER JOIN settings AS s ON t.account_id = s.account_id
INNER JOIN accounts AS a ON t.account_id = a.id
WHERE a.is_suspended = FALSE AND a.id != $3
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2
`

type GetRecentTweetMetadatasParams struct {
	Limit           int32
	Offset          int32
	ClientAccountID string
}

type GetRecentTweetMetadatasRow struct {
	ID            int64
	AccountID     string
	LikesCount    int32
	RetweetsCount int32
	RepliesCount  int32
	Label1        NullTweetLabel
	Label2        NullTweetLabel
	Label3        NullTweetLabel
}

func (q *Queries) GetRecentTweetMetadatas(ctx context.Context, arg GetRecentTweetMetadatasParams) ([]GetRecentTweetMetadatasRow, error) {
	rows, err := q.db.QueryContext(ctx, getRecentTweetMetadatas, arg.Limit, arg.Offset, arg.ClientAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRecentTweetMetadatasRow
	for rows.Next() {
		var i GetRecentTweetMetadatasRow
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.LikesCount,
			&i.RetweetsCount,
			&i.RepliesCount,
			&i.Label1,
			&i.Label2,
			&i.Label3,
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

const getTweetCountByAccountID = `-- name: GetTweetCountByAccountID :one
SELECT COUNT(*) FROM tweets
WHERE account_id = $1
`

func (q *Queries) GetTweetCountByAccountID(ctx context.Context, accountID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTweetCountByAccountID, accountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getTweetInfosByAccountID = `-- name: GetTweetInfosByAccountID :many
SELECT
    t.id, t.account_id, t.is_pinned, t.content, t.code, t.likes_count, t.replies_count, t.retweets_count, t.is_reply, t.is_quote, t.media, t.created_at,
    COALESCE(l.has_liked, FALSE) AS has_liked,
    COALESCE(r.has_retweeted, FALSE) AS has_retweeted
FROM tweets AS t
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_liked
    FROM likes
    WHERE liking_account_id = $3
) AS l ON t.id = l.original_tweet_id
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_retweeted
    FROM retweets
    WHERE retweeting_account_id = $3
) AS r ON t.id = r.original_tweet_id
WHERE t.account_id = $4
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2
`

type GetTweetInfosByAccountIDParams struct {
	Limit           int32
	Offset          int32
	ClientAccountID string
	TargetAccountID string
}

type GetTweetInfosByAccountIDRow struct {
	ID            int64
	AccountID     string
	IsPinned      bool
	Content       sql.NullString
	Code          pqtype.NullRawMessage
	LikesCount    int32
	RepliesCount  int32
	RetweetsCount int32
	IsReply       bool
	IsQuote       bool
	Media         pqtype.NullRawMessage
	CreatedAt     time.Time
	HasLiked      bool
	HasRetweeted  bool
}

func (q *Queries) GetTweetInfosByAccountID(ctx context.Context, arg GetTweetInfosByAccountIDParams) ([]GetTweetInfosByAccountIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getTweetInfosByAccountID,
		arg.Limit,
		arg.Offset,
		arg.ClientAccountID,
		arg.TargetAccountID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTweetInfosByAccountIDRow
	for rows.Next() {
		var i GetTweetInfosByAccountIDRow
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.IsPinned,
			&i.Content,
			&i.Code,
			&i.LikesCount,
			&i.RepliesCount,
			&i.RetweetsCount,
			&i.IsReply,
			&i.IsQuote,
			&i.Media,
			&i.CreatedAt,
			&i.HasLiked,
			&i.HasRetweeted,
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

const getTweetInfosByIDs = `-- name: GetTweetInfosByIDs :many
SELECT
    t.id, t.account_id, t.is_pinned, t.content, t.code, t.likes_count, t.replies_count, t.retweets_count, t.is_reply, t.is_quote, t.media, t.created_at,
    COALESCE(l.has_liked, FALSE) AS has_liked,
    COALESCE(r.has_retweeted, FALSE) AS has_retweeted
FROM tweets AS t
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_liked
    FROM likes
    WHERE liking_account_id = $1
) AS l ON t.id = l.original_tweet_id
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_retweeted
    FROM retweets
    WHERE retweeting_account_id = $1
) AS r ON t.id = r.original_tweet_id
WHERE t.id = ANY($2::BIGINT[])
`

type GetTweetInfosByIDsParams struct {
	ClientAccountID string
	TweetIds        []int64
}

type GetTweetInfosByIDsRow struct {
	ID            int64
	AccountID     string
	IsPinned      bool
	Content       sql.NullString
	Code          pqtype.NullRawMessage
	LikesCount    int32
	RepliesCount  int32
	RetweetsCount int32
	IsReply       bool
	IsQuote       bool
	Media         pqtype.NullRawMessage
	CreatedAt     time.Time
	HasLiked      bool
	HasRetweeted  bool
}

func (q *Queries) GetTweetInfosByIDs(ctx context.Context, arg GetTweetInfosByIDsParams) ([]GetTweetInfosByIDsRow, error) {
	rows, err := q.db.QueryContext(ctx, getTweetInfosByIDs, arg.ClientAccountID, pq.Array(arg.TweetIds))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTweetInfosByIDsRow
	for rows.Next() {
		var i GetTweetInfosByIDsRow
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.IsPinned,
			&i.Content,
			&i.Code,
			&i.LikesCount,
			&i.RepliesCount,
			&i.RetweetsCount,
			&i.IsReply,
			&i.IsQuote,
			&i.Media,
			&i.CreatedAt,
			&i.HasLiked,
			&i.HasRetweeted,
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

const searchTweets = `-- name: SearchTweets :many
SELECT id, account_id, is_pinned, content, code, likes_count, replies_count, retweets_count, is_reply, is_quote, media, created_at FROM tweets
WHERE content LIKE $1 OR code LIKE $2
ORDER BY created_at DESC
LIMIT $3 OFFSET $4
`

type SearchTweetsParams struct {
	Content sql.NullString
	Code    pqtype.NullRawMessage
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
			&i.IsReply,
			&i.IsQuote,
			&i.Media,
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

const searchTweetsOrderByCreatedAt = `-- name: SearchTweetsOrderByCreatedAt :many
SELECT
    t.id, t.account_id, t.is_pinned, t.content, t.code, t.likes_count, t.replies_count, t.retweets_count, t.is_reply, t.is_quote, t.media, t.created_at,
    COALESCE(l.has_liked, FALSE) AS has_liked,
    COALESCE(r.has_retweeted, FALSE) AS has_retweeted
FROM tweets AS t
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_liked
    FROM likes
    WHERE liking_account_id = $3
) AS l ON t.id = l.original_tweet_id
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_retweeted
    FROM retweets
    WHERE retweeting_account_id = $3
) AS r ON t.id = r.original_tweet_id
WHERE
    t.content ILIKE CONCAT('%', $4::VARCHAR, '%')
    OR t.code->>'Content' ILIKE CONCAT('%', $4::VARCHAR, '%')
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2
`

type SearchTweetsOrderByCreatedAtParams struct {
	Limit           int32
	Offset          int32
	ClientAccountID string
	Keyword         string
}

type SearchTweetsOrderByCreatedAtRow struct {
	ID            int64
	AccountID     string
	IsPinned      bool
	Content       sql.NullString
	Code          pqtype.NullRawMessage
	LikesCount    int32
	RepliesCount  int32
	RetweetsCount int32
	IsReply       bool
	IsQuote       bool
	Media         pqtype.NullRawMessage
	CreatedAt     time.Time
	HasLiked      bool
	HasRetweeted  bool
}

func (q *Queries) SearchTweetsOrderByCreatedAt(ctx context.Context, arg SearchTweetsOrderByCreatedAtParams) ([]SearchTweetsOrderByCreatedAtRow, error) {
	rows, err := q.db.QueryContext(ctx, searchTweetsOrderByCreatedAt,
		arg.Limit,
		arg.Offset,
		arg.ClientAccountID,
		arg.Keyword,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchTweetsOrderByCreatedAtRow
	for rows.Next() {
		var i SearchTweetsOrderByCreatedAtRow
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.IsPinned,
			&i.Content,
			&i.Code,
			&i.LikesCount,
			&i.RepliesCount,
			&i.RetweetsCount,
			&i.IsReply,
			&i.IsQuote,
			&i.Media,
			&i.CreatedAt,
			&i.HasLiked,
			&i.HasRetweeted,
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

const searchTweetsOrderByEngagementScore = `-- name: SearchTweetsOrderByEngagementScore :many
SELECT
    t.id, t.account_id, t.is_pinned, t.content, t.code, t.likes_count, t.replies_count, t.retweets_count, t.is_reply, t.is_quote, t.media, t.created_at,
    COALESCE(l.has_liked, FALSE) AS has_liked,
    COALESCE(r.has_retweeted, FALSE) AS has_retweeted
FROM tweets AS t
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_liked
    FROM likes
    WHERE liking_account_id = $3
) AS l ON t.id = l.original_tweet_id
LEFT JOIN (
    SELECT
        original_tweet_id,
        TRUE AS has_retweeted
    FROM retweets
    WHERE retweeting_account_id = $3
) AS r ON t.id = r.original_tweet_id
WHERE
    t.content ILIKE CONCAT('%', $4::VARCHAR, '%')
    OR t.code->>'Content' ILIKE CONCAT('%', $4::VARCHAR, '%')
ORDER BY
    (t.likes_count * 30 + t.retweets_count * 20 + t.replies_count * 1) DESC,
    t.created_at DESC
LIMIT $1 OFFSET $2
`

type SearchTweetsOrderByEngagementScoreParams struct {
	Limit           int32
	Offset          int32
	ClientAccountID string
	Keyword         string
}

type SearchTweetsOrderByEngagementScoreRow struct {
	ID            int64
	AccountID     string
	IsPinned      bool
	Content       sql.NullString
	Code          pqtype.NullRawMessage
	LikesCount    int32
	RepliesCount  int32
	RetweetsCount int32
	IsReply       bool
	IsQuote       bool
	Media         pqtype.NullRawMessage
	CreatedAt     time.Time
	HasLiked      bool
	HasRetweeted  bool
}

func (q *Queries) SearchTweetsOrderByEngagementScore(ctx context.Context, arg SearchTweetsOrderByEngagementScoreParams) ([]SearchTweetsOrderByEngagementScoreRow, error) {
	rows, err := q.db.QueryContext(ctx, searchTweetsOrderByEngagementScore,
		arg.Limit,
		arg.Offset,
		arg.ClientAccountID,
		arg.Keyword,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchTweetsOrderByEngagementScoreRow
	for rows.Next() {
		var i SearchTweetsOrderByEngagementScoreRow
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.IsPinned,
			&i.Content,
			&i.Code,
			&i.LikesCount,
			&i.RepliesCount,
			&i.RetweetsCount,
			&i.IsReply,
			&i.IsQuote,
			&i.Media,
			&i.CreatedAt,
			&i.HasLiked,
			&i.HasRetweeted,
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

const setTweetAsPinned = `-- name: SetTweetAsPinned :execresult
UPDATE tweets
SET is_pinned = TRUE
WHERE id = $1 AND account_id = $2
`

type SetTweetAsPinnedParams struct {
	ID        int64
	AccountID string
}

func (q *Queries) SetTweetAsPinned(ctx context.Context, arg SetTweetAsPinnedParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, setTweetAsPinned, arg.ID, arg.AccountID)
}

const unsetTweetAsPinned = `-- name: UnsetTweetAsPinned :execresult
UPDATE tweets
SET is_pinned = FALSE
WHERE id = $1 AND account_id = $2
`

type UnsetTweetAsPinnedParams struct {
	ID        int64
	AccountID string
}

func (q *Queries) UnsetTweetAsPinned(ctx context.Context, arg UnsetTweetAsPinnedParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, unsetTweetAsPinned, arg.ID, arg.AccountID)
}
