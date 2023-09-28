package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/product-rpc/application"
	"github.com/umardev500/product-rpc/config"
)

func init() {
	// Setup logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	// Load environment file
	// use for development only
	// production we use system env for best practice
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	// create main context with signal
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	// Create new application instance
	// and start the application
	DB := config.NewMongo().Database("store_product")
	app := application.NewApp(DB)
	err := app.Start(ctx)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
