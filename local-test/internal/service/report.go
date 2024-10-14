package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/utils"
	"net/http"
)

func (s *Service) CreateReportByUserID(ctx context.Context, arg *model.CreateReportByUserIDParams) error {
	// Get account id by user id
	RepotedAccountID, err := s.repo.GetAccountIDByUserId(ctx, arg.ReportedUserID)
	if err != nil {
		if errors.Is(err, &utils.ErrRecordNotFound{}) {
			return &utils.AppError{
				Status:  http.StatusNotFound,
				Code:    "ACCOUNT_NOT_FOUND",
				Message: "Account not found",
				Err:     utils.WrapServiceError(
					&utils.ErrOperationFailed{
						Operation: "get account ID by user ID",
						Err: err,
					},
				),
			}
		}
		return utils.WrapServiceError(
			&utils.ErrOperationFailed{
				Operation: "get account ID by user ID",
				Err: err,
			},
		)
	}

	// Convert params
	params := &model.CreateReportParams{
		ReporterAccountID: arg.ReporterAccountID,
		ReportedAccountID: RepotedAccountID,
		Reason:            arg.Reason,
		Content:           arg.Content,
	}

	// Validate request
	if err := params.Validate(); err != nil {
		return &utils.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     utils.WrapServiceError(
				&utils.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		}
	}

	// Create report
	if err := s.repo.CreateReport(ctx, params); err != nil {
		return utils.WrapServiceError(
			&utils.ErrOperationFailed{
				Operation: "create report",
				Err: err,
			},
		)
	}

	return nil
}