package service

import (
	"zopsmart/gofr-curd/model"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Productservice interface {
	GetByID(ctx *gofr.Context, id string) (model.Product, error)
}
