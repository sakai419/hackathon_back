package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"net/http"
)

func (s *Service) CreateReportByUserID(ctx context.Context, arg *model.CreateReportServiceParams) error {
	// Get reported account ID
	reportedAccountID, err := s.repo.GetAccountIDByUserID(ctx, arg.ReportedUserID)
	if err != nil {
		return &apperrors.AppError{
			Status:  http.StatusNotFound,
			Code:    "NOT_FOUND",
			Message: "Account not found",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID by user ID",
					Err: err,
				},
			),
		}
	}

	// Validate request
	params := &model.CreateReportRepositoryParams{
		ReporterAccountID: arg.ReporterAccountID,
		ReportedAccountID: reportedAccountID,
		Reason:            arg.Reason,
		Content: 		   arg.Content,
	}
	if err := validateReportParams(params.ReporterAccountID, params.ReportedAccountID, params.Content.String, params.Reason); err != nil {
		return &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		}
	}

	// Create report
	if err := s.repo.CreateReport(ctx, params); err != nil {
		return apperrors.WrapServiceError(
			&apperrors.ErrOperationFailed{
				Operation: "create report",
				Err: err,
			},
		)
	}

	return nil
}