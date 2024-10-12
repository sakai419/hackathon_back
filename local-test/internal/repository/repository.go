package repository

import (
	"database/sql"
	sqlcgen "local-test/internal/sqlc/generated"
)

type Repository struct {
	db *sql.DB
	q *sqlcgen.Queries
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
		q: sqlcgen.New(db),
	}
}