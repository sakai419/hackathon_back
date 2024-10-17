// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accounts.sql

package sqlcgen

import (
	"context"
	"database/sql"
	"time"

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

const getAccountCreationDate = `-- name: GetAccountCreationDate :one
SELECT created_at FROM accounts
WHERE id = $1
`

func (q *Queries) GetAccountCreationDate(ctx context.Context, id string) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, getAccountCreationDate, id)
	var created_at time.Time
	err := row.Scan(&created_at)
	return created_at, err
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

const getUserAndProfileInfos = `-- name: GetUserAndProfileInfos :many
SELECT a.user_id, a.user_name, p.bio, p.profile_image_url
FROM accounts a
JOIN profiles p ON a.id = p.account_id
WHERE a.id = ANY($3::VARCHAR[])
ORDER BY a.created_at DESC
LIMIT $1 OFFSET $2
`

type GetUserAndProfileInfosParams struct {
	Limit  int32
	Offset int32
	Ids    []string
}

type GetUserAndProfileInfosRow struct {
	UserID          string
	UserName        string
	Bio             sql.NullString
	ProfileImageUrl sql.NullString
}

func (q *Queries) GetUserAndProfileInfos(ctx context.Context, arg GetUserAndProfileInfosParams) ([]GetUserAndProfileInfosRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserAndProfileInfos, arg.Limit, arg.Offset, pq.Array(arg.Ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserAndProfileInfosRow
	for rows.Next() {
		var i GetUserAndProfileInfosRow
		if err := rows.Scan(
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
SELECT id, user_id, user_name, is_suspended, is_admin, created_at, updated_at FROM accounts
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

const searchAccountsByUserName = `-- name: SearchAccountsByUserName :many
SELECT id, user_id, user_name, is_suspended, is_admin, created_at, updated_at FROM accounts
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

const updateAccountUserID = `-- name: UpdateAccountUserID :exec
UPDATE accounts
SET user_id = $1
WHERE id = $2
`

type UpdateAccountUserIDParams struct {
	UserID string
	ID     string
}

func (q *Queries) UpdateAccountUserID(ctx context.Context, arg UpdateAccountUserIDParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountUserID, arg.UserID, arg.ID)
	return err
}

const updateAccountUserName = `-- name: UpdateAccountUserName :exec
UPDATE accounts
SET user_name = $1
WHERE id = $2
`

type UpdateAccountUserNameParams struct {
	UserName string
	ID       string
}

func (q *Queries) UpdateAccountUserName(ctx context.Context, arg UpdateAccountUserNameParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountUserName, arg.UserName, arg.ID)
	return err
}
