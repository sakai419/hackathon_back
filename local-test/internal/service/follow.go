package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) FollowAndNotify(ctx context.Context, params *model.FollowAndNotifyParams) error {
	// Validate params
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

	// Create follow
	if err := s.repo.FollowAndNotify(ctx, params); err != nil {
		return apperrors.NewDuplicateEntryAppError("Follow", "follow", err)
	}

	return nil
}

func (s *Service) Unfollow(ctx context.Context, params *model.UnfollowParams) error {
	// Validate params
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

	// Unfollow
	if err := s.repo.Unfollow(ctx, params); err != nil {
		return apperrors.NewNotFoundAppError("Follow", "unfollow", err)
	}

	return nil
}

func (s *Service) GetFollowerInfos(ctx context.Context, params *model.GetFollowerInfosParams) ([]*model.UserInfo, error) {
	// Validate params
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get follower account ids
	followerAccountIDs, err := s.repo.GetFollowerAccountIDs(ctx, &model.GetFollowerAccountIDsParams{
		FollowingAccountID: params.FollowingAccountID,
		Limit:              params.Limit,
		Offset:             params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get follower account ids", err)
	}

	// Get user and profile info
	infos, err := s.repo.GetUserInfos(ctx, followerAccountIDs)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("follower info", "get follower infos", err)
	}

	// Sort user infos
	followerInfos := sortUserInfos(infos, followerAccountIDs)

	return followerInfos, nil
}

func (s *Service) GetFollowingInfos(ctx context.Context, params *model.GetFollowingInfosParams) ([]*model.UserInfo, error) {
	// Validate params
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get following account ids
	followingAccountIDs, err := s.repo.GetFollowingAccountIDs(ctx, &model.GetFollowingAccountIDsParams{
		FollowerAccountID: params.FollowerAccountID,
		Limit:             params.Limit,
		Offset:            params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get following account ids", err)
	}

	// Get user and profile info
	infos, err := s.repo.GetUserInfos(ctx, followingAccountIDs)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("following user info", "get following user infos", err)
	}

	// Sort user infos
	followingInfos := sortUserInfos(infos, followingAccountIDs)

	return followingInfos, nil
}

func (s *Service) RequestFollowAndNotify(ctx context.Context, params *model.RequestFollowAndNotifyParams) error {
	// Validate params
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

	// Request follow
	if err := s.repo.RequestFollowAndNotify(ctx, params); err != nil {
		return apperrors.NewDuplicateEntryAppError("Follow request", "request follow", err)
	}

	return nil
}

func (s *Service) AcceptFollowRequestAndNotify(ctx context.Context, params *model.AcceptFollowRequestAndNotifyParams) error {
	// Validate params
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

    // Accept follow request
    if err := s.repo.AcceptFollowRequestAndNotify(ctx, params); err != nil {
		return apperrors.NewNotFoundAppError("Follow request", "accept follow request", err)
    }

    return nil
}

func (s *Service) RejectFollowRequest(ctx context.Context, params *model.RejectFollowRequestParams) error {
	// Validate params
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

	// Reject follow request
	if err := s.repo.RejectFollowRequest(ctx, params); err != nil {
		return apperrors.NewNotFoundAppError("Follow request", "reject follow request", err)
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