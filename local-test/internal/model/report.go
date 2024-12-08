package model

import (
	"local-test/pkg/apperrors"
)

type ReportReason string

const (
	ReportReasonSpam ReportReason = "spam"
	ReportReasonHarrassment ReportReason = "harassment"
	ReportReasonInappropriateContent ReportReason = "inappropriate_content"
	ReportReasonOther ReportReason = "other"
)

// CreateReport
type CreateReportParams struct {
	ReporterAccountID string
	ReportedAccountID string
	Reason            ReportReason
	Content           *string
}

func (p *CreateReportParams) Validate() error {
	if p.ReporterAccountID == p.ReportedAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "reporter and reported account ID must be different",
		}
	}

	switch p.Reason {
	case ReportReasonSpam, ReportReasonHarrassment, ReportReasonInappropriateContent:
		return nil
	case ReportReasonOther:
		if p.Content == nil {
			return &apperrors.ErrInvalidInput{
				Message: "content is required for 'other' report reason",
			}
		}
	default:
		return &apperrors.ErrInvalidInput{
			Message: "invalid report reason",
		}
	}

	return nil
}

type GetReportedUserInfosOrderByReportCountParams struct {
	Limit  int32
	Offset int32
}

func (p *GetReportedUserInfosOrderByReportCountParams) Validate() error {
	if p.Limit <= 0 {
		return &apperrors.ErrInvalidInput{
			Message: "limit must be greater than 0",
		}
	}

	if p.Offset < 0 {
		return &apperrors.ErrInvalidInput{
			Message: "offset must be greater than or equal to 0",
		}
	}

	return nil
}

type GetReportedAccountIDsOrderByReportCountParams struct {
	Limit  int32
	Offset int32
}

type ReportedUserInfoInternal struct {
	AccountID   string
	ReportCount int64
}

type ReportedUserInfo struct {
	UserInfo UserInfo
	ReportCount int64
}