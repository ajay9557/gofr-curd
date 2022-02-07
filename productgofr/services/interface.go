package service

import (
	models "zopsmart/productgofr/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Services interface {
	GetProdByID(ctx *gofr.Context, id int) (*models.Product, error)
	CreateProduct(ctx *gofr.Context, pro *models.Product) (*models.Product,error)
	UpdateProduct(ctx *gofr.Context,pro models.Product) (*models.Product,error)
	DeleteProduct(ctx *gofr.Context, id int) (error)
	GetAllProd(ctx *gofr.Context) ([]*models.Product,error)
}
