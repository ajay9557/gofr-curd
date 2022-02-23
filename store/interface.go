package store

import (
	"zopsmart/gofr-curd/model"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Productstorer interface {
	GetProductByID(*gofr.Context, int) (model.Product, error)
	GetProducts(*gofr.Context) ([]model.Product, error)
	AddProduct(*gofr.Context, model.Product) (int, error)
	DeleteByID(*gofr.Context, int) error
	UpdateByID(*gofr.Context, model.Product) (model.Product, error)
}
