package account

import (
	"context"
	"errors"
	"fmt"
	"local-test/internal/database/sqlc"
	"local-test/pkg/utils"
	"net/http"
)

type AccountService struct {
    repo *AccountRepository
}

func NewAccountService(repo *AccountRepository) *AccountService {
    return &AccountService{
        repo: repo,
    }
}

func (s *AccountService) CreateAccount(ctx context.Context, arg *sqlc.CreateAccountParams) error {
    // 入力値のバリデーション
    if err := validateCreateAccountParams(arg); err != nil {
        return &utils.AppError{
            Status:  http.StatusBadRequest,
            Code:    "INVALID_INPUT",
            Message: "Invalid input for account creation",
            Err:     err,
        }
    }

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

func (s *AccountService) DeleteAccount(ctx context.Context, id string) error {
    return s.performAccountOperation(ctx, id, s.repo.DeleteAccount, "delete")
}

func (s *AccountService) SuspendAccount(ctx context.Context, id string) error {
    return s.performAccountOperation(ctx, id, s.repo.SuspendAccount, "suspend")
}

func (s *AccountService) UnsuspendAccount(ctx context.Context, id string) error {
    return s.performAccountOperation(ctx, id, s.repo.UnsuspendAccount, "unsuspend")
}

func (s *AccountService) performAccountOperation(ctx context.Context, id string, operation func(context.Context, string) error, operationName string) error {
    if err := operation(ctx, id); err != nil {
        if errors.Is(err, ErrAccountNotFound) {
            return &utils.AppError{
                Status:  http.StatusNotFound,
                Code:    "ACCOUNT_NOT_FOUND",
                Message: "Account not found",
                Err:     fmt.Errorf("service: failed to %s account: %w", operationName, err),
            }
        }
        return &utils.AppError{
            Status:  http.StatusInternalServerError,
            Code:    "DATABASE_ERROR",
            Message: fmt.Sprintf("Failed to %s account", operationName),
            Err:     fmt.Errorf("service: failed to %s account: %w", operationName, err),
        }
    }
    return nil
}

func validateCreateAccountParams(arg *sqlc.CreateAccountParams) error {
    if arg.ID == "" {
        return errors.New("ID is required")
    } else if arg.UserID == "" {
        return errors.New("UserID is required")
    } else if arg.UserName == "" {
        return errors.New("UserName is required")
    }
    return nil
}