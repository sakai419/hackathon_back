package account

import (
	"context"
	"time"
	"local-test/internal/database/sqlc"
)

// AccountRepository はアカウント関連の操作を定義するインターフェース
type AccountRepository interface {
	CreateAccount(ctx context.Context, params sqlc.CreateAccountParams) error
	GetAccountById(ctx context.Context, id string) (sqlc.Account, error)
	GetAccountByUserId(ctx context.Context, userID string) (sqlc.Account, error)
	GetAccountByUserName(ctx context.Context, userName string) (sqlc.Account, error)
	UpdateAccountUserName(ctx context.Context, params sqlc.UpdateAccountUserNameParams) error
	DeleteAccount(ctx context.Context, id string) error
	CheckUserIdExists(ctx context.Context, userID string) (bool, error)
	CheckUserNameExists(ctx context.Context, userName string) (bool, error)
	CountAccounts(ctx context.Context) (int64, error)
	GetAccountCreationDate(ctx context.Context, id string) (time.Time, error)
	SearchAccountsByUserId(ctx context.Context, params sqlc.SearchAccountsByUserIdParams) ([]sqlc.Account, error)
	SearchAccountsByUserName(ctx context.Context, params sqlc.SearchAccountsByUserNameParams) ([]sqlc.Account, error)
}

// SQLAccountRepository はsqlcを使用してAccountRepositoryを実装する構造体
type SQLAccountRepository struct {
	q *sqlc.Queries
}

// NewSQLAccountRepository は新しいSQLAccountRepositoryインスタンスを作成します
func NewSQLAccountRepository(q *sqlc.Queries) *SQLAccountRepository {
	return &SQLAccountRepository{q: q}
}

// CreateAccount はアカウントを作成します
func (r *SQLAccountRepository) CreateAccount(ctx context.Context, params sqlc.CreateAccountParams) error {
	return r.q.CreateAccount(ctx, params)
}

// GetAccountById はIDでアカウントを取得します
func (r *SQLAccountRepository) GetAccountById(ctx context.Context, id string) (sqlc.Account, error) {
	return r.q.GetAccountById(ctx, id)
}

// GetAccountByUserId はユーザーIDでアカウントを取得します
func (r *SQLAccountRepository) GetAccountByUserId(ctx context.Context, userID string) (sqlc.Account, error) {
	return r.q.GetAccountByUserId(ctx, userID)
}

// GetAccountByUserName はユーザー名でアカウントを取得します
func (r *SQLAccountRepository) GetAccountByUserName(ctx context.Context, userName string) (sqlc.Account, error) {
	return r.q.GetAccountByUserName(ctx, userName)
}

// UpdateAccountUserName はアカウントのユーザー名を更新します
func (r *SQLAccountRepository) UpdateAccountUserName(ctx context.Context, params sqlc.UpdateAccountUserNameParams) error {
	return r.q.UpdateAccountUserName(ctx, params)
}

// DeleteAccount はアカウントを削除します
func (r *SQLAccountRepository) DeleteAccount(ctx context.Context, id string) error {
	return r.q.DeleteAccount(ctx, id)
}

// CheckUserIdExists はユーザーIDが存在するかチェックします
func (r *SQLAccountRepository) CheckUserIdExists(ctx context.Context, userID string) (bool, error) {
	return r.q.CheckUserIdExists(ctx, userID)
}

// CheckUserNameExists はユーザー名が存在するかチェックします
func (r *SQLAccountRepository) CheckUserNameExists(ctx context.Context, userName string) (bool, error) {
	return r.q.CheckUserNameExists(ctx, userName)
}

// CountAccounts はアカウントの総数を取得します
func (r *SQLAccountRepository) CountAccounts(ctx context.Context) (int64, error) {
	return r.q.CountAccounts(ctx)
}

// GetAccountCreationDate はアカウントの作成日を取得します
func (r *SQLAccountRepository) GetAccountCreationDate(ctx context.Context, id string) (time.Time, error) {
	return r.q.GetAccountCreationDate(ctx, id)
}

// SearchAccountsByUserId はユーザーIDでアカウントを検索します
func (r *SQLAccountRepository) SearchAccountsByUserId(ctx context.Context, params sqlc.SearchAccountsByUserIdParams) ([]sqlc.Account, error) {
	return r.q.SearchAccountsByUserId(ctx, params)
}

// SearchAccountsByUserName はユーザー名でアカウントを検索します
func (r *SQLAccountRepository) SearchAccountsByUserName(ctx context.Context, params sqlc.SearchAccountsByUserNameParams) ([]sqlc.Account, error) {
	return r.q.SearchAccountsByUserName(ctx, params)
}