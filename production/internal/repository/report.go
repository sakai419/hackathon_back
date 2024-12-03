package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) CreateReport(ctx context.Context, params *model.CreateReportParams) error {
	// Convert content to sql.NullString
	var content sql.NullString
	if params.Content != nil {
		content.String = *params.Content
		content.Valid = true
	} else {
		content.Valid = false
	}

	// Create report
	if err := r.q.CreateReport(ctx, sqlcgen.CreateReportParams{
		ReporterAccountID: params.ReporterAccountID,
		ReportedAccountID: params.ReportedAccountID,
		Reason:            sqlcgen.ReportReason(params.Reason),
		Content:           content,
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