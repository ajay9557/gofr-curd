package services

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Service interface {
	GetByUserID(ctx *gofr.Context, id int) (models.Product, error)
	DeleteByProductID(ctx *gofr.Context, id int) error
	UpdateByProductID(ctx *gofr.Context, product models.Product) error
	InsertProduct(ctx *gofr.Context, product models.Product) (models.Product, error)
	GetProducts(ctx *gofr.Context) ([]models.Product, error)
}
