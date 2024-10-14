package repository

import (
	"context"
	"database/sql"
	"local-test/internal/model"
	sqlcgen "local-test/internal/sqlc/generated"
	"local-test/pkg/utils"

	"github.com/go-sql-driver/mysql"
)

func (r *Repository) FollowAndNotify(ctx context.Context, arg *model.FollowAndNotifyParams) error {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return utils.WrapRepositoryError(
			&utils.ErrOperationFailed{
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
		if err.(*mysql.MySQLError).Number == 1062 {
			return utils.WrapRepositoryError(
				&utils.ErrDuplicateEntry{
					Entity: "follower/following account id",
					Err: err,
				},
			)
		}

		return utils.WrapRepositoryError(
			&utils.ErrOperationFailed{
				Operation: "create follow",
				Err: err,
			},
		)
	}

	// Create notification
	CreateNotificationParams := sqlcgen.CreateNotificationParams{
		SenderAccountID: sql.NullString{String: arg.FollowerAccountID, Valid: true},
		RecipientAccountID: arg.FollowingAccountID,
		Type: sqlcgen.NotificationsTypeFollow,
		Content: sql.NullString{String: "You have a new follower", Valid: true},
	}
	if err := q.CreateNotification(ctx, CreateNotificationParams); err != nil {
		tx.Rollback()
		return utils.WrapRepositoryError(
			&utils.ErrOperationFailed{
				Operation: "create notification",
				Err: err,
			},
		)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return utils.WrapRepositoryError(
			&utils.ErrOperationFailed{
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
		return utils.WrapRepositoryError(
			&utils.ErrOperationFailed{
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
		return utils.WrapRepositoryError(
			&utils.ErrOperationFailed{
				Operation: "delete follow",
				Err: err,
			},
		)
	}

	// Check if follow is deleted
	num, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return utils.WrapRepositoryError(
			&utils.ErrOperationFailed{
				Operation: "check if follow is deleted",
				Err: err,
			},
		)
	}
	if num == 0 {
		tx.Rollback()
		return utils.WrapRepositoryError(
			&utils.ErrRecordNotFound{
				Condition: "follow",
			},
		)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return utils.WrapRepositoryError(
			&utils.ErrOperationFailed{
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
		return nil, utils.WrapRepositoryError(
			&utils.ErrOperationFailed{
				Operation: "get follower account ids",
				Err: err,
			},
		)
	}

	return followerAccountIDs, nil
}