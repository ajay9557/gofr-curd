package services

import (
	"product/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Service interface {
	GetProductById(ctx *gofr.Context, id int) (models.Product, error)
	GetAllProduct(ctx *gofr.Context) ([]models.Product, error)
	AddProduct(ctx *gofr.Context, product models.Product) error
	UpdateProduct(ctx *gofr.Context, product models.Product) error
	DeleteProduct(ctx *gofr.Context, id int) error
}
