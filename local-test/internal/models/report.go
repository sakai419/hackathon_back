package models

import (
	"database/sql"
)

type ReportReason string

const (
	ReportReasonSpam ReportReason = "spam"
	ReportReasonHarrassment ReportReason = "harassment"
	ReportReasonInappropriateContent ReportReason = "inappropriate_content"
	ReportReasonOther ReportReason = "other"
)

type CreateReportByUserIDParams struct {
	ReporterAccountID string
	ReportedUserID    string
	Reason            ReportReason
	Content           sql.NullString
}

type CreateReportParams struct {
	ReporterAccountID string
	ReportedAccountID string
	Reason            ReportReason
	Content           sql.NullString
}