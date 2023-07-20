package clients

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"firebase.google.com/go/v4/auth"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/twitter-remake/auth/config"
	"golang.org/x/sync/errgroup"
)

type Clients struct {
	ctx context.Context

	PostgreSQL      *pgxpool.Pool
	Auth            *auth.Client
	ServiceRegistry *consulapi.Client
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

		if err := c.registerServiceToConsul(); err != nil {
			return errors.Wrap(err, "initializing consul")
		}
		return nil
	})

	group.Go(func() error {
		c.UserService = NewUserClient("")
		return nil
	})

	if err := group.Wait(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Clients) registerServiceToConsul() error {
	var address string
	if config.Environment() == "dev" {
		// assuming consul is running in docker
		address = "host.docker.internal"
	} else {
		address = config.Host()
	}

	check := &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s/", net.JoinHostPort(address, config.Port())),
		Interval:                       "10s",
		Timeout:                        "30s",
		CheckID:                        fmt.Sprintf("service:%s:http", config.AppName()),
		DeregisterCriticalServiceAfter: "1m",
		TLSSkipVerify:                  func() bool { return config.Environment() == "dev" }(),
	}

	port, _ := strconv.Atoi(config.Port())

	serviceDefinition := &consulapi.AgentServiceRegistration{
		ID:      config.AppName(),
		Name:    config.AppName() + "_master",
		Port:    port,
		Address: address,
		Tags:    []string{config.Environment(), "auth-service"},
		Check:   check,
	}

	if err := c.ServiceRegistry.Agent().ServiceRegister(serviceDefinition); err != nil {
		return err
	}

	return nil
}
