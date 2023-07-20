package repository

import "github.com/jackc/pgx/v5/pgxpool"

// Repository is the main repository struct for the data access layer
type Dependency struct {
	db *pgxpool.Pool
}

// New creates a new Repository struct
func New(db *pgxpool.Pool) *Dependency {
	return &Dependency{db: db}
}
