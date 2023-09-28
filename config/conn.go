package config

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	host := os.Getenv("MONGO_HOST")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal().Msg(err.Error())
		err := client.Disconnect(ctx)
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}
	log.Info().Msg("Connected to mongodb")

	return client
}
