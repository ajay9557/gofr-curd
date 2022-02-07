package store

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Store interface {
	GetByID(id int, ctx *gofr.Context) (models.Product, error)
	GetAllProducts(ctx *gofr.Context) ([]models.Product, error)
	InsertProduct(product models.Product, ctx *gofr.Context) error
	UpdateProduct(product models.Product, ctx *gofr.Context) error
	DeleteByID(id int, ctx *gofr.Context) error
}
