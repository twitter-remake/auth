package repository

import (
	"context"
	"time"
)

type Identity struct {
	ID        string
	UID       string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type RegisterParams struct {
	UUID  string
	UID   string
	Email string
}

func (r *Dependency) SaveIdentity(ctx context.Context, params RegisterParams) error {
	query := "INSERT INTO identities (id, uid, email) VALUES ($1, $2, $3)"
	if _, err := r.db.Exec(ctx, query, params.UUID, params.UID, params.Email); err != nil {
		return err
	}

	return nil
}

func (r *Dependency) IdentityExistsByUID(ctx context.Context, uid string) (bool, error) {
	var exists bool

	query := "SELECT EXISTS (SELECT 1 FROM identities WHERE uid = $1)"
	if err := r.db.QueryRow(ctx, query, uid).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}
