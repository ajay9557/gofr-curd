package store

import (
	"zopsmart/gofr-curd/model"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Productstorer interface {
	GetProductById(*gofr.Context, int) (model.Product, error)
}
