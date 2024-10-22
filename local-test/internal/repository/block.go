package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"

	"github.com/lib/pq"
)

func (r *Repository) BlockUser(ctx context.Context, params *model.BlockUserParams) error {
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

	// Create block
	if err := q.CreateBlock(ctx, sqlcgen.CreateBlockParams{
		BlockerAccountID: params.BlockerAccountID,
		BlockedAccountID: params.BlockedAccountID,
	}); err != nil {
		tx.Rollback()
		if err.(*pq.Error).Code == ErrCodeDuplicateEntry {
			return apperrors.WrapRepositoryError(
				&apperrors.ErrDuplicateEntry{
					Entity: "blocker/blocked account id",
					Err: err,
				},
			)
		}

		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create block",
				Err: err,
			},
		)
	}

	// Delete follow
	_, err = q.DeleteFollow(ctx, sqlcgen.DeleteFollowParams{
		FollowerAccountID: params.BlockedAccountID,
		FollowingAccountID: params.BlockerAccountID,
	})
	if err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "delete follow",
				Err: err,
			},
		)
	}

	// Delete follow request
	_, err = q.DeleteFollowRequest(ctx, sqlcgen.DeleteFollowRequestParams{
		FollowerAccountID: params.BlockedAccountID,
		FollowingAccountID: params.BlockerAccountID,
	})
	if err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "delete follow request",
				Err: err,
			},
		)
	}

	// Delete notification
	senderAccountID := sql.NullString{String: params.BlockedAccountID, Valid: true}
	if err := q.DeleteAllNotificationsFromSender(ctx, sqlcgen.DeleteAllNotificationsFromSenderParams{
		SenderAccountID: senderAccountID,
		RecipientAccountID: params.BlockerAccountID,
	}); err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "delete notification",
				Err: err,
			},
		)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "commit transaction",
				Err: err,
			},
		)
	}

	return nil
}

func (r *Repository) UnblockUser(ctx context.Context, params *model.UnblockUserParams) error {
	// Delete block
	res, err := r.q.DeleteBlock(ctx, sqlcgen.DeleteBlockParams{
		BlockerAccountID: params.BlockerAccountID,
		BlockedAccountID: params.BlockedAccountID,
	})
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "delete block",
				Err: err,
			},
		)
	}

	// Check if block is deleted
	num, err := res.RowsAffected()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "check if block is deleted",
				Err: err,
			},
		)
	}
	if num == 0 {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "block",
			},
		)
	}

	return nil
}

func (r *Repository) GetBlockedAccountIDs(ctx context.Context, params *model.GetBlockedAccountIDsParams) ([]string, error) {
	// Get blocked account IDs
	blockedAccountIDs, err := r.q.GetBlockedAccountIDs(ctx, sqlcgen.GetBlockedAccountIDsParams{
		BlockerAccountID: params.BlockerAccountID,
		Limit:            params.Limit,
		Offset:           params.Offset,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get blocked account IDs",
				Err: err,
			},
		)
	}

	return blockedAccountIDs, nil
}

func (r *Repository) GetBlockCount(ctx context.Context, accountID string) (int64, error) {
	count, err := r.q.GetBlockCount(ctx, accountID)
	if err != nil {
		return 0, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get block count",
				Err: err,
			},
		)
	}

	return count, nil
}

func (r *Repository) IsBlocked(ctx context.Context, params *model.IsBlockedParams) (bool, error) {
	// Get block
	is_blocked, err := r.q.CheckBlockExists(ctx, sqlcgen.CheckBlockExistsParams{
		BlockerAccountID: params.BlockerAccountID,
		BlockedAccountID: params.BlockedAccountID,
	})
	if err != nil {
		return false, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "check is_blocked",
				Err: err,
			},
		)
	}

	return is_blocked, nil
}