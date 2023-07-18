package clients

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func NewFirebaseAuthClient(ctx context.Context, credentialFile string) (*auth.Client, error) {
	opt := option.WithCredentialsFile(credentialFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
