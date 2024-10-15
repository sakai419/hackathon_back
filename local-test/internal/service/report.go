package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"net/http"
)

func (s *Service) CreateReportByUserID(ctx context.Context, arg *model.CreateReportByUserIDParams) error {
	// Get account id by user id
	RepotedAccountID, err := s.repo.GetAccountIDByUserID(ctx, arg.ReportedUserID)
	if err != nil {
		if errors.Is(err, &apperrors.ErrRecordNotFound{}) {
			return &apperrors.AppError{
				Status:  http.StatusNotFound,
				Code:    "ACCOUNT_NOT_FOUND",
				Message: "Account not found",
				Err:     apperrors.WrapServiceError(
					&apperrors.ErrOperationFailed{
						Operation: "get account ID by user ID",
						Err: err,
					},
				),
			}
		}

		return apperrors.WrapServiceError(
			&apperrors.ErrOperationFailed{
				Operation: "get account ID by user ID",
				Err: err,
			},
		)
	}

	// Convert params
	createReportParams := &model.CreateReportParams{
		ReporterAccountID: arg.ReporterAccountID,
		ReportedAccountID: RepotedAccountID,
		Reason:            arg.Reason,
		Content:           arg.Content,
	}

	// Validate request
	if err := createReportParams.Validate(); err != nil {
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
	if err := s.repo.CreateReport(ctx, createReportParams); err != nil {
		return apperrors.WrapServiceError(
			&apperrors.ErrOperationFailed{
				Operation: "create report",
				Err: err,
			},
		)
	}

	return nil
}