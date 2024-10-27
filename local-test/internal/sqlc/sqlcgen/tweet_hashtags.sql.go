// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tweet_hashtags.sql

package sqlcgen

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/sqlc-dev/pqtype"
)

const associateTweetWithHashtags = `-- name: AssociateTweetWithHashtags :exec
INSERT INTO tweet_hashtags (tweet_id, hashtag_id)
VALUES
    ($1, unnest($2::bigint[]))
`

type AssociateTweetWithHashtagsParams struct {
	TweetID    int64
	HashtagIds []int64
}

func (q *Queries) AssociateTweetWithHashtags(ctx context.Context, arg AssociateTweetWithHashtagsParams) error {
	_, err := q.db.ExecContext(ctx, associateTweetWithHashtags, arg.TweetID, pq.Array(arg.HashtagIds))
	return err
}

const checkTweetHashtagExists = `-- name: CheckTweetHashtagExists :one
SELECT EXISTS(
    SELECT 1 FROM tweet_hashtags
    WHERE tweet_id = $1 AND hashtag_id = $2
) AS hashtag_exists
`

type CheckTweetHashtagExistsParams struct {
	TweetID   int64
	HashtagID int64
}

func (q *Queries) CheckTweetHashtagExists(ctx context.Context, arg CheckTweetHashtagExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkTweetHashtagExists, arg.TweetID, arg.HashtagID)
	var hashtag_exists bool
	err := row.Scan(&hashtag_exists)
	return hashtag_exists, err
}

const deleteAllHashtagsForTweet = `-- name: DeleteAllHashtagsForTweet :exec
DELETE FROM tweet_hashtags
WHERE tweet_id = $1
`

func (q *Queries) DeleteAllHashtagsForTweet(ctx context.Context, tweetID int64) error {
	_, err := q.db.ExecContext(ctx, deleteAllHashtagsForTweet, tweetID)
	return err
}

const getHashtagsByTweetID = `-- name: GetHashtagsByTweetID :many
SELECT h.id, h.tag, h.created_at
FROM hashtags h
JOIN tweet_hashtags th ON h.id = th.hashtag_id
WHERE th.tweet_id = $1
`

func (q *Queries) GetHashtagsByTweetID(ctx context.Context, tweetID int64) ([]Hashtag, error) {
	rows, err := q.db.QueryContext(ctx, getHashtagsByTweetID, tweetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Hashtag
	for rows.Next() {
		var i Hashtag
		if err := rows.Scan(&i.ID, &i.Tag, &i.CreatedAt); err != nil {
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

const getMostUsedHashtags = `-- name: GetMostUsedHashtags :many
SELECT h.id, h.tag, h.created_at, COUNT(th.tweet_id) as usage_count
FROM hashtags h
JOIN tweet_hashtags th ON h.id = th.hashtag_id
GROUP BY h.id
ORDER BY usage_count DESC
LIMIT $1 OFFSET $2
`

type GetMostUsedHashtagsParams struct {
	Limit  int32
	Offset int32
}

type GetMostUsedHashtagsRow struct {
	ID         int64
	Tag        string
	CreatedAt  time.Time
	UsageCount int64
}

func (q *Queries) GetMostUsedHashtags(ctx context.Context, arg GetMostUsedHashtagsParams) ([]GetMostUsedHashtagsRow, error) {
	rows, err := q.db.QueryContext(ctx, getMostUsedHashtags, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMostUsedHashtagsRow
	for rows.Next() {
		var i GetMostUsedHashtagsRow
		if err := rows.Scan(
			&i.ID,
			&i.Tag,
			&i.CreatedAt,
			&i.UsageCount,
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

const getRecentTweetsWithHashtag = `-- name: GetRecentTweetsWithHashtag :many
SELECT t.id, t.account_id, t.is_pinned, t.content, t.code, t.likes_count, t.replies_count, t.retweets_count, t.is_reply, t.is_quote, t.media, t.created_at, t.updated_at, h.tag
FROM tweets t
JOIN tweet_hashtags th ON t.id = th.tweet_id
JOIN hashtags h ON th.hashtag_id = h.id
WHERE h.tag = $1
ORDER BY t.created_at DESC
LIMIT $2 OFFSET $3
`

type GetRecentTweetsWithHashtagParams struct {
	Tag    string
	Limit  int32
	Offset int32
}

type GetRecentTweetsWithHashtagRow struct {
	ID            int64
	AccountID     string
	IsPinned      bool
	Content       sql.NullString
	Code          sql.NullString
	LikesCount    int32
	RepliesCount  int32
	RetweetsCount int32
	IsReply       bool
	IsQuote       bool
	Media         pqtype.NullRawMessage
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Tag           string
}

func (q *Queries) GetRecentTweetsWithHashtag(ctx context.Context, arg GetRecentTweetsWithHashtagParams) ([]GetRecentTweetsWithHashtagRow, error) {
	rows, err := q.db.QueryContext(ctx, getRecentTweetsWithHashtag, arg.Tag, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRecentTweetsWithHashtagRow
	for rows.Next() {
		var i GetRecentTweetsWithHashtagRow
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
			&i.UpdatedAt,
			&i.Tag,
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

const getTweetCountByHashtagID = `-- name: GetTweetCountByHashtagID :one
SELECT COUNT(DISTINCT tweet_id)
FROM tweet_hashtags
WHERE hashtag_id = $1
`

func (q *Queries) GetTweetCountByHashtagID(ctx context.Context, hashtagID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTweetCountByHashtagID, hashtagID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getTweetsByHashtagID = `-- name: GetTweetsByHashtagID :many
SELECT t.id, t.account_id, t.is_pinned, t.content, t.code, t.likes_count, t.replies_count, t.retweets_count, t.is_reply, t.is_quote, t.media, t.created_at, t.updated_at
FROM tweets t
JOIN tweet_hashtags th ON t.id = th.tweet_id
WHERE th.hashtag_id = $1
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2
`

type GetTweetsByHashtagIDParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetTweetsByHashtagID(ctx context.Context, arg GetTweetsByHashtagIDParams) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, getTweetsByHashtagID, arg.Limit, arg.Offset)
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
