package application

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	golib "github.com/umardev500/go-lib"
	"github.com/umardev500/product-rpc/domain"
	"github.com/umardev500/product-rpc/domain/model"
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

// DeleteProduct is method to delete the product by id
func (p *productUsecase) DeleteProduct(
	ctx context.Context,
	req *proto.DeleteProductRequest,
) error {
	id, _ := primitive.ObjectIDFromHex(req.Id)
	return p.repo.Delete(ctx, id)
}

func (p *productUsecase) FindProduct(ctx context.Context) (products []*proto.Product, err error) {
	cur, err := p.repo.Find(ctx, 1, 5)
	if err != nil {
		return
	}

	for cur.Next(ctx) {
		var row model.ProductModel
		err := cur.Decode(&row)
		if err != nil {
			log.Error().Msg(err.Error())
			break
		}
		// copy model to proto struct generated
		var productProto *proto.Product
		golib.CopyStruct(row, &productProto)
		var timeTract *proto.TimeTrack
		golib.CopyStruct(row.TimeTrack, timeTract)
		productProto.TimeTrack = timeTract

		// p.parseToRPCStruct(row, proto.Product{})

		products = append(products, productProto)
	}

	return
}

func (p *productUsecase) CountProducts(ctx context.Context) (int64, error) {
	return p.repo.Count(ctx)
}
