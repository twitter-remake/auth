package backend

import (
	"firebase.google.com/go/v4/auth"
	"github.com/twitter-remake/auth/clients"
	"github.com/twitter-remake/auth/repository"
)

// Backend is the main backend struct for the business logic layer
type Dependency struct {
	repo        *repository.Dependency
	auth        *auth.Client
	userService *clients.UserServiceClient
}

// New creates a new Backend struct
func New(repo *repository.Dependency, auth *auth.Client, userService *clients.UserServiceClient) *Dependency {
	return &Dependency{
		repo:        repo,
		auth:        auth,
		userService: userService,
	}
}
