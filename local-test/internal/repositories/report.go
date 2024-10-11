package repositories

import (
	"context"
	"local-test/internal/models"
	sqlcgen "local-test/internal/sqlc/generated"
	"local-test/pkg/utils"
)

func (r *Repository) CreateReport(ctx context.Context, arg *models.CreateReportParams) error {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "begin transaction", Err: err})
	}

	// Create query object with transaction
	q := r.q.WithTx(tx)

	// Create report
	params := sqlcgen.CreateReportParams{
		ReporterAccountID: arg.ReporterAccountID,
		ReportedAccountID: arg.ReportedAccountID,
		Reason:            sqlcgen.ReportsReason(arg.Reason),
		Content:           arg.Content,
	}
	if err := q.CreateReport(ctx, params); err != nil {
		tx.Rollback()
		return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "create report", Err: err})
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return utils.WrapRepositoryError(&utils.ErrOperationFailed{Operation: "commit transaction", Err: err})
	}

	return nil
}