package store

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/tejas/gofr-crud/models"
)

type ProductStore interface {
	GetProductById(ctx *gofr.Context, id int) (models.Product, error)
	GetAllProducts(ctx *gofr.Context) ([]models.Product, error)
	UpdateProduct(ctx *gofr.Context, prod models.Product) (models.Product, error)
	CreateProduct(ctx *gofr.Context, prod models.Product) (models.Product, error)
	DeleteProduct(ctx *gofr.Context, id int) error

}
