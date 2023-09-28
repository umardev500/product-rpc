package application

import (
	"context"

	"github.com/umardev500/product-rpc/domain"
	"github.com/umardev500/store/proto"
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
	err := p.usecase.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return &proto.CreateProductResponse{Success: true}, nil
}

func (p *productDelivery) UpdateProduct(
	ctx context.Context,
	req *proto.UpdateProductRequest,
) (*proto.UpdateProductResponse, error) {
	err := p.usecase.UpdateProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateProductResponse{Success: true}, nil
}
