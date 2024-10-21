package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"

	"github.com/lib/pq"
)

func (r *Repository) FollowAndNotify(ctx context.Context, params *model.FollowAndNotifyParams) error {
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

	// Create follow
	if err := q.CreateFollow(ctx, sqlcgen.CreateFollowParams{
		FollowerAccountID: params.FollowerAccountID,
		FollowingAccountID: params.FollowingAccountID,
	}); err != nil {
		tx.Rollback()
		if err.(*pq.Error).Code == ErrCodeDuplicateEntry {
			return apperrors.WrapRepositoryError(
				&apperrors.ErrDuplicateEntry{
					Entity: "follower/following account id",
					Err: err,
				},
			)
		}

		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create follow",
				Err: err,
			},
		)
	}

	// Notify following user
	if err := q.CreateNotification(ctx, sqlcgen.CreateNotificationParams{
		SenderAccountID: sql.NullString{String: params.FollowerAccountID, Valid: true},
		RecipientAccountID: params.FollowingAccountID,
		Type: sqlcgen.NotificationTypeFollow,
	}); err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create notification",
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

func (r *Repository) Unfollow(ctx context.Context, params *model.UnfollowParams) error {
	// Delete follow
	res, err := r.q.DeleteFollow(ctx, sqlcgen.DeleteFollowParams{
		FollowerAccountID: params.FollowerAccountID,
		FollowingAccountID: params.FollowingAccountID,
	})
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "delete follow",
				Err: err,
			},
		)
	}

	// Check if follow is deleted
	num, err := res.RowsAffected()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "check if follow is deleted",
				Err: err,
			},
		)
	}
	if num == 0 {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "follow",
			},
		)
	}

	return nil
}

func (r *Repository) GetFollowerAccountIDs(ctx context.Context, params *model.GetFollowerAccountIDsParams) ([]string, error) {
	// Get follower account ids
	followerAccountIDs, err := r.q.GetFollowerAccountIDs(ctx, sqlcgen.GetFollowerAccountIDsParams{
		FollowingAccountID: params.FollowingAccountID,
		Limit: params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get follower account ids",
				Err: err,
			},
		)
	}

	return followerAccountIDs, nil
}

func (r *Repository) GetFollowingAccountIDs(ctx context.Context, params *model.GetFollowingAccountIDsParams) ([]string, error) {
	// Get following account ids
	followingAccountIDs, err := r.q.GetFollowingAccountIDs(ctx, sqlcgen.GetFollowingAccountIDsParams{
		FollowerAccountID: params.FollowerAccountID,
		Limit: params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get following account ids",
				Err: err,
			},
		)
	}

	return followingAccountIDs, nil
}

func (r *Repository) RequestFollowAndNotify(ctx context.Context, params *model.RequestFollowAndNotifyParams) error {
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

	// Create follow request
	if err := q.CreateFollowRequest(ctx, sqlcgen.CreateFollowRequestParams{
		FollowerAccountID: params.RequesterAccountID,
		FollowingAccountID: params.RequestedAccountID,
	}); err != nil {
		tx.Rollback()
		if err.(*pq.Error).Code == ErrCodeDuplicateEntry {
			return apperrors.WrapRepositoryError(
				&apperrors.ErrDuplicateEntry{
					Entity: "requester/requested account id",
					Err: err,
				},
			)
		}

		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create follow request",
				Err: err,
			},
		)
	}

	// Notify requested user
	if err := q.CreateNotification(ctx, sqlcgen.CreateNotificationParams{
		SenderAccountID: sql.NullString{String: params.RequesterAccountID, Valid: true},
		RecipientAccountID: params.RequestedAccountID,
		Type: sqlcgen.NotificationTypeFollowRequest,
	}); err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create notification",
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

func (r *Repository) AcceptFollowRequestAndNotify(ctx context.Context, params *model.AcceptFollowRequestAndNotifyParams) error {
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

	// Accept follow request
	res, err := q.AcceptFollowRequest(ctx, sqlcgen.AcceptFollowRequestParams{
		FollowerAccountID: params.RequesterAccountID,
		FollowingAccountID: params.RequestedAccountID,
	})
	if err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "accept follow request",
				Err: err,
			},
		)
	}

	// Check if follow request is accepted
	num, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "check if follow request is accepted",
				Err: err,
			},
		)
	}
	if num == 0 {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "follow request",
			},
		)
	}

	// Notify requester
	if err := q.CreateNotification(ctx, sqlcgen.CreateNotificationParams{
		SenderAccountID: sql.NullString{String: params.RequestedAccountID, Valid: true},
		RecipientAccountID: params.RequesterAccountID,
		Type: sqlcgen.NotificationTypeRequestAccepted,
	}); err != nil {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "create notification",
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

func (r *Repository) RejectFollowRequest(ctx context.Context, params *model.RejectFollowRequestParams) error {
	// Delete follow request
	res, err := r.q.DeleteFollowRequest(ctx, sqlcgen.DeleteFollowRequestParams{
		FollowerAccountID: params.RequesterAccountID,
		FollowingAccountID: params.RequestedAccountID,
	})
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "delete follow request",
				Err: err,
			},
		)
	}

	// Check if follow request is deleted
	num, err := res.RowsAffected()
	if err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "check if follow request is deleted",
				Err: err,
			},
		)
	}
	if num == 0 {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "follow request",
			},
		)
	}

	return nil
}