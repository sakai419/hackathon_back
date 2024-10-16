package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"net/http"
)

func (s *Service) FollowAndNotify(ctx context.Context, arg *model.FollowAndNotifyParams) error {
	// Validate params
	if err := arg.Validate(); err != nil {
		return &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		}
	}

	// Create follow
	if err := s.repo.FollowAndNotify(ctx, arg); err != nil {
		// Check if the error is a duplicate entry error
		var duplicateErr *apperrors.ErrDuplicateEntry
		if errors.As(err, &duplicateErr) {
			return &apperrors.AppError{
				Status:  http.StatusConflict,
				Code:    "DUPLICATE_ENTRY",
				Message: "Follow already exists",
				Err:     apperrors.WrapServiceError(
					&apperrors.ErrOperationFailed{
						Operation: "create follow",
						Err: duplicateErr,
					},
				),
			}
		}

		return &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to create follow",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "create follow",
					Err: err,
				},
			),
		}
	}

	return nil
}

func (s *Service) Unfollow(ctx context.Context, arg *model.UnfollowParams) error {
	// Validate params
	if err := arg.Validate(); err != nil {
		return &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		}
	}

	// Unfollow
	if err := s.repo.Unfollow(ctx, arg); err != nil {
		// Check if the error is a record not found error
		var notFoundErr *apperrors.ErrRecordNotFound
		if errors.As(err, &notFoundErr) {
			return &apperrors.AppError{
				Status:  http.StatusNotFound,
				Code:    "FOLLOW_NOT_FOUND",
				Message: "Follow not found",
				Err:     apperrors.WrapServiceError(
					&apperrors.ErrOperationFailed{
						Operation: "unfollow",
						Err: notFoundErr,
					},
				),
			}
		}

		return &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to unfollow",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "unfollow",
					Err: err,
				},
			),
		}
	}

	return nil
}

func (s *Service) GetFollowerInfos(ctx context.Context, arg *model.GetFollowerInfosParams) ([]*model.UserAndProfileInfo, error) {
	// Validate params
	if err := arg.Validate(); err != nil {
		return nil, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		}
	}

	// Convert params
	getFollowerAccountIDsParams := &model.GetFollowerAccountIDsParams{
		FollowingAccountID: arg.FollowingAccountID,
		Limit:              arg.Limit,
		Offset:             arg.Offset,
	}

	// Get follower account ids
	followerAccountIDs, err := s.repo.GetFollowerAccountIDs(ctx, getFollowerAccountIDsParams)
	if err != nil {
		return nil, apperrors.WrapServiceError(
			&apperrors.ErrOperationFailed{
				Operation: "get follower account ids",
				Err: err,
			},
		)
	}

	// Convert params
	getUserAndProfileInfoByAccountIDs := &model.GetUserAndProfileInfoByAccountIDsParams{
		Limit:  arg.Limit,
		Offset: arg.Offset,
		IDs:    followerAccountIDs,
	}

	// Get user and profile info
	followerInfos, err := s.repo.GetUserAndProfileInfoByAccountIDs(ctx, getUserAndProfileInfoByAccountIDs)
	if err != nil {
		return nil, apperrors.WrapServiceError(
			&apperrors.ErrOperationFailed{
				Operation: "get user and profile info by account ids",
				Err: err,
			},
		)
	}

	return followerInfos, nil
}

