package domain

import (
	"context"

	"github.com/umardev500/store/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductUsecase a contract of method the usecase
type ProductUsecase interface {
	CreateProduct(ctx context.Context, req *proto.CreateProductRequest) error
	UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest) error
	DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) error
	FindProduct(
		ctx context.Context,
		pageNum,
		pageSize int64,
	) ([]*proto.Product, error)
	CountProducts(
		ctx context.Context,
		pageNum,
		pageSize int64,
	) (int64, error)
}

// ProductRepository a contract for product repo
type ProductRepository interface {
	Create(ctx context.Context, product bson.D) error
	Update(ctx context.Context, product bson.D, id primitive.ObjectID) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	Find(
		ctx context.Context,
		pageNum,
		pageSize int64,
	) (*mongo.Cursor, error)
	Count(
		ctx context.Context,
		pageNum,
		pageSize int64,
	) (int64, error)
}
