package application

import (
	"context"

	golib "github.com/umardev500/go-lib"
	"github.com/umardev500/product-rpc/domain"
	"github.com/umardev500/store/proto"
	"go.mongodb.org/mongo-driver/bson"
)

type productDelivery struct {
	proto.UnsafeProductServiceServer
	usecase domain.ProductUsecase
}

// NewProductDelivery create new instance of product delivery
func NewProductDelivery(usecase domain.ProductUsecase) proto.ProductServiceServer {
	return &productDelivery{
		usecase: usecase,
	}
}

// CreateProduct create product handler
func (p *productDelivery) CreateProduct(
	ctx context.Context,
	req *proto.CreateProductRequest,
) (*proto.CreateProductResponse, error) {
	lib := golib.NewLib()
	// Append created at struct
	req.Product.TimeTrack = &proto.TimeTrack{
		CreatedAt: 100,
	}
	product := lib.StructToBson(req, "json")
	val := product[0].Value.(bson.D)
	err := p.usecase.CreateProduct(ctx, val)
	if err != nil {
		return nil, err
	}

	return &proto.CreateProductResponse{Success: true}, nil
}
