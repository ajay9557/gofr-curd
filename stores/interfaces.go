package stores

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Store interface {
	GetId(ctx *gofr.Context, id int) (models.Product, error)
	DeleteId(ctx *gofr.Context, id int) error
	UpdateId(ctx *gofr.Context, product models.Product) error
	CreateProducts(ctx *gofr.Context, product models.Product) (models.Product, error)
	GetAll(ctx *gofr.Context) ([]models.Product, error)
}
