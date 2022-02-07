package stores

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Store interface {
	GetID(ctx *gofr.Context, id int) (models.Product, error)
	DeleteID(ctx *gofr.Context, id int) error
	UpdateID(ctx *gofr.Context, product models.Product) error
	CreateProducts(ctx *gofr.Context, product models.Product) (models.Product, error)
	GetAll(ctx *gofr.Context) ([]models.Product, error)
}
