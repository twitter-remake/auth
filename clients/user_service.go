package clients

import (
	"net/http"

	userpb "github.com/twitter-remake/auth/proto/gen/go/user"
)

type UserServiceClient struct {
	Listing userpb.Listing
	Profile userpb.Profile
}

func NewUserClient(baseURL string) *UserServiceClient {
	listingClient := userpb.NewListingProtobufClient(baseURL, &http.Client{})
	profileClient := userpb.NewProfileProtobufClient(baseURL, &http.Client{})

	return &UserServiceClient{
		Listing: listingClient,
		Profile: profileClient,
	}
}
