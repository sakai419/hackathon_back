package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) UpdateSettings(ctx context.Context, arg *model.UpdateSettingsParams) error {
	// Update settings
	if err := s.repo.UpdateSettings(ctx, arg); err != nil {
		return apperrors.NewNotFoundAppError("settings", "update settings", err)
	}

	return nil
}