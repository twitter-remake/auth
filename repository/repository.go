package repository

import "github.com/jackc/pgx/v5/pgxpool"

// Repository is the main repository struct for the data access layer
type Repository struct {
	db *pgxpool.Pool
}

// New creates a new Repository struct
func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}
