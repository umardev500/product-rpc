package application

import (
	"context"

	"github.com/umardev500/product-rpc/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Update create update for the product
func (p *productRepository) Update(ctx context.Context, product bson.D, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": product}
	_, err := p.collection.UpdateOne(ctx, filter, update)
	return err
}

// Delete deletet the product by the id
func (p *productRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := p.collection.DeleteOne(ctx, filter)
	return err
}

func (p *productRepository) Find(ctx context.Context) (*mongo.Cursor, error) {
	filter := bson.M{}
	return p.collection.Find(ctx, filter)
}
