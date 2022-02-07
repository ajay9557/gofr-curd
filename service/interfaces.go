package service

import (
	"gofr-curd/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Services interface {
	Get(ctx *gofr.Context) ([]*models.Product, error)
	GetByID(ctx *gofr.Context, id int) (*models.Product, error)
	Create(ctx *gofr.Context, p models.Product) (*models.Product, error)
	Update(ctx *gofr.Context, p models.Product) (*models.Product, error)
	Delete(ctx *gofr.Context, id int) error
}
