package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// ProductUsecase a contract of method the usecase
type ProductUsecase interface {
	CreateProduct(ctx context.Context, p bson.D) error
}

// ProductRepository a contract for product repo
type ProductRepository interface {
	Create(ctx context.Context, product bson.D) error
	Update(ctx context.Context, product bson.D) error
}
