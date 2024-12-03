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

const filterAccessibleAccountIDs = `-- name: FilterAccessibleAccountIDs :many
SELECT a.id::VARCHAR as account_id
FROM unnest($1::VARCHAR[]) AS a(id)
LEFT JOIN blocks AS b
    ON (b.blocked_account_id = a.id
    AND b.blocker_account_id = $2)
    OR (b.blocked_account_id = $2
    AND b.blocker_account_id = a.id)
LEFT JOIN settings AS s
    ON s.account_id = a.id
LEFT JOIN follows AS f
    ON f.following_account_id = a.id
    AND f.follower_account_id = $2
WHERE
    b.blocked_account_id IS NULL
    AND (s.is_private = FALSE OR f.follower_account_id IS NOT NULL)
`

type FilterAccessibleAccountIDsParams struct {
	AccountIds      []string
	ClientAccountID string
}

func (q *Queries) FilterAccessibleAccountIDs(ctx context.Context, arg FilterAccessibleAccountIDsParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, filterAccessibleAccountIDs, pq.Array(arg.AccountIds), arg.ClientAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var account_id string
		if err := rows.Scan(&account_id); err != nil {
			return nil, err
		}
		items = append(items, account_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
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

const getUserInfo = `-- name: GetUserInfo :one
SELECT a.id, a.user_id, a.user_name, a.is_admin, a.created_at, p.bio, p.profile_image_url, p.banner_image_url, s.is_private
FROM accounts a
JOIN profiles p ON a.id = p.account_id
JOIN settings s ON a.id = s.account_id
WHERE a.id = $1
`

type GetUserInfoRow struct {
	ID              string
	UserID          string
	UserName        string
	IsAdmin         bool
	CreatedAt       time.Time
	Bio             sql.NullString
	ProfileImageUrl sql.NullString
	BannerImageUrl  sql.NullString
	IsPrivate       sql.NullBool
}

func (q *Queries) GetUserInfo(ctx context.Context, id string) (GetUserInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getUserInfo, id)
	var i GetUserInfoRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.UserName,
		&i.IsAdmin,
		&i.CreatedAt,
		&i.Bio,
		&i.ProfileImageUrl,
		&i.BannerImageUrl,
		&i.IsPrivate,
	)
	return i, err
}

const getUserInfos = `-- name: GetUserInfos :many
SELECT a.id, a.user_id, a.user_name, a.is_admin, a.created_at, p.bio, p.profile_image_url, p.banner_image_url, s.is_private
FROM accounts a
JOIN profiles p ON a.id = p.account_id
JOIN settings s ON a.id = s.account_id
WHERE a.id = ANY($1::VARCHAR[]) and a.is_suspended = FALSE
ORDER BY a.created_at DESC
`

type GetUserInfosRow struct {
	ID              string
	UserID          string
	UserName        string
	IsAdmin         bool
	CreatedAt       time.Time
	Bio             sql.NullString
	ProfileImageUrl sql.NullString
	BannerImageUrl  sql.NullString
	IsPrivate       sql.NullBool
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
			&i.IsAdmin,
			&i.CreatedAt,
			&i.Bio,
			&i.ProfileImageUrl,
			&i.BannerImageUrl,
			&i.IsPrivate,
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
WHERE user_id ILIKE CONCAT('%', $1, '%')
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
WHERE user_name ILIKE CONCAT('%', $1, '%')
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

const searchAccountsOrderByCreatedAt = `-- name: SearchAccountsOrderByCreatedAt :many
SELECT a.id, a.is_admin, a.user_id, a.user_name, a.created_at, p.bio, p.profile_image_url, s.is_private FROM accounts AS a
JOIN profiles AS p ON a.id = p.account_id
JOIN settings AS s ON a.id = s.account_id
WHERE
    a.user_id ILIKE CONCAT('%', $3::VARCHAR, '%')
    OR a.user_name ILIKE CONCAT('%', $3::VARCHAR, '%')
    OR p.bio ILIKE CONCAT('%', $3::VARCHAR, '%')
    AND a.is_suspended = FALSE
ORDER BY a.created_at DESC
LIMIT $1 OFFSET $2
`

type SearchAccountsOrderByCreatedAtParams struct {
	Limit   int32
	Offset  int32
	Keyword string
}

type SearchAccountsOrderByCreatedAtRow struct {
	ID              string
	IsAdmin         bool
	UserID          string
	UserName        string
	CreatedAt       time.Time
	Bio             sql.NullString
	ProfileImageUrl sql.NullString
	IsPrivate       sql.NullBool
}

func (q *Queries) SearchAccountsOrderByCreatedAt(ctx context.Context, arg SearchAccountsOrderByCreatedAtParams) ([]SearchAccountsOrderByCreatedAtRow, error) {
	rows, err := q.db.QueryContext(ctx, searchAccountsOrderByCreatedAt, arg.Limit, arg.Offset, arg.Keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchAccountsOrderByCreatedAtRow
	for rows.Next() {
		var i SearchAccountsOrderByCreatedAtRow
		if err := rows.Scan(
			&i.ID,
			&i.IsAdmin,
			&i.UserID,
			&i.UserName,
			&i.CreatedAt,
			&i.Bio,
			&i.ProfileImageUrl,
			&i.IsPrivate,
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
