package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) CreateReport(ctx context.Context, params *model.CreateReportParams) error {
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

	// Create report
	if err := s.repo.CreateReport(ctx, params); err != nil {
		return apperrors.NewInternalAppError("create report", err)
	}

	return nil
}