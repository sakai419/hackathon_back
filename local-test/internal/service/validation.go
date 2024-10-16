package service

import (
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func validateAccountParams(id, userID, userName string) error {
	if len(id) != 28 {
		return &apperrors.ErrInvalidInput{
			Message: "ID must be 28 characters",
		}
	}
	if len(userID) > 30 {
		return &apperrors.ErrInvalidInput{
			Message: "userID is too long",
		}
	}
	if len(userName) > 30 {
		return &apperrors.ErrInvalidInput{
			Message: "userName is too long",
		}
	}
	return nil
}

func validateReportParams(reporterAccountID, reportedAccountID, content string, reason model.ReportReason) error {
	if reporterAccountID == reportedAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "Reporter account ID and reported account ID must be different",
		}
	}

	switch reason {
	case model.ReportReasonSpam:
		return nil
	case model.ReportReasonHarrassment:
		return nil
	case model.ReportReasonInappropriateContent:
		return nil
	case model.ReportReasonOther:
		if content == "" {
			return &apperrors.ErrInvalidInput{
				Message: "Content must be provided for 'Other' reason",
			}
		}
		return nil
	default:
		return &apperrors.ErrInvalidInput{
			Message: "Invalid report reason",
		}
	}
}