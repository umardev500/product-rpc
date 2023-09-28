package application

import (
	"context"
	"net"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/store/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Application struct {
	DB *mongo.Database
}

func NewApp(DB *mongo.Database) *Application {
	return &Application{
		DB: DB,
	}
}

func (a *Application) Start(ctx context.Context) error {
	host := os.Getenv("HOST")
	// Setup listener
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer lis.Close()

	// Create new grpc server
	srv := grpc.NewServer()

	// injection
	repo := NewProductRepository(a.DB.Collection("products"))
	usecase := NewProductUsecase(repo)
	product := NewProductDelivery(usecase)
	proto.RegisterProductServiceServer(srv, product)

	// Start server to goroutine
	// so we can handle multiple state with non blocking
	ch := make(chan error, 1)
	go func() {
		log.Info().Msg("gRPC server started")
		// Start server
		err := srv.Serve(lis)
		if err != nil {
			ch <- err
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		log.Info().Msg("Gracfully shutdown gRPC server")
		srv.GracefulStop()
	}

	return nil
}
