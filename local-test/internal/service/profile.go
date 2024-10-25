package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) UpdateProfiles(ctx context.Context, arg *model.UpdateProfilesParams) error {
	// Validate the input
	if err := arg.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

	// Update profiles
	if err := s.repo.UpdateProfiles(ctx, arg); err != nil {
		var duplicateErr *apperrors.ErrDuplicateEntry
		if errors.As(err, &duplicateErr) {
			return apperrors.NewDuplicateEntryAppError("profiles", "update profiles", err)
		}
		return apperrors.NewNotFoundAppError("profiles", "update profiles", err)
	}

	return nil
}