package application

import (
	"context"

	"github.com/umardev500/product-rpc/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type productUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}

func (p *productUsecase) CreateProduct(ctx context.Context, product bson.D) error {
	return p.repo.Create(ctx, product)
}
