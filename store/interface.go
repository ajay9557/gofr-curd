package store

import (
	"zopsmart/gofr-curd/model"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Productstorer interface {
	GetProductById(*gofr.Context, int) (model.Product, error)
	GetProducts(*gofr.Context) ([]model.Product, error)
	AddProduct(*gofr.Context, model.Product) (int, error)
	DeleteById(*gofr.Context,int) error
	UpdateById(*gofr.Context,model.Product) (model.Product,error)
}
