package store

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/ridhdhish-desai-zs/product-gofr/models"
)

type Product interface {
	GetByID(ctx *gofr.Context, id int) (*models.Product, error)
	Get(ctx *gofr.Context) ([]*models.Product, error)
	Create(ctx *gofr.Context, pr models.Product) error
	UpdateByID(ctx *gofr.Context, id int, pr models.Product) error
	DeleteByID(ctx *gofr.Context, id int) error
}
