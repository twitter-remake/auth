package main

import (
	"context"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/twitter-remake/auth/api"
	"github.com/twitter-remake/auth/backend"
	"github.com/twitter-remake/auth/clients"
	"github.com/twitter-remake/auth/config"
	"github.com/twitter-remake/auth/repository"
)

func init() {
	// Setup logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.With().Caller().Stack().Logger()
	if os.Getenv("ENVIRONMENT") == "dev" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	log.Info().Msg("Starting Twitter Auth Service")
	config.Init()
}

func main() {
	ctx := context.Background()

	// Setup clients
	firebaseAuth, err := clients.NewFirebaseAuthClient(ctx, "./firebase-credentials.json")
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	postgresClient, err := clients.NewPostgreSQLClient(ctx, config.DatabaseURL())
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	// Initialize layers
	repository := repository.New(postgresClient)
	backend := backend.New(repository, firebaseAuth)
	api := api.New(backend)

	// Start server and wait for shutdown signals
	exitSignal := api.Start(config.Host(), config.Port())

	// If a shutdown signal is received (e.g. CTRL + C or kill) shutdown gracefully
	// signal stored in variable for logging purposes
	signal := <-exitSignal
	api.Shutdown(ctx, signal)
}
