package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) CreateAccount(ctx context.Context, params *model.CreateAccountParams) error {
	// Validate input
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

    if err := s.repo.CreateAccount(ctx, params); err != nil {
		return apperrors.NewDuplicateEntryAppError("account", "create account", err)
    }

    return nil
}

func (s *Service) DeleteMyAccount(ctx context.Context, accountID string) error {
	// Delete account
	if err := s.repo.DeleteMyAccount(ctx, accountID); err != nil {
		apperrors.NewNotFoundAppError("account", "delete account", err)
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
		sortedUserInfos[i] = convertToUserInfo(temp)
	}

	return sortedUserInfos
}

func convertToUserInfo(userInfo *model.UserInfoInternal) *model.UserInfo {
	return &model.UserInfo{
		UserID:          userInfo.UserID,
		UserName:        userInfo.UserName,
		Bio:             userInfo.Bio,
		ProfileImageURL: userInfo.ProfileImageURL,
		IsPrivate:       userInfo.IsPrivate,
		IsAdmin:         userInfo.IsAdmin,
		IsFollowing:     userInfo.IsFollowing,
		IsFollowed:      userInfo.IsFollowed,
	}
}

func convertToUserInfos(userInfos []*model.UserInfoInternal) []*model.UserInfo {
	userInfo := make([]*model.UserInfo, len(userInfos))
	for i, u := range userInfos {
		userInfo[i] = convertToUserInfo(u)
	}

	return userInfo
}

func convertToUserInfoWithoutBio(userInfo *model.UserInfoInternal) *model.UserInfoWithoutBio {
	return &model.UserInfoWithoutBio{
		UserID:          userInfo.UserID,
		UserName:        userInfo.UserName,
		ProfileImageURL: userInfo.ProfileImageURL,
		IsPrivate:       userInfo.IsPrivate,
		IsAdmin:         userInfo.IsAdmin,
		IsFollowing:     userInfo.IsFollowing,
		IsFollowed:      userInfo.IsFollowed,
	}
}

func convertToUserInfoWithoutBios(userInfos []*model.UserInfoInternal) []*model.UserInfoWithoutBio {
	userInfo := make([]*model.UserInfoWithoutBio, len(userInfos))
	for i, u := range userInfos {
		userInfo[i] = convertToUserInfoWithoutBio(u)
	}

	return userInfo
}