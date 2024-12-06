package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) ExecuteCode(ctx context.Context, params *model.ExecuteCodeParams) (*model.ExecuteResult, error) {
	// Validate input
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Execute code
	switch params.Code.Language {
	case "c":
		if ret, err := s.repo.ExecuteCCode(ctx, params.Code.Content); err != nil {
			return nil, apperrors.NewInternalAppError("execute code", err)
		} else {
			return ret, nil
		}
	default:
		return nil, apperrors.NewInternalAppError("execute code", errors.New("unsupported language"))
	}
}