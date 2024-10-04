package account

import (
	"context"
	"database/sql"
	"fmt"
	"local-test/internal/database/sqlc"
)

type AccountRepository struct {
	q *sqlc.Queries
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		q: sqlc.New(db),
	}
}

func (r *AccountRepository) CreateAccount(ctx context.Context, arg *sqlc.CreateAccountParams) (error) {
	if err := r.q.CreateAccount(ctx, *arg); err != nil {
		return fmt.Errorf("repository: error creating account: %v", err)
	}
	return nil
}