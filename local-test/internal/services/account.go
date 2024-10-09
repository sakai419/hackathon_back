package services

import (
	"context"
	"errors"
	"fmt"
	"local-test/internal/models"
	"local-test/internal/repositories"
	"local-test/pkg/utils"
	"net/http"
)

func (s *Service) CreateAccount(ctx context.Context, arg *models.CreateAccountParams) error {
    if err := s.repo.CreateAccount(ctx, arg); err != nil {
        return &utils.AppError{
            Status:  http.StatusInternalServerError,
            Code:    "DATABASE_ERROR",
            Message: "Failed to create account",
            Err:     fmt.Errorf("service: failed to create account: %w", err),
        }
    }

    return nil
}

func (s *Service) DeleteMyAccount(ctx context.Context, id string) error {
	if err := s.repo.DeleteMyAccount(ctx, id); err != nil {
		if errors.Is(err, repositories.ErrAccountNotFound) {
			return &utils.AppError{
				Status:  http.StatusNotFound,
				Code:    "ACCOUNT_NOT_FOUND",
				Message: "Account not found",
				Err:     fmt.Errorf("service: failed to delete account: %w", err),
			}
		}
		return &utils.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "DATABASE_ERROR",
			Message: "Failed to delete account",
			Err:     fmt.Errorf("service: failed to delete account: %w", err),
		}
	}

	return nil
}