package store

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Store interface {
	Get(ctx *gofr.Context) ([]*models.Product, error)
	GetByID(ctx *gofr.Context, id int) (*models.Product, error)
	Create(ctx *gofr.Context, pd models.Product) error
	Update(ctx *gofr.Context, pd models.Product) error
	Delete(ctx *gofr.Context, id int) error
}
