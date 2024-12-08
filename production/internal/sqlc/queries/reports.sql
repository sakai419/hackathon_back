-- name: CreateReport :exec
INSERT INTO reports (reporter_account_id, reported_account_id, reason, content)
VALUES ($1, $2, $3, $4);

-- name: GetReportByID :one
SELECT *
FROM reports
WHERE id = $1;

-- name: GetReportedAccountIDsOrderByReportCount :many
SELECT accounts.id, COUNT(reports.id) AS report_count
FROM accounts
LEFT JOIN reports ON accounts.id = reports.reported_account_id
GROUP BY accounts.id
HAVING COUNT(reports.id) <> 0
ORDER BY report_count DESC
LIMIT $1 OFFSET $2;


-- name: GetReportsByReportedAccountID :many
SELECT *
FROM reports
WHERE reported_account_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: DeleteReport :exec
DELETE FROM reports
WHERE id = $1;

-- name: DeleteReportsByReportedAccount :exec
DELETE FROM reports
WHERE reported_account_id = $1;