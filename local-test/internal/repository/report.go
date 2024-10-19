package repository

import (
	"context"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) CreateReport(ctx context.Context, params *model.CreateReportParams) error {
	// Create report
	if err := r.q.CreateReport(ctx, sqlcgen.CreateReportParams{
		ReporterAccountID: params.ReporterAccountID,
		ReportedAccountID: params.ReportedAccountID,
		Reason:            sqlcgen.ReportReason(params.Reason),
		Content:           params.Content,
	}); err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create report",
				Err: err,
			},
		)
	}

	return nil
}