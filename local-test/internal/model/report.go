package model

import (
	"database/sql"
	"local-test/pkg/apperrors"
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

func (p *CreateReportParams) Validate() error {
	if p.ReporterAccountID == p.ReportedAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "Reporter account ID and reported account ID must be different",
		}
	}

	switch p.Reason {
	case ReportReasonSpam:
		return nil
	case ReportReasonHarrassment:
		return nil
	case ReportReasonInappropriateContent:
		return nil
	case ReportReasonOther:
		if !p.Content.Valid {
			return &apperrors.ErrInvalidInput{
				Message: "Content is required for other reason",
			}
		}
		return nil
	default:
		return &apperrors.ErrInvalidInput{
			Message: "Invalid report reason",
		}
	}
}