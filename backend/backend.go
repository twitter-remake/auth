package backend

import (
	"firebase.google.com/go/v4/auth"
	"github.com/twitter-remake/auth/repository"
)

// Backend is the main backend struct for the business logic layer
type Backend struct {
	repo *repository.Repository
	auth *auth.Client
}

// New creates a new Backend struct
func New(repo *repository.Repository, auth *auth.Client) *Backend {
	return &Backend{
		repo: repo,
		auth: auth,
	}
}
