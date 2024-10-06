package repositories

import (
	"context"
	"fmt"
	"local-test/pkg/database/generated"
)

func (r *Repository) CreateAccount(ctx context.Context, arg *queries.CreateAccountParams) (error) {
	if err := r.q.CreateAccount(ctx, *arg); err != nil {
		return fmt.Errorf("repository: failed to create account: %w", err)
	}
	return nil
}

func (r *Repository) DeleteAccount(ctx context.Context, id string) (error) {
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

func (r *Repository) SuspendAccount(ctx context.Context, id string) (error) {
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

func (r *Repository) UnsuspendAccount(ctx context.Context, id string) (error) {
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