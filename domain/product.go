package domain

import (
	"context"

	"github.com/umardev500/store/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductUsecase a contract of method the usecase
type ProductUsecase interface {
	CreateProduct(ctx context.Context, req *proto.CreateProductRequest) error
	UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest) error
	DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) error
}

// ProductRepository a contract for product repo
type ProductRepository interface {
	Create(ctx context.Context, product bson.D) error
	Update(ctx context.Context, product bson.D, id primitive.ObjectID) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
