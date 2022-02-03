package store

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/ridhdhish-desai-zs/product-gofr/models"
)

type Product interface {
	GetById(ctx *gofr.Context, id int) (*models.Product, error)
	Get(ctx *gofr.Context) ([]*models.Product, error)
}
