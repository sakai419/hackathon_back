package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"

	"github.com/lib/pq"
)

func (r *Repository) FollowAndNotify(ctx context.Context, arg *model.FollowAndNotifyParams) error {
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
	createFollowParams := sqlcgen.CreateFollowParams{
		FollowerAccountID: arg.FollowerAccountID,
		FollowingAccountID: arg.FollowingAccountID,
	}
	if err := q.CreateFollow(ctx, createFollowParams); err != nil {
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
	createNotificationParams := sqlcgen.CreateNotificationParams{
		SenderAccountID: sql.NullString{String: arg.FollowerAccountID, Valid: true},
		RecipientAccountID: arg.FollowingAccountID,
		Type: sqlcgen.NotificationTypeFollow,
	}
	if err := q.CreateNotification(ctx, createNotificationParams); err != nil {
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

func (r *Repository) Unfollow(ctx context.Context, arg *model.UnfollowParams) error {
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

	// Delete follow
	deleteFollowParams := sqlcgen.DeleteFollowParams{
		FollowerAccountID: arg.FollowerAccountID,
		FollowingAccountID: arg.FollowingAccountID,
	}
	res, err := q.DeleteFollow(ctx, deleteFollowParams)
	if err != nil {
		tx.Rollback()
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
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "check if follow is deleted",
				Err: err,
			},
		)
	}
	if num == 0 {
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrRecordNotFound{
				Condition: "follow",
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

func (r *Repository) GetFollowerAccountIDs(ctx context.Context, arg *model.GetFollowerAccountIDsParams) ([]string, error) {
	// Get follower account ids
	query := sqlcgen.GetFollowerAccountIDsParams{
		FollowingAccountID: arg.FollowingAccountID,
		Limit: arg.Limit,
		Offset: arg.Offset,
	}
	followerAccountIDs, err := r.q.GetFollowerAccountIDs(ctx, query)
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

func (r *Repository) GetFollowingAccountIDs(ctx context.Context, arg *model.GetFollowingAccountIDsParams) ([]string, error) {
	// Get following account ids
	query := sqlcgen.GetFollowingAccountIDsParams{
		FollowerAccountID: arg.FollowerAccountID,
		Limit: arg.Limit,
		Offset: arg.Offset,
	}
	followingAccountIDs, err := r.q.GetFollowingAccountIDs(ctx, query)
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

func (r *Repository) RequestFollowAndNotify(ctx context.Context, arg *model.RequestFollowAndNotifyParams) error {
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
	createFollowRequestParams := sqlcgen.CreateFollowRequestParams{
		FollowerAccountID: arg.RequesterAccountID,
		FollowingAccountID: arg.RequestedAccountID,
	}
	if err := q.CreateFollowRequest(ctx, createFollowRequestParams); err != nil {
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
	createNotificationParams := sqlcgen.CreateNotificationParams{
		SenderAccountID: sql.NullString{String: arg.RequesterAccountID, Valid: true},
		RecipientAccountID: arg.RequestedAccountID,
		Type: sqlcgen.NotificationTypeFollowRequest,
	}
	if err := q.CreateNotification(ctx, createNotificationParams); err != nil {
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

func (r *Repository) AcceptFollowRequestAndNotify(ctx context.Context, arg *model.AcceptFollowRequestAndNotifyParams) error {
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
	acceptFollowRequestParams := sqlcgen.AcceptFollowRequestParams{
		FollowerAccountID: arg.RequesterAccountID,
		FollowingAccountID: arg.RequestedAccountID,
	}
	res, err := q.AcceptFollowRequest(ctx, acceptFollowRequestParams)
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
	createNotificationParams := sqlcgen.CreateNotificationParams{
		SenderAccountID: sql.NullString{String: arg.RequestedAccountID, Valid: true},
		RecipientAccountID: arg.RequesterAccountID,
		Type: sqlcgen.NotificationTypeRequestAccepted,
	}
	if err := q.CreateNotification(ctx, createNotificationParams); err != nil {
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

func (r *Repository) RejectFollowRequest(ctx context.Context, arg *model.RejectFollowRequestParams) error {
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

	// Delete follow request
	deleteFollowRequestParams := sqlcgen.DeleteFollowRequestParams{
		FollowerAccountID: arg.RequesterAccountID,
		FollowingAccountID: arg.RequestedAccountID,
	}
	res, err := q.DeleteFollowRequest(ctx, deleteFollowRequestParams)
	if err != nil {
		tx.Rollback()
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
		tx.Rollback()
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "check if follow request is deleted",
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