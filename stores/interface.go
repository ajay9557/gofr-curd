package stores

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Istore interface {
	GetProductByID(ctx *gofr.Context, id int) (*models.Product, error)
	GetAllProducts(ctx *gofr.Context) ([]*models.Product, error)
	// CreateProduct(ctx *gofr.Context, prd models.Product) (int, error)
	CreateProduct(ctx *gofr.Context, prd models.Product) (int, error)
	DeleteByID(ctx *gofr.Context, id int) error
	UpdateByID(ctx *gofr.Context, id int, prd models.Product) (int, error)
}
