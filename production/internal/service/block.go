package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) BlockUser(ctx context.Context, params *model.BlockUserParams) error {
	// Validate input parameters
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

	// Create block
	if err := s.repo.BlockUser(ctx, params); err != nil {
		return apperrors.NewDuplicateEntryAppError("Block", "block", err)
	}

	return nil
}

func (s *Service) UnblockUser(ctx context.Context, params *model.UnblockUserParams) error {
	// Validate input parameters
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

	// Unblock
	if err := s.repo.UnblockUser(ctx, params); err != nil {
		return apperrors.NewNotFoundAppError("Block", "unblock", err)
	}

	return nil
}

func (s *Service) GetBlockedInfos(ctx context.Context, params *model.GetBlockedInfosParams) ([]*model.UserInfo, error) {
	// Validate input parameters
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get blocked account ids
	blockedAccountIDs, err := s.repo.GetBlockedAccountIDs(ctx, &model.GetBlockedAccountIDsParams{
		BlockerAccountID: params.BlockerAccountID,
		Limit:            params.Limit,
		Offset:           params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get blocked account ids", err)
	}

	// Get user and profile info
	infos, err := s.repo.GetUserInfos(ctx, blockedAccountIDs, "dammy")
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("blocked info", "get blocked infos", err)
	}

	// Sort user infos
	blockedInfos := sortUserInfos(infos, blockedAccountIDs)

	return blockedInfos, nil
}

func (s *Service) GetBlockCount(ctx context.Context, accountID string) (int64, error) {
	// Get block count
	count, err := s.repo.GetBlockCount(ctx, accountID)
	if err != nil {
		return 0, apperrors.NewInternalAppError("get block count", err)
	}

	return count, nil
}