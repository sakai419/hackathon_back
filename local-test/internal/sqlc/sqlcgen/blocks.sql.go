// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: blocks.sql

package sqlcgen

import (
	"context"
	"database/sql"
)

const checkBlockExists = `-- name: CheckBlockExists :one
SELECT EXISTS(
    SELECT 1 FROM blocks
    WHERE blocker_account_id = $1 AND blocked_account_id = $2
) AS is_blocked
`

type CheckBlockExistsParams struct {
	BlockerAccountID string
	BlockedAccountID string
}

func (q *Queries) CheckBlockExists(ctx context.Context, arg CheckBlockExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkBlockExists, arg.BlockerAccountID, arg.BlockedAccountID)
	var is_blocked bool
	err := row.Scan(&is_blocked)
	return is_blocked, err
}

const createBlock = `-- name: CreateBlock :exec
INSERT INTO blocks (blocker_account_id, blocked_account_id)
VALUES ($1, $2)
`

type CreateBlockParams struct {
	BlockerAccountID string
	BlockedAccountID string
}

func (q *Queries) CreateBlock(ctx context.Context, arg CreateBlockParams) error {
	_, err := q.db.ExecContext(ctx, createBlock, arg.BlockerAccountID, arg.BlockedAccountID)
	return err
}

const deleteBlock = `-- name: DeleteBlock :execresult
DELETE FROM blocks
WHERE blocker_account_id = $1 AND blocked_account_id = $2
`

type DeleteBlockParams struct {
	BlockerAccountID string
	BlockedAccountID string
}

func (q *Queries) DeleteBlock(ctx context.Context, arg DeleteBlockParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteBlock, arg.BlockerAccountID, arg.BlockedAccountID)
}

const getBlockCount = `-- name: GetBlockCount :one
SELECT COUNT(*) FROM blocks
WHERE blocker_account_id = $1
`

func (q *Queries) GetBlockCount(ctx context.Context, blockerAccountID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getBlockCount, blockerAccountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getBlockedAccountIDs = `-- name: GetBlockedAccountIDs :many
SELECT blocked_account_id
FROM blocks
WHERE blocker_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetBlockedAccountIDsParams struct {
	BlockerAccountID string
	Limit            int32
	Offset           int32
}

func (q *Queries) GetBlockedAccountIDs(ctx context.Context, arg GetBlockedAccountIDsParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getBlockedAccountIDs, arg.BlockerAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var blocked_account_id string
		if err := rows.Scan(&blocked_account_id); err != nil {
			return nil, err
		}
		items = append(items, blocked_account_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
