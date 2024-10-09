package repositories

import (
	"database/sql"
	"local-test/internal/sqlc/generated"
)

type Repository struct {
	q *sqlcgen.Queries
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		q: sqlcgen.New(db),
	}
}