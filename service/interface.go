package service

import (
	"zopsmart/gofr-curd/model"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Productservice interface {
	GetByID(ctx *gofr.Context, id string) (model.Product, error)
	GetProducts(ctx *gofr.Context) ([]model.Product, error)
	AddProduct(*gofr.Context, model.Product) (model.Product, error)
	DeleteByID(ctx *gofr.Context, id string) error
	UpdateByID(ctx *gofr.Context, prod model.Product, id string) (model.Product, error)
}
