package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) SearchUsers(ctx context.Context, params *model.SearchUsersParams) ([]*model.UserInfo, error) {
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

    var users []*model.UserInfoInternal
	switch params.SortType {
		case model.SortTypeNewest, "":
			temp, err := s.repo.SearchUsersOrderByCreatedAt(ctx, &model.SearchUsersOrderByCreatedAtParams{
				Keyword: params.Keyword,
				Offset: params.Offset,
				Limit: params.Limit,
			})
			if err != nil {
				return nil, apperrors.NewInternalAppError("Failed to search users", err)
			}
			users = temp
		default:
			return nil, apperrors.NewInternalAppError("SortType is not supported", errors.New("SortType is not supported"))
	}

	// extract accountIDs
	accountIDs := make([]string, 0)
	for _, u := range users {
		accountIDs = append(accountIDs, u.ID)
	}

	// Filter accessible account ids
	accessibleAccountIDs, err := s.repo.FilterAccessibleAccountIDs(ctx, &model.FilterAccesibleAccountIDsParams{
		ClientAccountID: params.ClientAccountID,
		AccountIDs: accountIDs,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("Failed to filter accessible account ids", err)
	}

	// Filter users
	userInfos := filterUsers(users, accessibleAccountIDs)

	return userInfos, nil
}

func filterUsers(users []*model.UserInfoInternal, accessibleAccountIDs []string) []*model.UserInfo {
	// create map of accessible account ids
	accessibleAccountIDsMap := make(map[string]bool)
	for _, id := range accessibleAccountIDs {
		accessibleAccountIDsMap[id] = true
	}

	filteredUsers := make([]*model.UserInfo, 0)
	for _, u := range users {
		if _, ok := accessibleAccountIDsMap[u.ID]; ok {
			filteredUsers = append(filteredUsers, &model.UserInfo{
				UserID: u.UserID,
				UserName: u.UserName,
				Bio: u.Bio,
				ProfileImageURL: u.ProfileImageURL,
				IsPrivate: u.IsPrivate,
				IsAdmin: u.IsAdmin,
			})
		}
	}
	return filteredUsers
}