package service

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Service interface {
	GetByProductID(id int, ctx *gofr.Context) (models.Product, error)
	GetProducts(ctx *gofr.Context) ([]models.Product, error)
	InsertProductDetails(product models.Product, ctx *gofr.Context) error
	UpdateProductDetails(product models.Product, ctx *gofr.Context) error
	DeleteProductByID(id int, ctx *gofr.Context) error
}
