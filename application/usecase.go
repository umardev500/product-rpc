package application

import (
	"context"
	"time"

	golib "github.com/umardev500/go-lib"
	"github.com/umardev500/product-rpc/domain"
	"github.com/umardev500/store/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}

// CreateProduct is method to create new product
func (p *productUsecase) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) error {
	// Append created at struct
	req.Product.TimeTrack = &proto.TimeTrack{
		CreatedAt: 100,
	}
	product := golib.StructToBson(req, false, "json")[0].Value.(bson.D)
	return p.repo.Create(ctx, product)
}

// UpdateProduct is method to update the product
func (p *productUsecase) UpdateProduct(
	ctx context.Context,
	req *proto.UpdateProductRequest,
) error {
	req.Product.TimeTrack = &proto.TimeTrack{
		UpdatedAt: time.Now().UTC().Unix(),
	}
	id, _ := primitive.ObjectIDFromHex(req.Id)
	product := golib.StructToBson(req.Product, true, "json")
	return p.repo.Update(ctx, product, id)
}
