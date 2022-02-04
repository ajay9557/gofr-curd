package services

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Iservice interface {
	GetProductById(ctx *gofr.Context, id string) (*models.Product, error)
	GetAllProducts(ctx *gofr.Context) ([]*models.Product, error)
	// CreateProduct(ctx *gofr.Context, prd models.Product) (int, error)
	CreateProduct(ctx *gofr.Context, prd models.Product) (*models.Product, error)
	DeleteById(ctx *gofr.Context, id string) error
	UpdateById(ctx *gofr.Context, id string, prd models.Product) (*models.Product, error)
}
