package clients

import (
	"context"
	"net"
	"strconv"

	"firebase.google.com/go/v4/auth"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/twitter-remake/auth/config"
	"golang.org/x/sync/errgroup"
)

type Clients struct {
	ctx context.Context

	PostgreSQL      *pgxpool.Pool
	Auth            *auth.Client
	ServiceRegistry *Consul
	UserService     *UserServiceClient
}

func New(ctx context.Context) (*Clients, error) {
	c := new(Clients)
	c.ctx = ctx

	var group errgroup.Group

	group.Go(func() error {
		var err error
		c.PostgreSQL, err = NewPostgreSQLClient(ctx, config.DatabaseURL())
		if err != nil {
			return errors.Wrap(err, "initializing postgresql")
		}
		return nil
	})

	group.Go(func() error {
		var err error
		c.Auth, err = NewFirebaseAuthClient(ctx, "./firebase-credentials.json")
		if err != nil {
			return errors.Wrap(err, "initializing firebase auth")
		}
		return nil
	})

	group.Go(func() error {
		var err error
		c.ServiceRegistry, err = NewConsulAPI()
		if err != nil {
			return errors.Wrap(err, "initializing consul")
		}

		if err := c.ServiceRegistry.Register(&RegisterCfg{
			ID:          config.AppName(),
			Host:        config.Host(),
			Port:        config.Port(),
			Environment: config.Environment(),
		}); err != nil {
			return errors.Wrap(err, "initializing consul")
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		return nil, err
	}

	services, err := c.ServiceRegistry.Agent().Services()
	if err != nil {
		return nil, errors.Wrap(err, "finding user service")
	}

	userService, ok := services["user-service"]
	if !ok {
		return nil, errors.New("user service not found")
	}

	userHost := net.JoinHostPort(
		userService.Address,
		strconv.Itoa(userService.Port))

	log.Debug().Msgf("user service found at %s", userHost)

	c.UserService = NewUserClient(userHost)

	return c, nil
}
