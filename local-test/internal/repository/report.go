package repository

import (
	"context"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) CreateReport(ctx context.Context, arg *model.CreateReportParams) error {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "begin transaction",
				Err: err,
				},
			)
	}

	// Create query object with transaction
	q := r.q.WithTx(tx)

	// Create report
	createReportParams := sqlcgen.CreateReportParams{
		ReporterAccountID: arg.ReporterAccountID,
		ReportedAccountID: arg.ReportedAccountID,
		Reason:            sqlcgen.ReportReason(arg.Reason),
		Content:           arg.Content,
	}
	if err := q.CreateReport(ctx, createReportParams); err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create report",
				Err: err,
			},
		)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "commit transaction",
				Err: err,
			},
		)
	}

	return nil
}