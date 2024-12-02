// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: follows.sql

package sqlcgen

import (
	"context"
	"database/sql"
	"time"
)

const acceptFollowRequest = `-- name: AcceptFollowRequest :execresult
UPDATE follows
SET status = 'accepted'
WHERE follower_account_id = $1 AND following_account_id = $2 AND status = 'pending'
`

type AcceptFollowRequestParams struct {
	FollowerAccountID  string
	FollowingAccountID string
}

func (q *Queries) AcceptFollowRequest(ctx context.Context, arg AcceptFollowRequestParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, acceptFollowRequest, arg.FollowerAccountID, arg.FollowingAccountID)
}

const checkIsFollowed = `-- name: CheckIsFollowed :one
SELECT EXISTS(
    SELECT 1
    FROM follows
    WHERE follower_account_id = $1 AND following_account_id = $2 AND status = 'accepted'
)
`

type CheckIsFollowedParams struct {
	FollowerAccountID  string
	FollowingAccountID string
}

func (q *Queries) CheckIsFollowed(ctx context.Context, arg CheckIsFollowedParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkIsFollowed, arg.FollowerAccountID, arg.FollowingAccountID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createFollow = `-- name: CreateFollow :exec
INSERT INTO follows (follower_account_id, following_account_id, status)
VALUES ($1, $2, 'accepted')
`

type CreateFollowParams struct {
	FollowerAccountID  string
	FollowingAccountID string
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) error {
	_, err := q.db.ExecContext(ctx, createFollow, arg.FollowerAccountID, arg.FollowingAccountID)
	return err
}

const createFollowRequest = `-- name: CreateFollowRequest :exec
INSERT INTO follows (follower_account_id, following_account_id, status)
VALUES ($1, $2, 'pending')
`

type CreateFollowRequestParams struct {
	FollowerAccountID  string
	FollowingAccountID string
}

func (q *Queries) CreateFollowRequest(ctx context.Context, arg CreateFollowRequestParams) error {
	_, err := q.db.ExecContext(ctx, createFollowRequest, arg.FollowerAccountID, arg.FollowingAccountID)
	return err
}

const deleteFollow = `-- name: DeleteFollow :execresult
DELETE FROM follows
WHERE follower_account_id = $1 AND following_account_id = $2 AND status = 'accepted'
`

type DeleteFollowParams struct {
	FollowerAccountID  string
	FollowingAccountID string
}

func (q *Queries) DeleteFollow(ctx context.Context, arg DeleteFollowParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteFollow, arg.FollowerAccountID, arg.FollowingAccountID)
}

const deleteFollowRequest = `-- name: DeleteFollowRequest :execresult
DELETE FROM follows
WHERE follower_account_id = $1 AND following_account_id = $2 AND status = 'pending'
`

type DeleteFollowRequestParams struct {
	FollowerAccountID  string
	FollowingAccountID string
}

func (q *Queries) DeleteFollowRequest(ctx context.Context, arg DeleteFollowRequestParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteFollowRequest, arg.FollowerAccountID, arg.FollowingAccountID)
}

const getFollowCounts = `-- name: GetFollowCounts :one
SELECT
    COUNT(CASE WHEN following_account_id = $1 AND status = 'accepted' THEN 1 END) AS follower_count,
    COUNT(CASE WHEN follower_account_id = $1 AND status = 'accepted' THEN 1 END) AS following_count
FROM follows
`

type GetFollowCountsRow struct {
	FollowerCount  int64
	FollowingCount int64
}

func (q *Queries) GetFollowCounts(ctx context.Context, followingAccountID string) (GetFollowCountsRow, error) {
	row := q.db.QueryRowContext(ctx, getFollowCounts, followingAccountID)
	var i GetFollowCountsRow
	err := row.Scan(&i.FollowerCount, &i.FollowingCount)
	return i, err
}

const getFollowRequestCount = `-- name: GetFollowRequestCount :one
SELECT COUNT(*)
FROM follows
WHERE following_account_id = $1 AND status = 'pending'
`

func (q *Queries) GetFollowRequestCount(ctx context.Context, followingAccountID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getFollowRequestCount, followingAccountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getFollowStatus = `-- name: GetFollowStatus :one
SELECT status
FROM follows
WHERE follower_account_id = $1 AND following_account_id = $2
`

type GetFollowStatusParams struct {
	FollowerAccountID  string
	FollowingAccountID string
}

func (q *Queries) GetFollowStatus(ctx context.Context, arg GetFollowStatusParams) (FollowStatus, error) {
	row := q.db.QueryRowContext(ctx, getFollowStatus, arg.FollowerAccountID, arg.FollowingAccountID)
	var status FollowStatus
	err := row.Scan(&status)
	return status, err
}

const getFollowerAccountIDs = `-- name: GetFollowerAccountIDs :many
SELECT follower_account_id
FROM follows
WHERE following_account_id = $1 AND status = 'accepted'
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

const getFollowingAccountIDs = `-- name: GetFollowingAccountIDs :many
SELECT following_account_id
FROM follows
WHERE follower_account_id = $1 AND status = 'accepted'
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetFollowingAccountIDsParams struct {
	FollowerAccountID string
	Limit             int32
	Offset            int32
}

func (q *Queries) GetFollowingAccountIDs(ctx context.Context, arg GetFollowingAccountIDsParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getFollowingAccountIDs, arg.FollowerAccountID, arg.Limit, arg.Offset)
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

const isPrivateAndNotFollowing = `-- name: IsPrivateAndNotFollowing :one
SELECT
    CASE
        WHEN s.is_private = TRUE AND f.follower_account_id IS NULL THEN TRUE
        ELSE FALSE
    END AS is_private_and_not_following
FROM
    settings AS s
LEFT JOIN
    follows AS f
ON
    s.account_id = f.following_account_id
    AND f.follower_account_id = $1
WHERE
    s.account_id = $2
`

type IsPrivateAndNotFollowingParams struct {
	ClientAccountID string
	TargetAccountID string
}

func (q *Queries) IsPrivateAndNotFollowing(ctx context.Context, arg IsPrivateAndNotFollowingParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isPrivateAndNotFollowing, arg.ClientAccountID, arg.TargetAccountID)
	var is_private_and_not_following bool
	err := row.Scan(&is_private_and_not_following)
	return is_private_and_not_following, err
}

const listPendingRequests = `-- name: ListPendingRequests :many
SELECT follower_account_id, following_account_id, created_at
FROM follows
WHERE following_account_id = $1 AND status = 'pending'
`

type ListPendingRequestsRow struct {
	FollowerAccountID  string
	FollowingAccountID string
	CreatedAt          time.Time
}

func (q *Queries) ListPendingRequests(ctx context.Context, followingAccountID string) ([]ListPendingRequestsRow, error) {
	rows, err := q.db.QueryContext(ctx, listPendingRequests, followingAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPendingRequestsRow
	for rows.Next() {
		var i ListPendingRequestsRow
		if err := rows.Scan(&i.FollowerAccountID, &i.FollowingAccountID, &i.CreatedAt); err != nil {
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