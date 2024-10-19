package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"net/http"
)

func (s *Service) FollowAndNotify(ctx context.Context, params *model.FollowAndNotifyParams) error {
	// Validate params
	if err := params.Validate(); err != nil {
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
	if err := s.repo.FollowAndNotify(ctx, params); err != nil {
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

func (s *Service) Unfollow(ctx context.Context, params *model.UnfollowParams) error {
	// Validate params
	if err := params.Validate(); err != nil {
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
	if err := s.repo.Unfollow(ctx, params); err != nil {
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

func (s *Service) GetFollowerInfos(ctx context.Context, params *model.GetFollowerInfosParams) ([]*model.UserInfo, error) {
	// Validate params
	if err := params.Validate(); err != nil {
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

	// Get follower account ids
	followerAccountIDs, err := s.repo.GetFollowerAccountIDs(ctx, &model.GetFollowerAccountIDsParams{
		FollowingAccountID: params.FollowingAccountID,
		Limit:              params.Limit,
		Offset:             params.Offset,
	})
	if err != nil {
		return nil, apperrors.WrapServiceError(
			&apperrors.ErrOperationFailed{
				Operation: "get follower account ids",
				Err: err,
			},
		)
	}

	// Get user and profile info
	infos, err := s.repo.GetUserInfos(ctx, followerAccountIDs)
	if err != nil {
		return nil, apperrors.WrapServiceError(
			&apperrors.ErrOperationFailed{
				Operation: "get user and profile info by account ids",
				Err: err,
			},
		)
	}

	// Sort user infos
	followerInfos := sortUserInfos(infos, followerAccountIDs)

	return followerInfos, nil
}

func (s *Service) GetFollowingInfos(ctx context.Context, params *model.GetFollowingInfosParams) ([]*model.UserInfo, error) {
	// Validate params
	if err := params.Validate(); err != nil {
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

	// Get following account ids
	followingAccountIDs, err := s.repo.GetFollowingAccountIDs(ctx, &model.GetFollowingAccountIDsParams{
		FollowerAccountID: params.FollowerAccountID,
		Limit:             params.Limit,
		Offset:            params.Offset,
	})
	if err != nil {
		return nil, apperrors.WrapServiceError(
			&apperrors.ErrOperationFailed{
				Operation: "get following account ids",
				Err: err,
			},
		)
	}

	// Get user and profile info
	infos, err := s.repo.GetUserInfos(ctx, followingAccountIDs)
	if err != nil {
		return nil, apperrors.WrapServiceError(
			&apperrors.ErrOperationFailed{
				Operation: "get user and profile info by account ids",
				Err: err,
			},
		)
	}

	// Sort user infos
	followingInfos := sortUserInfos(infos, followingAccountIDs)

	return followingInfos, nil
}

func (s *Service) RequestFollowAndNotify(ctx context.Context, params *model.RequestFollowAndNotifyParams) error {
	// Validate params
	if err := params.Validate(); err != nil {
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
	if err := s.repo.RequestFollowAndNotify(ctx, params); err != nil {
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

func (s *Service) AcceptFollowRequestAndNotify(ctx context.Context, params *model.AcceptFollowRequestAndNotifyParams) error {
	// Validate params
	if err := params.Validate(); err != nil {
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
    if err := s.repo.AcceptFollowRequestAndNotify(ctx, params); err != nil {
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

func (s *Service) RejectFollowRequest(ctx context.Context, params *model.RejectFollowRequestParams) error {
	// Validate params
	if err := params.Validate(); err != nil {
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
	if err := s.repo.RejectFollowRequest(ctx, params); err != nil {
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

func sortUserInfos(userInfos []*model.UserInfoInternal, ids []string) []*model.UserInfo {
	userInfoMap := make(map[string]*model.UserInfoInternal)
	for _, userInfo := range userInfos {
		userInfoMap[userInfo.ID] = userInfo
	}

	sortedUserInfos := make([]*model.UserInfo, len(ids))
	for i, id := range ids {
		temp := userInfoMap[id]
		sortedUserInfos[i] = &model.UserInfo{
			UserID:          temp.UserID,
			UserName:        temp.UserName,
			Bio:             temp.Bio,
			ProfileImageURL: temp.ProfileImageURL,
		}
	}

	return sortedUserInfos
}