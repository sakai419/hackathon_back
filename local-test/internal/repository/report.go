package repository

import (
	"context"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) CreateReport(ctx context.Context, arg *model.CreateReportParams) error {
	// Create report
	params := sqlcgen.CreateReportParams{
		ReporterAccountID: arg.ReporterAccountID,
		ReportedAccountID: arg.ReportedAccountID,
		Reason:            sqlcgen.ReportReason(arg.Reason),
		Content:           arg.Content,
	}
	if err := r.q.CreateReport(ctx, params); err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create report",
				Err: err,
			},
		)
	}

	return nil
}