package application

import (
	"context"

	"github.com/umardev500/product-rpc/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// Find retrieves documents from the MongoDB collection.
func (p *productRepository) Find(
	ctx context.Context,
	pageNum,
	pageSize int64,
) (*mongo.Cursor, error) {
	filter := bson.M{}
	skip := (pageNum - 1) * pageSize
	findOpt := options.Find()
	findOpt.SetSkip(skip)
	findOpt.SetLimit(pageSize)

	return p.collection.Find(ctx, filter, findOpt)
}

// Count returns the total count of documents in the product collection.
func (p *productRepository) Count(
	ctx context.Context,
	pageNum,
	pageSize int64,
) (int64, error) {
	filter := bson.M{}
	skip := (pageNum - 1) * pageSize
	count := options.Count()
	count.SetSkip(skip)
	count.SetLimit(pageSize)

	return p.collection.CountDocuments(ctx, filter, count)
}
