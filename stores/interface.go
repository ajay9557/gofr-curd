package stores

import (
	"productGofr/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Istore interface {
	UserById(ctx *gofr.Context, id int) (*models.Product, error)
	GetAllUsers(ctx *gofr.Context) ([]*models.Product, error)
	// CreateProduct(ctx *gofr.Context, prd models.Product) (int, error)
	CreateProduct(ctx *gofr.Context, prd models.Product) (int, error)
	DeleteById(ctx *gofr.Context, id int) error
	UpdateById(ctx *gofr.Context, id int, prd models.Product) (int, error)
}
