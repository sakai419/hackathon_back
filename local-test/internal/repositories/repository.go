package repositories

import (
	"database/sql"
	"local-test/pkg/database/generated"
)

type Repository struct {
	q *queries.Queries
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		q: queries.New(db),
	}
}