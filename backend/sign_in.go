package backend

import (
	"context"
	"database/sql"
	"html"
	"time"

	"github.com/pkg/errors"

	"github.com/rs/zerolog/log"
	userpb "github.com/twitter-remake/auth/proto/gen/go/user"
	"github.com/twitter-remake/auth/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RegisterInput struct {
	IDToken    string
	Name       string
	ScreenName string
	BirthDate  time.Time
}

func (b *Dependency) SignIn(ctx context.Context, input RegisterInput) error {
	// TODO: send other params to user service
	idToken, err := b.auth.VerifyIDToken(ctx, input.IDToken)
	if err != nil {
		return errors.Wrap(err, "failed to verify firebase id token")
	}

	log.Debug().Any("firebase_identity", idToken).Send()

	exists, err := b.repo.IdentityExistsByUID(ctx, idToken.UID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return errors.Wrap(err, "failed to check if identity exists")
	}

	if exists {
		return nil
	}

	var email string
	var photoURL string
	user, err := b.auth.GetUser(ctx, idToken.UID)
	if err != nil {
		log.Err(err).Send()
		// Get the email from the firebase idtoken
		email = idToken.Firebase.Identities["email"].([]any)[0].(string)

		// Use boring avatars API to generate a profile image
		// TODO: use a better API or have a default profile image
		photoURL = "https://hostedboringavatars.vercel.app/api/beam?colors=1DA1F2,14171A,657786,F5F8FA&name=" + html.EscapeString(user.DisplayName)
	} else {
		email = user.Email
		photoURL = user.PhotoURL
	}

	uuid := repository.NewUUID()
	// send to user service first before persisting identity.
	// in case of failure, the identity won't be persisted
	_, err = b.userService.Listing.Register(ctx, &userpb.RegisterRequest{
		Uuid:            uuid,
		Name:            input.Name,
		ScreenName:      input.ScreenName,
		Email:           email,
		BirthDate:       timestamppb.New(input.BirthDate),
		ProfileImageUrl: photoURL,
	})
	if err != nil {
		return errors.Wrap(err, "user-service: failed to register user")
	}

	// Persist identity and send user data to user service
	err = b.repo.SaveIdentity(ctx, repository.RegisterParams{
		UUID:  uuid,
		UID:   idToken.UID,
		Email: email,
	})
	if err != nil {
		return errors.Wrap(err, "failed to save identity")
	}

	return nil
}
