package services

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Service interface {
	GetByUserId(ctx *gofr.Context, id int) (models.Product, error)
	DeleteByProductId(ctx *gofr.Context, id int) error
	UpdateByProductId(ctx *gofr.Context, product models.Product) error
	InsertProduct(ctx *gofr.Context, product models.Product) (models.Product, error)
	GetProducts(ctx *gofr.Context) ([]models.Product, error)
}
