// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: follows.sql

package sqlcgen

import (
	"context"
	"database/sql"
)

const checkFollowExists = `-- name: CheckFollowExists :one
SELECT EXISTS(
    SELECT 1 FROM follows
    WHERE follower_account_id = $1 AND following_account_id = $2
) AS is_following
`

type CheckFollowExistsParams struct {
	FollowerAccountID  string
	FollowingAccountID string
}

func (q *Queries) CheckFollowExists(ctx context.Context, arg CheckFollowExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkFollowExists, arg.FollowerAccountID, arg.FollowingAccountID)
	var is_following bool
	err := row.Scan(&is_following)
	return is_following, err
}

const createFollow = `-- name: CreateFollow :exec
INSERT INTO follows (follower_account_id, following_account_id)
VALUES ($1, $2)
`

type CreateFollowParams struct {
	FollowerAccountID  string
	FollowingAccountID string
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) error {
	_, err := q.db.ExecContext(ctx, createFollow, arg.FollowerAccountID, arg.FollowingAccountID)
	return err
}

const deleteFollow = `-- name: DeleteFollow :execresult
DELETE FROM follows
WHERE follower_account_id = $1 AND following_account_id = $2
`

type DeleteFollowParams struct {
	FollowerAccountID  string
	FollowingAccountID string
}

func (q *Queries) DeleteFollow(ctx context.Context, arg DeleteFollowParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteFollow, arg.FollowerAccountID, arg.FollowingAccountID)
}

const getFollowSuggestions = `-- name: GetFollowSuggestions :many
SELECT DISTINCT f2.following_account_id
FROM follows f1
JOIN follows f2 ON f1.following_account_id = f2.follower_account_id
WHERE f1.follower_account_id = $1
    AND f2.following_account_id != f1.follower_account_id
    AND NOT EXISTS (
        SELECT 1 FROM follows f3
        WHERE f3.follower_account_id = f1.follower_account_id
            AND f3.following_account_id = f2.following_account_id
    )
LIMIT $2 OFFSET $3
`

type GetFollowSuggestionsParams struct {
	FollowerAccountID string
	Limit             int32
	Offset            int32
}

func (q *Queries) GetFollowSuggestions(ctx context.Context, arg GetFollowSuggestionsParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getFollowSuggestions, arg.FollowerAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var following_account_id string
		if err := rows.Scan(&following_account_id); err != nil {
			return nil, err
		}
		items = append(items, following_account_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowerAccountIDs = `-- name: GetFollowerAccountIDs :many
SELECT follower_account_id
FROM follows
WHERE following_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetFollowerAccountIDsParams struct {
	FollowingAccountID string
	Limit              int32
	Offset             int32
}

func (q *Queries) GetFollowerAccountIDs(ctx context.Context, arg GetFollowerAccountIDsParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getFollowerAccountIDs, arg.FollowingAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var follower_account_id string
		if err := rows.Scan(&follower_account_id); err != nil {
			return nil, err
		}
		items = append(items, follower_account_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowerCount = `-- name: GetFollowerCount :one
SELECT COUNT(*) FROM follows
WHERE following_account_id = $1
`

func (q *Queries) GetFollowerCount(ctx context.Context, followingAccountID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getFollowerCount, followingAccountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getFollowing = `-- name: GetFollowing :many
SELECT following_account_id
FROM follows
WHERE follower_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetFollowingParams struct {
	FollowerAccountID string
	Limit             int32
	Offset            int32
}

func (q *Queries) GetFollowing(ctx context.Context, arg GetFollowingParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getFollowing, arg.FollowerAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var following_account_id string
		if err := rows.Scan(&following_account_id); err != nil {
			return nil, err
		}
		items = append(items, following_account_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowingCount = `-- name: GetFollowingCount :one
SELECT COUNT(*) FROM follows
WHERE follower_account_id = $1
`

func (q *Queries) GetFollowingCount(ctx context.Context, followerAccountID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getFollowingCount, followerAccountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getMutualFollows = `-- name: GetMutualFollows :many
SELECT f1.following_account_id
FROM follows f1
JOIN follows f2 ON f1.following_account_id = f2.follower_account_id
WHERE f1.follower_account_id = $1 AND f2.following_account_id = $2
LIMIT $3 OFFSET $4
`

type GetMutualFollowsParams struct {
	FollowerAccountID  string
	FollowingAccountID string
	Limit              int32
	Offset             int32
}

func (q *Queries) GetMutualFollows(ctx context.Context, arg GetMutualFollowsParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getMutualFollows,
		arg.FollowerAccountID,
		arg.FollowingAccountID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var following_account_id string
		if err := rows.Scan(&following_account_id); err != nil {
			return nil, err
		}
		items = append(items, following_account_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
