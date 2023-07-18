package backend

import (
	"context"
	"time"

	"github.com/twitter-remake/auth/repository"
)

type RegisterInput struct {
	IDToken    string
	Name       string
	ScreenName string
	Location   string
	BirthDate  time.Time
}

func (b *Backend) SignIn(ctx context.Context, input RegisterInput) error {
	// TODO: send other params to user service
	idToken, err := b.auth.VerifyIDToken(ctx, input.IDToken)
	if err != nil {
		return err
	}

	exists, err := b.repo.IdentityExistsByUID(ctx, idToken.UID)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	email := idToken.Firebase.Identities["email"].([]any)[0].(string)
	err = b.repo.SaveIdentity(ctx, repository.RegisterParams{
		UID:   idToken.UID,
		Email: email,
	})
	if err != nil {
		return err
	}

	return nil
}
