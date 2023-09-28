package application

import (
	"context"

	"github.com/umardev500/product-rpc/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) domain.ProductRepository {
	return &productRepository{
		collection: collection,
	}
}

// Create insert new product to database
func (p *productRepository) Create(ctx context.Context, data bson.D) error {
	_, err := p.collection.InsertOne(ctx, data)
	return err
}
