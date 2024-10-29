// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accounts.sql

package sqlcgen

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createAccount = `-- name: CreateAccount :exec
INSERT INTO accounts (id, user_id, user_name)
VALUES ($1, $2, $3)
`

type CreateAccountParams struct {
	ID       string
	UserID   string
	UserName string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) error {
	_, err := q.db.ExecContext(ctx, createAccount, arg.ID, arg.UserID, arg.UserName)
	return err
}

const deleteAccount = `-- name: DeleteAccount :execresult
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteAccount, id)
}

const getAccountIDByUserID = `-- name: GetAccountIDByUserID :one
SELECT id FROM accounts
WHERE user_id = $1
`

func (q *Queries) GetAccountIDByUserID(ctx context.Context, userID string) (string, error) {
	row := q.db.QueryRowContext(ctx, getAccountIDByUserID, userID)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getAccountInfo = `-- name: GetAccountInfo :one
SELECT a.is_suspended, a.is_admin, s.is_private
FROM accounts a
JOIN settings s ON a.id = s.account_id
WHERE a.id = $1
`

type GetAccountInfoRow struct {
	IsSuspended bool
	IsAdmin     bool
	IsPrivate   sql.NullBool
}

func (q *Queries) GetAccountInfo(ctx context.Context, id string) (GetAccountInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getAccountInfo, id)
	var i GetAccountInfoRow
	err := row.Scan(&i.IsSuspended, &i.IsAdmin, &i.IsPrivate)
	return i, err
}

const getUserInfos = `-- name: GetUserInfos :many
SELECT a.id, a.user_id, a.user_name, p.bio, p.profile_image_url
FROM accounts a
JOIN profiles p ON a.id = p.account_id
WHERE a.id = ANY($1::VARCHAR[]) and a.is_suspended = FALSE
ORDER BY a.created_at DESC
`

type GetUserInfosRow struct {
	ID              string
	UserID          string
	UserName        string
	Bio             sql.NullString
	ProfileImageUrl sql.NullString
}

func (q *Queries) GetUserInfos(ctx context.Context, ids []string) ([]GetUserInfosRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserInfos, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserInfosRow
	for rows.Next() {
		var i GetUserInfosRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.UserName,
			&i.Bio,
			&i.ProfileImageUrl,
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

const isAdmin = `-- name: IsAdmin :one
SELECT is_admin FROM accounts
WHERE id = $1
`

func (q *Queries) IsAdmin(ctx context.Context, id string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isAdmin, id)
	var is_admin bool
	err := row.Scan(&is_admin)
	return is_admin, err
}

const isSuspended = `-- name: IsSuspended :one
SELECT is_suspended FROM accounts
WHERE id = $1
`

func (q *Queries) IsSuspended(ctx context.Context, id string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isSuspended, id)
	var is_suspended bool
	err := row.Scan(&is_suspended)
	return is_suspended, err
}

const searchAccountsByUserID = `-- name: SearchAccountsByUserID :many
SELECT id, user_id, user_name, is_suspended, is_admin, created_at FROM accounts
WHERE user_id LIKE CONCAT('%', $1, '%')
ORDER BY user_id
LIMIT $2 OFFSET $3
`

type SearchAccountsByUserIDParams struct {
	Concat interface{}
	Limit  int32
	Offset int32
}

func (q *Queries) SearchAccountsByUserID(ctx context.Context, arg SearchAccountsByUserIDParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, searchAccountsByUserID, arg.Concat, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.UserName,
			&i.IsSuspended,
			&i.IsAdmin,
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

const searchAccountsByUserName = `-- name: SearchAccountsByUserName :many
SELECT id, user_id, user_name, is_suspended, is_admin, created_at FROM accounts
WHERE user_name LIKE CONCAT('%', $1, '%')
ORDER BY user_name
LIMIT $2 OFFSET $3
`

type SearchAccountsByUserNameParams struct {
	Concat interface{}
	Limit  int32
	Offset int32
}

func (q *Queries) SearchAccountsByUserName(ctx context.Context, arg SearchAccountsByUserNameParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, searchAccountsByUserName, arg.Concat, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.UserName,
			&i.IsSuspended,
			&i.IsAdmin,
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

const suspendAccount = `-- name: SuspendAccount :execresult
UPDATE accounts
SET is_suspended = TRUE
WHERE id = $1
`

func (q *Queries) SuspendAccount(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, suspendAccount, id)
}

const unsuspendAccount = `-- name: UnsuspendAccount :execresult
UPDATE accounts
SET is_suspended = FALSE
WHERE id = $1
`

func (q *Queries) UnsuspendAccount(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, unsuspendAccount, id)
}

const updateAccountInfos = `-- name: UpdateAccountInfos :execresult
UPDATE accounts
SET user_id = COALESCE(NULLIF($1, ''), user_id),
    user_name = COALESCE(NULLIF($2, ''), user_name)
WHERE id = $3
`

type UpdateAccountInfosParams struct {
	UserID   interface{}
	UserName interface{}
	ID       string
}

func (q *Queries) UpdateAccountInfos(ctx context.Context, arg UpdateAccountInfosParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateAccountInfos, arg.UserID, arg.UserName, arg.ID)
}
