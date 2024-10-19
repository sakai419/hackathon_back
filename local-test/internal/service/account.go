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