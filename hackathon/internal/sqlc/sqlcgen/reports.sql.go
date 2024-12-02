// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: reports.sql

package sqlcgen

import (
	"context"
	"database/sql"
)

const createReport = `-- name: CreateReport :exec
INSERT INTO reports (reporter_account_id, reported_account_id, reason, content)
VALUES ($1, $2, $3, $4)
`

type CreateReportParams struct {
	ReporterAccountID string
	ReportedAccountID string
	Reason            ReportReason
	Content           sql.NullString
}

func (q *Queries) CreateReport(ctx context.Context, arg CreateReportParams) error {
	_, err := q.db.ExecContext(ctx, createReport,
		arg.ReporterAccountID,
		arg.ReportedAccountID,
		arg.Reason,
		arg.Content,
	)
	return err
}

const deleteReport = `-- name: DeleteReport :exec
DELETE FROM reports
WHERE id = $1
`

func (q *Queries) DeleteReport(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteReport, id)
	return err
}

const deleteReportsByReportedAccount = `-- name: DeleteReportsByReportedAccount :exec
DELETE FROM reports
WHERE reported_account_id = $1
`

func (q *Queries) DeleteReportsByReportedAccount(ctx context.Context, reportedAccountID string) error {
	_, err := q.db.ExecContext(ctx, deleteReportsByReportedAccount, reportedAccountID)
	return err
}

const getReportByID = `-- name: GetReportByID :one
SELECT id, reporter_account_id, reported_account_id, reason, content, created_at
FROM reports
WHERE id = $1
`

func (q *Queries) GetReportByID(ctx context.Context, id int64) (Report, error) {
	row := q.db.QueryRowContext(ctx, getReportByID, id)
	var i Report
	err := row.Scan(
		&i.ID,
		&i.ReporterAccountID,
		&i.ReportedAccountID,
		&i.Reason,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const getReportsByReportedAccount = `-- name: GetReportsByReportedAccount :many
SELECT id, reporter_account_id, reported_account_id, reason, content, created_at
FROM reports
WHERE reported_account_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetReportsByReportedAccount(ctx context.Context, reportedAccountID string) ([]Report, error) {
	rows, err := q.db.QueryContext(ctx, getReportsByReportedAccount, reportedAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Report
	for rows.Next() {
		var i Report
		if err := rows.Scan(
			&i.ID,
			&i.ReporterAccountID,
			&i.ReportedAccountID,
			&i.Reason,
			&i.Content,
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

const getUsersOrderByReportCount = `-- name: GetUsersOrderByReportCount :many
SELECT accounts.id, accounts.user_name, COUNT(reports.id) AS report_count
FROM accounts
LEFT JOIN reports ON accounts.id = reports.reported_account_id
GROUP BY accounts.id
ORDER BY report_count DESC
LIMIT $1 OFFSET $2
`

type GetUsersOrderByReportCountParams struct {
	Limit  int32
	Offset int32
}

type GetUsersOrderByReportCountRow struct {
	ID          string
	UserName    string
	ReportCount int64
}

func (q *Queries) GetUsersOrderByReportCount(ctx context.Context, arg GetUsersOrderByReportCountParams) ([]GetUsersOrderByReportCountRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsersOrderByReportCount, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersOrderByReportCountRow
	for rows.Next() {
		var i GetUsersOrderByReportCountRow
		if err := rows.Scan(&i.ID, &i.UserName, &i.ReportCount); err != nil {
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