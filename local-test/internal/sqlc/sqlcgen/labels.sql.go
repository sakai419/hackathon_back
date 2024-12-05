// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: labels.sql

package sqlcgen

import (
	"context"
)

const createLabel = `-- name: CreateLabel :exec
INSERT INTO labels (tweet_id, label1, label2, label3)
VALUES ($1, $2, $3, $4)
`

type CreateLabelParams struct {
	TweetID int64
	Label1  NullTweetLabel
	Label2  NullTweetLabel
	Label3  NullTweetLabel
}

func (q *Queries) CreateLabel(ctx context.Context, arg CreateLabelParams) error {
	_, err := q.db.ExecContext(ctx, createLabel,
		arg.TweetID,
		arg.Label1,
		arg.Label2,
		arg.Label3,
	)
	return err
}

const deleteLabel = `-- name: DeleteLabel :exec
DELETE FROM labels
WHERE tweet_id = $1
`

func (q *Queries) DeleteLabel(ctx context.Context, tweetID int64) error {
	_, err := q.db.ExecContext(ctx, deleteLabel, tweetID)
	return err
}

const getLabelsByTweetID = `-- name: GetLabelsByTweetID :one
SELECT tweet_id, label1, label2, label3, created_at FROM labels
WHERE tweet_id = $1
`

func (q *Queries) GetLabelsByTweetID(ctx context.Context, tweetID int64) (Label, error) {
	row := q.db.QueryRowContext(ctx, getLabelsByTweetID, tweetID)
	var i Label
	err := row.Scan(
		&i.TweetID,
		&i.Label1,
		&i.Label2,
		&i.Label3,
		&i.CreatedAt,
	)
	return i, err
}

const getLikedTweetLabelsCount = `-- name: GetLikedTweetLabelsCount :many
WITH liked_tweets AS (
    SELECT l.original_tweet_id
    FROM likes l
    WHERE l.liking_account_id = $1
    ORDER BY l.created_at DESC
    LIMIT 100
)
SELECT
    label,
    COUNT(*) AS label_count
FROM (
    SELECT label1 AS label FROM labels WHERE tweet_id IN (SELECT original_tweet_id FROM liked_tweets)
    UNION ALL
    SELECT label2 AS label FROM labels WHERE tweet_id IN (SELECT original_tweet_id FROM liked_tweets)
    UNION ALL
    SELECT label3 AS label FROM labels WHERE tweet_id IN (SELECT original_tweet_id FROM liked_tweets)
) labels_combined
GROUP BY label
`

type GetLikedTweetLabelsCountRow struct {
	Label      NullTweetLabel
	LabelCount int64
}

func (q *Queries) GetLikedTweetLabelsCount(ctx context.Context, likingAccountID string) ([]GetLikedTweetLabelsCountRow, error) {
	rows, err := q.db.QueryContext(ctx, getLikedTweetLabelsCount, likingAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLikedTweetLabelsCountRow
	for rows.Next() {
		var i GetLikedTweetLabelsCountRow
		if err := rows.Scan(&i.Label, &i.LabelCount); err != nil {
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

const getRecentLabels = `-- name: GetRecentLabels :many
WITH all_labels AS (
    SELECT label1 AS label FROM labels
    UNION ALL
    SELECT label2 AS label FROM labels
    UNION ALL
    SELECT label3 AS label FROM labels
)
SELECT
    label,
    COUNT(*) AS label_count
FROM
    all_labels
WHERE
    label IS NOT NULL
GROUP BY
    label
ORDER BY
    label_count DESC
LIMIT $1
`

type GetRecentLabelsRow struct {
	Label      NullTweetLabel
	LabelCount int64
}

func (q *Queries) GetRecentLabels(ctx context.Context, limit int32) ([]GetRecentLabelsRow, error) {
	rows, err := q.db.QueryContext(ctx, getRecentLabels, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRecentLabelsRow
	for rows.Next() {
		var i GetRecentLabelsRow
		if err := rows.Scan(&i.Label, &i.LabelCount); err != nil {
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

const getTweetsByLabel = `-- name: GetTweetsByLabel :many
SELECT t.id, t.account_id, t.is_pinned, t.content, t.code, t.likes_count, t.replies_count, t.retweets_count, t.is_reply, t.is_quote, t.media, t.created_at FROM tweets t
JOIN labels l ON t.id = l.tweet_id
WHERE l.label1 = $1 OR l.label2 = $2 OR l.label3 = $3
ORDER BY t.created_at DESC
LIMIT $4 OFFSET $5
`

type GetTweetsByLabelParams struct {
	Label1 NullTweetLabel
	Label2 NullTweetLabel
	Label3 NullTweetLabel
	Limit  int32
	Offset int32
}

func (q *Queries) GetTweetsByLabel(ctx context.Context, arg GetTweetsByLabelParams) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, getTweetsByLabel,
		arg.Label1,
		arg.Label2,
		arg.Label3,
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

const getTweetsWithoutLabels = `-- name: GetTweetsWithoutLabels :many
SELECT t.id, t.account_id, t.is_pinned, t.content, t.code, t.likes_count, t.replies_count, t.retweets_count, t.is_reply, t.is_quote, t.media, t.created_at FROM tweets t
LEFT JOIN labels l ON t.id = l.tweet_id
WHERE l.tweet_id IS NULL
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2
`

type GetTweetsWithoutLabelsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetTweetsWithoutLabels(ctx context.Context, arg GetTweetsWithoutLabelsParams) ([]Tweet, error) {
	rows, err := q.db.QueryContext(ctx, getTweetsWithoutLabels, arg.Limit, arg.Offset)
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

const updateLabels = `-- name: UpdateLabels :exec
UPDATE labels
SET label1 = $1, label2 = $2, label3 = $3
WHERE tweet_id = $4
`

type UpdateLabelsParams struct {
	Label1  NullTweetLabel
	Label2  NullTweetLabel
	Label3  NullTweetLabel
	TweetID int64
}

func (q *Queries) UpdateLabels(ctx context.Context, arg UpdateLabelsParams) error {
	_, err := q.db.ExecContext(ctx, updateLabels,
		arg.Label1,
		arg.Label2,
		arg.Label3,
		arg.TweetID,
	)
	return err
}
