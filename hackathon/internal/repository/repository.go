package repository

import (
	"database/sql"
	"local-test/internal/sqlc/sqlcgen"
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