package service

import (
	"gofr-curd/model"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Service interface {
	GetByProductId(id int, ctx *gofr.Context) (model.ProductDetails, error)
	DeleteByProductId(ctx *gofr.Context, id int) error
	UpdateByProductId(ctx *gofr.Context, product model.ProductDetails) error
	InsertProduct(ctx *gofr.Context, product model.ProductDetails) (model.ProductDetails, error)
	GetProducts(ctx *gofr.Context) ([]model.ProductDetails, error)
}
