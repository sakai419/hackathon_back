package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"local-test/internal/database/sqlc"
)

type AccountRepository struct {
	q *sqlc.Queries
}

var ErrAccountNotFound = errors.New("repository: account not found")

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		q: sqlc.New(db),
	}
}

func (r *AccountRepository) CreateAccount(ctx context.Context, arg *sqlc.CreateAccountParams) (error) {
	if err := r.q.CreateAccount(ctx, *arg); err != nil {
		return fmt.Errorf("repository: failed to create account: %w", err)
	}
	return nil
}

func (r *AccountRepository) DeleteAccount(ctx context.Context, id string) (error) {
	res, err := r.q.DeleteAccount(ctx, id)
	if err != nil {
		return fmt.Errorf("repository: failed to delete account: %w", err)
	}

	num, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if num == 0 {
		return ErrAccountNotFound
	}
	return nil
}

func (r *AccountRepository) SuspendAccount(ctx context.Context, id string) (error) {
	res, err := r.q.SuspendAccount(ctx, id)
	if err != nil {
		return fmt.Errorf("repository: failed to suspend account: %w", err)
	}

	num, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if num == 0 {
		return ErrAccountNotFound
	}

	return nil
}

func (r *AccountRepository) UnsuspendAccount(ctx context.Context, id string) (error) {
	res, err := r.q.UnsuspendAccount(ctx, id)
	if err != nil {
		return fmt.Errorf("repository: failed to unsuspend account: %w", err)
	}

	num, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if num == 0 {
		return ErrAccountNotFound
	}

	return nil
}