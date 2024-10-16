package model

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

// CreateReport
type CreateReportServiceParams struct {
	ReporterAccountID string
	ReportedUserID    string
	Reason            ReportReason
	Content           sql.NullString
}

type CreateReportRepositoryParams struct {
	ReporterAccountID string
	ReportedAccountID    string
	Reason            ReportReason
	Content           sql.NullString
}