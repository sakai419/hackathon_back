-- name: CreateReport :exec
INSERT INTO reports (reporter_account_id, reported_account_id, reason, content)
VALUES (?, ?, ?, ?);

-- name: GetReportByID :one
SELECT *
FROM reports
WHERE id = ?;

-- name: GetUsersOrderByReportCount :many
SELECT accounts.id, accounts.user_name, COUNT(reports.id) AS report_count
FROM accounts
LEFT JOIN reports ON accounts.id = reports.reported_account_id
GROUP BY accounts.id
ORDER BY report_count DESC
LIMIT ?
OFFSET ?;

-- name: GetReportsByReportedAccount :many
SELECT *
FROM reports
WHERE reported_account_id = ?
ORDER BY created_at DESC;

-- name: DeleteReport :exec
DELETE FROM reports
WHERE id = ?;

-- name: DeleteReportsByReportedAccount :exec
DELETE FROM reports
WHERE reported_account_id = ?;