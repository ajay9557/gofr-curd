package store

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
)

//go:generate mockgen -package=store --source=interface.go --destination=mock_interface.go

type Store interface {
	GetAll(ctx *gofr.Context) ([]*models.Product, error)
	GetByID(ctx *gofr.Context, id int) (*models.Product, error)
	Update(ctx *gofr.Context, product *models.Product) error
	Delete(ctx *gofr.Context, id int) error
	Create(ctx *gofr.Context, product *models.Product) (*models.Product, error)
}
