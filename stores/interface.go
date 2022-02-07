package stores

import (
	"gofr-curd/model"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Store interface {
	GetProductById(id int, ctx *gofr.Context) (model.ProductDetails, error)
	DeleteProductId(ctx *gofr.Context, id int) error
	UpdateProductById(ctx *gofr.Context, prod model.ProductDetails) error
	CreateProducts(ctx *gofr.Context, product model.ProductDetails) (model.ProductDetails, error)
	GetAll(ctx *gofr.Context) ([]model.ProductDetails, error)
}