func (s *Service) GetFollowingInfos(ctx context.Context, arg *model.GetFollowingInfosParams) ([]*model.UserAndProfileInfo, error) {
	// Validate params
	if err := arg.Validate(); err != nil {
		return nil, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		}
	}

	// Convert params
	getFollowingAccountIDsParams := &model.GetFollowingAccountIDsParams{
		FollowerAccountID: arg.FollowerAccountID,
		Limit:             arg.Limit,
		Offset:            arg.Offset,
	}

	// Get following account ids
	followingAccountIDs, err := s.repo.GetFollowingAccountIDs(ctx, getFollowingAccountIDsParams)
	if err != nil {
		return nil, apperrors.WrapServiceError(
			&apperrors.ErrOperationFailed{
				Operation: "get following account ids",
				Err: err,
			},
		)
	}

	// Convert params
	getUserAndProfileInfoByAccountIDs := &model.GetUserAndProfileInfoByAccountIDsParams{
		Limit:  arg.Limit,
		Offset: arg.Offset,
		IDs:    followingAccountIDs,
	}

	// Get user and profile info
	followingInfos, err := s.repo.GetUserAndProfileInfoByAccountIDs(ctx, getUserAndProfileInfoByAccountIDs)
	if err != nil {
		return nil, apperrors.WrapServiceError(
			&apperrors.ErrOperationFailed{
				Operation: "get user and profile info by account ids",
				Err: err,
			},
		)
	}

	return followingInfos, nil
}

func (s *Service) RequestFollowAndNotify(ctx context.Context, arg *model.RequestFollowAndNotifyParams) error {
	// Validate params
	if err := arg.Validate(); err != nil {
		return &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		}
	}

	// Request follow
	if err := s.repo.RequestFollowAndNotify(ctx, arg); err != nil {
		// Check if the error is a duplicate entry error
		var duplicateErr *apperrors.ErrDuplicateEntry
		if errors.As(err, &duplicateErr) {
			return &apperrors.AppError{
				Status:  http.StatusConflict,
				Code:    "DUPLICATE_ENTRY",
				Message: "Follow request already exists",
				Err:     apperrors.WrapServiceError(
					&apperrors.ErrOperationFailed{
						Operation: "request follow",
						Err: duplicateErr,
					},
				),
			}
		}

		return &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to request follow",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "request follow",
					Err: err,
				},
			),
		}
	}

	return nil
}

func (s *Service) AcceptFollowRequestAndNotify(ctx context.Context, arg *model.AcceptFollowRequestAndNotifyParams) error {
	// Validate params
	if err := arg.Validate(); err != nil {
		return &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		}
	}

    // Accept follow request
    if err := s.repo.AcceptFollowRequestAndNotify(ctx, arg); err != nil {
        // Check if the error is a record not found error
        var notFoundErr *apperrors.ErrRecordNotFound
        if errors.As(err, &notFoundErr) {
            return &apperrors.AppError{
                Status:  http.StatusNotFound,
                Code:    "FOLLOW_REQUEST_NOT_FOUND",
                Message: "Follow request not found",
                Err:     apperrors.WrapServiceError(
                    &apperrors.ErrOperationFailed{
                        Operation: "accept follow request",
                        Err: notFoundErr,
                    },
                ),
            }
        }

        return &apperrors.AppError{
            Status:  http.StatusInternalServerError,
            Code:    "DATABASE_ERROR",
            Message: "Failed to accept follow request",
            Err:     apperrors.WrapServiceError(
                &apperrors.ErrOperationFailed{
                    Operation: "accept follow request",
                    Err: err,
                },
            ),
        }
    }

    return nil
}

func (s *Service) RejectFollowRequest(ctx context.Context, arg *model.RejectFollowRequestParams) error {
	// Validate params
	if err := arg.Validate(); err != nil {
		return &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		}
	}

	// Reject follow request
	if err := s.repo.RejectFollowRequest(ctx, arg); err != nil {
		// Check if the error is a record not found error
		var notFoundErr *apperrors.ErrRecordNotFound
		if errors.As(err, &notFoundErr) {
			return &apperrors.AppError{
				Status:  http.StatusNotFound,
				Code:    "FOLLOW_REQUEST_NOT_FOUND",
				Message: "Follow request not found",
				Err:     apperrors.WrapServiceError(
					&apperrors.ErrOperationFailed{
						Operation: "reject follow request",
						Err: notFoundErr,
					},
				),
			}
		}

		return &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to reject follow request",
			Err:     apperrors.WrapServiceError(
				&apperrors.ErrOperationFailed{
					Operation: "reject follow request",
					Err: err,
				},
			),
		}
	}

	return nil
}