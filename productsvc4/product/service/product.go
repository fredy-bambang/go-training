package service

import (
	"context"

	"github.com/rickywinata/go-training/productsvc4/product"
)

type (
	// CreateProductCommand represents the parameters for creating a product.
	CreateProductCommand struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}
)

// ProductService is an interface for product operations.
type ProductService interface {
	CreateProduct(ctx context.Context, cmd *CreateProductCommand) (*product.Product, error)
}

type productService struct {
	productRepo product.Repository
}

// NewProductService creates a new product service.
func NewProductService(productRepo1 product.Repository) ProductService {
	return &productService{
		productRepo: productRepo1,
	}
}

func (s *productService) CreateProduct(ctx context.Context, cmd *CreateProductCommand) (*product.Product, error) {
	p := &product.Product{
		Name:  cmd.Name,
		Price: cmd.Price,
	}

	if err := s.productRepo.Insert(ctx, p); err != nil {
		return nil, err
	}

	return p, nil
}
