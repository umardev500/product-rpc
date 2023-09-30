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

// FindProduct find all products
func (p *productDelivery) FindProduct(
	ctx context.Context,
	req *proto.FindProductRequest,
) (*proto.FindProductResponse, error) {
	products, err := p.usecase.FindProduct(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.FindProductResponse{Data: products}, nil
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

// UpdateProduct update the product by the id
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

// Delete product is delete the product by the id
func (p *productDelivery) DeleteProduct(
	ctx context.Context,
	req *proto.DeleteProductRequest,
) (*proto.DeleteProductResponse, error) {
	err := p.usecase.DeleteProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return &proto.DeleteProductResponse{Success: true}, nil
}
